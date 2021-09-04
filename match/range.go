package matcher

import (
	"fmt"
	"github.com/SimonAlong/Mikilin-go/util"
	"github.com/antonmedv/expr"
	"github.com/antonmedv/expr/compiler"
	"github.com/antonmedv/expr/parser"
	"github.com/antonmedv/expr/vm"
	log "github.com/sirupsen/logrus"
	"reflect"
	"regexp"
	"strings"
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
		"o":     fieldValue,
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
		rangeMatch.SetBlackMsg("属性 %v 的 %v 位于禁用的范围 %v 中", field.Name, fieldValue, rangeMatch.RangeExpress)
		return true
	} else {
		rangeMatch.SetWhiteMsg("属性 %v 的 %v 没有命中只允许的范围 %v", field.Name, fieldValue, rangeMatch.RangeExpress)
		return false
	}
}

func (rangeMatch *RangeMatch) IsEmpty() bool {
	return rangeMatch.Script == ""
}

/*
 * []：范围匹配
 */
var rangeRegex = regexp.MustCompile("^(\\(|\\[)(.*)(,|，)(\\s)*(.*)(\\)|\\])$")

// digitRegex 全是数字匹配（整数，浮点数，0，负数）
var digitRegex = regexp.MustCompile("^[-+]?(0)|([1-9]+\\d*|0\\.(\\d*)|[1-9]\\d*\\.(\\d*))$")

/*
 * 时间或者数字范围匹配
 */
//private Pattern rangePattern = Pattern.compile("^(\\(|\\[)(.*),(\\s)*(.*)(\\)|\\])$");

func BuildRangeMatcher(objectTypeFullName string, fieldKind reflect.Kind, objectFieldName string, subCondition string) {
	if !strings.Contains(subCondition, RANGE) || !strings.Contains(subCondition, EQUAL) {
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
			if RIGHT_EQUAL == endAli {
				script = "o <= end"
			} else if RIGHT_UN_EQUAL == endAli {
				script = "o < end"
			}
		}
	} else {
		if end == nil {
			if LEFT_EQUAL == beginAli {
				script = "begin <= o"
			} else if LEFT_UN_EQUAL == beginAli {
				script = "begin < o"
			}
		} else {
			if LEFT_EQUAL == beginAli && RIGHT_EQUAL == endAli {
				script = "begin <= o && o <= end"
			} else if LEFT_EQUAL == beginAli && RIGHT_UN_EQUAL == endAli {
				script = "begin <= o && o < end"
			} else if LEFT_UN_EQUAL == beginAli && RIGHT_EQUAL == endAli {
				script = "begin < o && o <= end"
			} else if LEFT_UN_EQUAL == beginAli && RIGHT_UN_EQUAL == endAli {
				script = "begin < o && o < end"
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
	if subCondition == "[1, 2]" {
		fmt.Println("ok")
	}
	subData := rangeRegex.FindAllStringSubmatch(subCondition, 1)
	//subData := rangeRegex.FindAllStringSubmatch("[1, 2]", 1)
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
			return &RangeEntity{beginAli: beginAli, begin: parseNum(fieldKind, begin), end: parseNum(fieldKind, end), endAli: endAli}
		}

		// 如果是解析动态时间，按照动态时间解析

		// 否则按照时间解析

	} else {
		// 匹配过去和未来的时间

		//if (input.equals(PAST) || input.equals(FUTURE)) {
		//	return parseRangeDate(input);
		//}
	}

	//// 如果是数字，则按照数字解析
	//if digitRegex.MatchString(begin) || digitRegex.MatchString(end) {
	//
	//}

	// 如果是解析动态时间，按照动态时间解析

	// 否则按照时间解析
	return nil
}

func parseNum(fieldKind reflect.Kind, valueStr string) interface{} {
	result, err := util.Cast(fieldKind, valueStr)
	if err != nil {
		return nil
	}
	return result
}
