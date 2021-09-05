package matcher

import (
	"fmt"
	"github.com/SimonAlong/Mikilin-go/constant"
	"github.com/SimonAlong/Mikilin-go/util"
	"github.com/antonmedv/expr"
	"github.com/antonmedv/expr/compiler"
	"github.com/antonmedv/expr/parser"
	"github.com/antonmedv/expr/vm"
	log "github.com/sirupsen/logrus"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type RangeMatch struct {
	BlackWhiteMatch

	RangeExpress string
	Script       string
	Begin        interface{}
	End          interface{}
	Program      *vm.Program
}

type RangeEntity struct {
	beginAli    string
	begin       interface{}
	end         interface{}
	endAli      string
	dateFlag    bool
	dynamicTime bool
}

type Predicate func(subCondition string) bool

func (rangeMatch *RangeMatch) Match(object interface{}, field reflect.StructField, fieldValue interface{}) bool {
	env := map[string]interface{}{
		"begin": rangeMatch.Begin,
		"end":   rangeMatch.End,
	}

	fieldKind := field.Type.Kind()
	if util.IsNumber(fieldKind) {
		env["value"] = fieldValue
	} else if fieldKind == reflect.String {
		env["value"] = len(fmt.Sprintf("%v", fieldValue))
	} else if fieldKind == reflect.Slice {
		env["value"] = reflect.ValueOf(fieldValue).Len()
	} else if field.Type.String() == "time.Time" {
		env["value"] = fieldValue.(time.Time).UnixNano()
	} else {
		// todo 如果是时间类型
	}

	output, err := expr.Run(rangeMatch.Program, env)
	if err != nil {
		log.Errorf("脚本 %v 执行失败: %v", rangeMatch.Script, err.Error())
		return false
	}

	result, err := util.CastBool(fmt.Sprintf("%v", output))
	if err != nil {
		return false
	}

	if result {
		if field.Type.Kind() == reflect.String {
			if len(fmt.Sprintf("%v", fieldValue)) > 1024 {
				rangeMatch.SetBlackMsg("属性 %v 值的字符串长度位于禁用的范围 %v 中", field.Name, fieldValue, rangeMatch.RangeExpress)
			} else {
				rangeMatch.SetBlackMsg("属性 %v 值 %v 的字符串长度位于禁用的范围 %v 中", field.Name, fieldValue, rangeMatch.RangeExpress)
			}
		} else if util.IsNumber(field.Type.Kind()) {
			rangeMatch.SetBlackMsg("属性 %v 值 %v 位于禁用的范围 %v 中", field.Name, fieldValue, rangeMatch.RangeExpress)
		} else if field.Type.Kind() == reflect.Slice {
			if reflect.ValueOf(fieldValue).Len() > 1024 {
				rangeMatch.SetBlackMsg("属性 %v 值的数组长度位于禁用的范围 %v 中", field.Name, fieldValue, rangeMatch.RangeExpress)
			} else {
				rangeMatch.SetBlackMsg("属性 %v 值 %v 的数组长度位于禁用的范围 %v 中", field.Name, fieldValue, rangeMatch.RangeExpress)
			}
		} else if field.Type.String() == "time.Time" {
			rangeMatch.SetBlackMsg("属性 %v 时间 %v 位于禁用时间段 %v 中", field.Name, fieldValue, rangeMatch.RangeExpress)
		} else {
			// todo
		}
		return true
	} else {
		if field.Type.Kind() == reflect.String {
			if len(fmt.Sprintf("%v", fieldValue)) > 1024 {
				rangeMatch.SetWhiteMsg("属性 %v 值的长度没有命中只允许的范围 %v", field.Name, fieldValue, rangeMatch.RangeExpress)
			} else {
				rangeMatch.SetWhiteMsg("属性 %v 值 %v 的长度没有命中只允许的范围 %v", field.Name, fieldValue, rangeMatch.RangeExpress)
			}
		} else if util.IsNumber(field.Type.Kind()) {
			rangeMatch.SetWhiteMsg("属性 %v 值 %v 没有命中只允许的范围 %v", field.Name, fieldValue, rangeMatch.RangeExpress)
		} else if field.Type.Kind() == reflect.Slice {
			if reflect.ValueOf(fieldValue).Len() > 1024 {
				rangeMatch.SetWhiteMsg("属性 %v 值的数组长度没有命中只允许的范围 %v", field.Name, fieldValue, rangeMatch.RangeExpress)
			} else {
				rangeMatch.SetWhiteMsg("属性 %v 值 %v 的数组长度没有命中只允许的范围 %v", field.Name, fieldValue, rangeMatch.RangeExpress)
			}
		} else if field.Type.String() == "time.Time" {
			rangeMatch.SetBlackMsg("属性 %v 时间 %v 的数组长度没有命中只允许的范围 %v 中", field.Name, fieldValue, rangeMatch.RangeExpress)
		} else {
			// todo
		}
		return false
	}
}

func (rangeMatch *RangeMatch) IsEmpty() bool {
	return rangeMatch.Script == ""
}

/*
 * []：时间或者数字范围匹配
 */
var rangeRegex = regexp.MustCompile("^(\\(|\\[)(.*)(,|，)(\\s)*(.*)(\\)|\\])$")

// digitRegex 全是数字匹配（整数，浮点数，0，负数）
var digitRegex = regexp.MustCompile("^[-+]?(0)|([1-9]+\\d*|0\\.(\\d*)|[1-9]\\d*\\.(\\d*))$")

// 时间的前后计算匹配：(-|+)yMd(h|H)msS
var timePlusRegex = regexp.MustCompile("^([-+])?(\\d*y)?(\\d*M)?(\\d*d)?(\\d*H|\\d*h)?(\\d*m)?(\\d*s)?$")

func BuildRangeMatcher(objectTypeFullName string, fieldKind reflect.Kind, objectFieldName string, subCondition string) {
	if !strings.Contains(subCondition, constant.RANGE) || !strings.Contains(subCondition, constant.EQUAL) {
		return
	}

	index := strings.Index(subCondition, "=")
	value := subCondition[index+1:]

	rangeEntity := parseRange(fieldKind, value)
	if rangeEntity == nil {
		return
	}

	beginAli := rangeEntity.beginAli
	begin := rangeEntity.begin
	end := rangeEntity.end
	endAli := rangeEntity.endAli

	var script string
	if begin == nil {
		if end == nil {
			return
		} else {
			if constant.RIGHT_EQUAL == endAli {
				script = "value <= end"
			} else if constant.RIGHT_UN_EQUAL == endAli {
				script = "value < end"
			}
		}
	} else {
		if end == nil {
			if constant.LEFT_EQUAL == beginAli {
				script = "begin <= value"
			} else if constant.LEFT_UN_EQUAL == beginAli {
				script = "begin < value"
			}
		} else {
			if constant.LEFT_EQUAL == beginAli && constant.RIGHT_EQUAL == endAli {
				script = "begin <= value && value <= end"
			} else if constant.LEFT_EQUAL == beginAli && constant.RIGHT_UN_EQUAL == endAli {
				script = "begin <= value && value < end"
			} else if constant.LEFT_UN_EQUAL == beginAli && constant.RIGHT_EQUAL == endAli {
				script = "begin < value && value <= end"
			} else if constant.LEFT_UN_EQUAL == beginAli && constant.RIGHT_UN_EQUAL == endAli {
				script = "begin < value && value < end"
			}
		}
	}

	tree, err := parser.Parse(script)
	if err != nil {
		log.Errorf("脚本：%v 解析异常：%v", script, err.Error())
		return
	}

	program, err := compiler.Compile(tree, nil)
	if err != nil {
		log.Errorf("脚本: %v 编译异常：%v", err.Error())
		return
	}

	addMatcher(objectTypeFullName, objectFieldName, &RangeMatch{Program: program, Begin: begin, End: end, Script: script, RangeExpress: value})
}

func parseRange(fieldKind reflect.Kind, subCondition string) *RangeEntity {
	subData := rangeRegex.FindAllStringSubmatch(subCondition, 1)
	if len(subData) > 0 {
		beginAli := subData[0][1]
		begin := subData[0][2]
		end := subData[0][5]
		endAli := subData[0][6]

		if (begin == "nil" || begin == "") && (end == "nil" || end == "") {
			log.Errorf("range匹配器格式输入错误，start和end不可都为null或者空字符, input=%v", subCondition)
			return nil
		} else if begin == "past" || begin == "future" {
			log.Errorf("range匹配器格式输入错误, start不可含有past或者future, input=%v", subCondition)
			return nil
		} else if end == "past" || end == "future" {
			log.Errorf("range匹配器格式输入错误, end不可含有past或者future, input=%v", subCondition)
			return nil
		}

		// 如果是数字，则按照数字解析
		if digitRegex.MatchString(begin) || digitRegex.MatchString(end) {
			// todo 添加begin要小于end的校验
			return &RangeEntity{beginAli: beginAli, begin: parseNum(fieldKind, begin), end: parseNum(fieldKind, end), endAli: endAli, dateFlag: true}
		} else if timePlusRegex.MatchString(begin) || timePlusRegex.MatchString(end) {
			// 解析动态时间 todo
		} else {
			beginTime := util.ParseTime(begin)
			endTime := util.ParseTime(end)

			beginTimeIsEmpty := util.IsTimeEmpty(beginTime)
			endTimeIsEmpty := util.IsTimeEmpty(endTime)

			if !beginTimeIsEmpty && !endTimeIsEmpty {
				if beginTime.After(endTime) {
					log.Errorf("时间的范围起始点不正确，起点时间不应该大于终点时间")
					return nil
				}
				return &RangeEntity{beginAli: beginAli, begin: beginTime.UnixNano(), end: endTime.UnixNano(), endAli: endAli, dateFlag: true}
			} else if beginTimeIsEmpty && endTimeIsEmpty {
				log.Errorf("range 匹配器格式输入错误，解析数字或者日期失败, time: %v", subData)
			} else {
				if !beginTimeIsEmpty {
					return &RangeEntity{beginAli: beginAli, begin: beginTime.UnixNano(), end: 0, endAli: endAli, dateFlag: true}
				} else if !endTimeIsEmpty {
					return &RangeEntity{beginAli: beginAli, begin: 0, end: endTime.UnixNano(), endAli: endAli, dateFlag: true}
				} else {
					return nil
				}
			}
		}
	} else {
		// 匹配过去和未来的时间
		if subCondition == constant.PAST {
			// 过去，则范围为(null, now)
			return &RangeEntity{beginAli: constant.LEFT_UN_EQUAL, begin: nil, end: constant.NOW, endAli: constant.RIGHT_UN_EQUAL, dateFlag: true}
		} else if subCondition == constant.FUTURE {
			// 未来，则范围为(now, null)
			return &RangeEntity{beginAli: constant.LEFT_UN_EQUAL, begin: constant.NOW, end: nil, endAli: constant.RIGHT_UN_EQUAL, dateFlag: true}
		}
		return nil
	}
	return nil
}

func parseNum(fieldKind reflect.Kind, valueStr string) interface{} {
	if util.IsNumber(fieldKind) {
		result, err := util.Cast(fieldKind, valueStr)
		if err != nil {
			return nil
		}
		return result
	} else if fieldKind == reflect.String || fieldKind == reflect.Slice {
		result, err := strconv.Atoi(valueStr)
		if err != nil {
			return nil
		}
		return result
	} else {
		return nil
	}
}
