package mikilin

import (
	matcher "github.com/SimonAlong/Mikilin-go/match"
	"github.com/SimonAlong/Mikilin-go/util"
	log "github.com/sirupsen/logrus"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"unicode"
)

var lock sync.Mutex

type Matcher interface {
	Match(object interface{}, field reflect.StructField, fieldValue interface{}) bool
	IsEmpty() bool
	GetWhitMsg() string
	GetBlackMsg() string
}

type FieldMatcher struct {

	// 属性名
	fieldName string
	// 异常名字
	errMsg string
	// 是否接受：true，则表示白名单，false，则表示黑名单
	accept bool
	// 是否禁用
	disable bool
	// 待转换的名字
	changeTo string
	// 匹配器列表
	Matchers []*Matcher
}

type InfoCollector func(objectTypeFullName string, fieldKind reflect.Kind, objectFieldName string, subCondition string)

type CollectorEntity struct {
	name         string
	infCollector InfoCollector
}

type CheckResult struct {
	Result bool
	ErrMsg string
}

var checkerEntities []CollectorEntity

/* key：类全名，value：key：属性名 */
var matcherMap = make(map[string]map[string]*FieldMatcher)

/* 核查的标签 */
var matchTagArray = []string{VALUE, IsBlank, RANGE, MODEL, CONDITION, REGEX, CUSTOMIZE}

/* 匹配后处理的标签 */
var handleTagArray = []string{ERR_MSG, CHANGE_TO, ACCEPT, DISABLE}

func Check(object interface{}, fieldNames ...string) (bool, string) {
	if object == nil {
		return true, ""
	}
	objType := reflect.TypeOf(object)
	objValue := reflect.ValueOf(object)

	// 指针类型按照指针类型
	if objType.Kind() == reflect.Ptr {
		objValue = objValue.Elem()
		return Check(objValue.Interface(), fieldNames...)
	}

	if objValue.Kind() != reflect.Struct {
		return true, ""
	}

	// 搜集核查器
	collectCollector(objType)

	ch := make(chan *CheckResult)
	for index, num := 0, objType.NumField(); index < num; index++ {
		field := objType.Field(index)
		fieldValue := objValue.Field(index)

		// 私有字段不处理
		if !isStartUpper(field.Name) {
			continue
		}

		// 过滤选择的列
		if !isSelectField(field.Name, fieldNames...) {
			continue
		}

		if fieldValue.Kind() == reflect.Ptr && !fieldValue.IsNil() {
			fieldValue = fieldValue.Elem()
		}

		// 基本类型
		if isCheckedKing(fieldValue.Kind()) {
			tagJudge := field.Tag.Get(MATCH)
			if len(tagJudge) == 0 {
				continue
			}

			// 核查结果：任何一个属性失败，则返回失败
			go check(object, field, fieldValue.Interface(), ch)
			checkResult := <-ch
			if !checkResult.Result {
				close(ch)
				return false, checkResult.ErrMsg
			}
		} else if fieldValue.Kind() == reflect.Struct {
			// struct 结构类型
			tagMatch := field.Tag.Get(MATCH)
			if len(tagMatch) == 0 || (len(tagMatch) == 1 && tagMatch != CHECK) {
				continue
			}
			result, err := Check(fieldValue.Interface())
			if !result {
				return false, err
			}
		} else if fieldValue.Kind() == reflect.Map {
			// map结构
			if fieldValue.Len() == 0 {
				continue
			}

			for mapR := fieldValue.MapRange(); mapR.Next(); {
				mapKey := mapR.Key()
				mapValue := mapR.Value()

				result, err := Check(mapKey.Interface())
				if !result {
					return false, err
				}
				result, err = Check(mapValue.Interface())
				if !result {
					return false, err
				}
			}
		} else if fieldValue.Kind() == reflect.Array || fieldValue.Kind() == reflect.Slice {
			// Array|Slice 结构
			arrayLen := fieldValue.Len()
			for arrayIndex := 0; arrayIndex < arrayLen; arrayIndex++ {
				fieldValueItem := fieldValue.Index(arrayIndex)
				result, err := Check(fieldValueItem.Interface())
				if !result {
					return false, err
				}
			}
		}
	}
	close(ch)
	return true, ""
}

// 搜集核查器
func collectCollector(objType reflect.Type) {
	objectFullName := objType.String()

	/* 搜集过则不再搜集 */
	if _, contain := matcherMap[objectFullName]; contain {
		return
	}

	lock.Lock()
	/* 搜集过则不再搜集 */
	if _, contain := matcherMap[objectFullName]; contain {
		return
	}

	doCollectCollector(objType)
	lock.Unlock()
}

func doCollectCollector(objType reflect.Type) {
	// 基本类型不需要搜集
	if isCheckedKing(objType.Kind()) {
		return
	}

	// 指针类型按照指针类型
	if objType.Kind() == reflect.Ptr {
		doCollectCollector(objType.Elem())
		return
	}

	objectFullName := objType.String()
	for fieldIndex, num := 0, objType.NumField(); fieldIndex < num; fieldIndex++ {
		field := objType.Field(fieldIndex)
		fieldKind := field.Type.Kind()

		// 不可访问字段不处理
		if !isStartUpper(field.Name) {
			continue
		}

		if fieldKind == reflect.Ptr {
			fieldKind = field.Type.Elem().Kind()
		}

		// 基本类型
		if isCheckedKing(fieldKind) {
			tagMatch := field.Tag.Get(MATCH)
			if len(tagMatch) == 0 {
				continue
			}

			if _, contain := matcherMap[objectFullName][field.Name]; !contain {
				collectChecker(objectFullName, fieldKind, field.Name, tagMatch)
			}
		} else if fieldKind == reflect.Struct {
			// struct 结构类型
			tagMatch := field.Tag.Get(MATCH)
			if len(tagMatch) == 0 || (len(tagMatch) == 1 && tagMatch != CHECK) {
				continue
			}

			doCollectCollector(field.Type)
		} else if fieldKind == reflect.Map {
			// Map 结构
			doCollectCollector(field.Type.Key())
			doCollectCollector(field.Type.Elem())
		} else if fieldKind == reflect.Array || fieldKind == reflect.Slice {
			// Array|Slice 结构
			doCollectCollector(field.Type.Elem())
		} else {
			// Uintptr 类型不处理
		}
	}
}

// 判断首字母是否大写
func isStartUpper(s string) bool {
	return unicode.IsUpper([]rune(s)[0])
}

// 是否是选择的列，没有选择也认为是选择的
func isSelectField(fieldName string, fieldNames ...string) bool {
	if len(fieldNames) == 0 {
		return true
	}
	for _, name := range fieldNames {
		if strings.EqualFold(name, fieldName) {
			return true
		}
	}
	return false
}

// 搜集处理器，对于有一些空格的也进行单独处理
func collectChecker(objectFullName string, fieldKind reflect.Kind, fieldName string, matchJudge string) {
	var subStrIndexes []int
	for _, tag := range matchTagArray {
		index := strings.Index(matchJudge, tag)
		if index != -1 {
			subStrIndexes = append(subStrIndexes, index)
		}
	}
	sort.Ints(subStrIndexes)

	lastIndex := 0
	for _, subIndex := range subStrIndexes {
		if lastIndex == subIndex {
			continue
		}
		subJudgeStr := matchJudge[lastIndex:subIndex]
		subJudgeStr = strings.Replace(subJudgeStr, " ", "", -1)
		buildChecker(objectFullName, fieldKind, fieldName, subJudgeStr)
		lastIndex = subIndex
	}

	subJudgeStr := matchJudge[lastIndex:]
	subJudgeStr = strings.Replace(subJudgeStr, " ", "", -1)
	buildChecker(objectFullName, fieldKind, fieldName, subJudgeStr)
}

func buildChecker(objectFullName string, fieldKind reflect.Kind, fieldName string, subStr string) {
	for _, entity := range checkerEntities {
		entity.infCollector(objectFullName, fieldKind, fieldName, subStr)
	}
}

func check(object interface{}, field reflect.StructField, fieldValue interface{}, ch chan *CheckResult) {
	objectType := reflect.TypeOf(object)

	if fieldMatcher, contain := matcherMap[objectType.String()][field.Name]; contain {
		accept := fieldMatcher.accept
		matchers := fieldMatcher.Matchers

		// 黑名单，而且匹配到，则核查失败
		if !accept {
			if matchResult, errMsg := judgeMatch(matchers, object, field, fieldValue, accept); matchResult {
				ch <- &CheckResult{Result: false, ErrMsg: errMsg}
				return
			}
		}

		// 白名单，没有匹配到，则核查失败
		if accept {
			if matchResult, errMsg := judgeMatch(matchers, object, field, fieldValue, accept); !matchResult {
				ch <- &CheckResult{Result: false, ErrMsg: errMsg}
				return
			}
		}
	}
	ch <- &CheckResult{Result: true}
	return
}

// 任何一个匹配上，则返回true，都没有匹配上则返回false
func judgeMatch(matchers []*Matcher, object interface{}, field reflect.StructField, fieldValue interface{}, accept bool) (bool, string) {
	var errMsgArray []string
	for _, match := range matchers {
		if (*match).IsEmpty() {
			continue
		}

		matchResult := (*match).Match(object, field, fieldValue)
		if matchResult {
			if !accept {
				errMsgArray = append(errMsgArray, (*match).GetWhitMsg())
			} else {
				errMsgArray = []string{}
			}
			return true, ""
		} else {
			if accept {
				errMsgArray = append(errMsgArray, (*match).GetWhitMsg())
			}
		}
	}
	return false, util.ArraysToString(errMsgArray)
}

// 包的初始回调
func init() {
	/* 搜集匹配后的操作参数 */
	//checkerEntities = append(checkerEntities, CollectorEntity{ERR_MSG, collectErrMsg})
	//checkerEntities = append(checkerEntities, CollectorEntity{CHANGE_TO, collectChangeTo})
	//checkerEntities = append(checkerEntities, CollectorEntity{ACCEPT, collectAccept})
	//checkerEntities = append(checkerEntities, CollectorEntity{DISABLE, collectDisable})

	/* 搜集匹配器 */
	checkerEntities = append(checkerEntities, CollectorEntity{VALUE, buildValuesMatcher})
	checkerEntities = append(checkerEntities, CollectorEntity{IsBlank, buildIsBlankMatcher})
	checkerEntities = append(checkerEntities, CollectorEntity{RANGE, buildRangeMatcher})
	//checkerEntities = append(checkerEntities, CollectorEntity{MODEL, buildModelMatcher})
	//checkerEntities = append(checkerEntities, CollectorEntity{ENUM_TYPE, buildEnumTypeMatcher})
	//checkerEntities = append(checkerEntities, CollectorEntity{CONDITION, buildConditionMatcher})
	//checkerEntities = append(checkerEntities, CollectorEntity{CUSTOMIZE, buildCustomizeMatcher})
	checkerEntities = append(checkerEntities, CollectorEntity{REGEX, buildRegexMatcher})
}

func collectErrMsg(objectTypeFullName string, fieldKind reflect.Kind, objectFieldName string, subCondition string) {

}

func collectChangeTo(objectTypeFullName string, fieldKind reflect.Kind, objectFieldName string, subCondition string) {

}

func collectAccept(objectTypeFullName string, fieldKind reflect.Kind, objectFieldName string, subCondition string) {

}

func collectDisable(objectTypeFullName string, fieldKind reflect.Kind, objectFieldName string, subCondition string) {

}

func buildValuesMatcher(objectTypeFullName string, fieldKind reflect.Kind, objectFieldName string, subCondition string) {
	if !strings.Contains(subCondition, VALUE) || !strings.Contains(subCondition, EQUAL) {
		return
	}

	index := strings.Index(subCondition, "=")
	value := subCondition[index+1:]

	var availableValues []interface{}
	if strings.HasPrefix(value, "{") && strings.HasSuffix(value, "}") {
		value = value[1 : len(value)-1]
		for _, subValue := range strings.Split(value, ",") {
			subValue = strings.TrimSpace(subValue)
			if chgValue, err := cast(fieldKind, subValue); err == nil {
				availableValues = append(availableValues, chgValue)
			} else {
				log.Error(err.Error())
				continue
			}
		}
	} else {
		value = strings.TrimSpace(value)
		if chgValue, err := cast(fieldKind, value); err == nil {
			availableValues = append(availableValues, chgValue)
		} else {
			log.Error(err.Error())
			return
		}
	}
	addMatcher(objectTypeFullName, objectFieldName, &matcher.ValueMatch{Values: availableValues})
}

func buildIsBlankMatcher(objectTypeFullName string, fieldKind reflect.Kind, objectFieldName string, subCondition string) {
	if !strings.Contains(subCondition, IsBlank) {
		return
	}

	value := "true"
	if strings.Contains(subCondition, EQUAL) {
		index := strings.Index(subCondition, "=")
		value = strings.TrimSpace(subCondition[index+1:])
	}

	if strings.EqualFold("true", value) || strings.EqualFold("false", value) {
		var isBlank bool
		if chgValue, err := strconv.ParseBool(value); err == nil {
			isBlank = chgValue
		} else {
			log.Error(err.Error())
			return
		}
		addMatcher(objectTypeFullName, objectFieldName, &matcher.IsBlankMatch{IsBlank: isBlank, HaveSet: 1})
	}
}

func buildRangeMatcher(objectTypeFullName string, fieldKind reflect.Kind, objectFieldName string, subCondition string) {
	// todo
}

func buildModelMatcher(objectTypeFullName string, fieldKind reflect.Kind, objectFieldName string, subCondition string) {

}

func buildEnumTypeMatcher(objectTypeFullName string, fieldKind reflect.Kind, objectFieldName string, subCondition string) {

}

func buildConditionMatcher(objectTypeFullName string, fieldKind reflect.Kind, objectFieldName string, subCondition string) {

}

func buildRegexMatcher(objectTypeFullName string, fieldKind reflect.Kind, objectFieldName string, subCondition string) {
	if !strings.Contains(subCondition, REGEX) || !strings.Contains(subCondition, EQUAL) {
		return
	}

	index := strings.Index(subCondition, "=")
	value := subCondition[index+1:]

	reg, err := regexp.Compile(value)
	if err != nil {
		return
	}
	addMatcher(objectTypeFullName, objectFieldName, &matcher.RegexMatch{Reg: reg})
}

func buildCustomizeMatcher(objectTypeFullName string, fieldKind reflect.Kind, objectFieldName string, subCondition string) {

}

func addMatcher(objectTypeFullName string, objectFieldName string, matcher Matcher) {
	// 添加匹配器到map
	fieldMatcherMap, c1 := matcherMap[objectTypeFullName]
	if !c1 {
		fieldMap := make(map[string]*FieldMatcher)

		var matchers []*Matcher
		matchers = append(matchers, &matcher)

		fieldMap[objectFieldName] = &FieldMatcher{fieldName: objectFieldName, Matchers: matchers, accept: true}
		matcherMap[objectTypeFullName] = fieldMap
	} else {
		fieldMatcher, c2 := fieldMatcherMap[objectFieldName]
		if !c2 {
			var matchers []*Matcher
			matchers = append(matchers, &matcher)

			fieldMatcherMap[objectFieldName] = &FieldMatcher{fieldName: objectFieldName, Matchers: matchers, accept: true}
		} else {
			fieldMatcher.Matchers = append(fieldMatcher.Matchers, &matcher)
		}
	}
}

// 判断是否是核查的类型
func isCheckedKing(fieldKing reflect.Kind) bool {
	switch fieldKing {
	case reflect.Int:
		return true
	case reflect.Int8:
		return true
	case reflect.Int16:
		return true
	case reflect.Int32:
		return true
	case reflect.Int64:
		return true
	case reflect.Uint:
		return true
	case reflect.Uint8:
		return true
	case reflect.Uint16:
		return true
	case reflect.Uint32:
		return true
	case reflect.Uint64:
		return true
	case reflect.Float32:
		return true
	case reflect.Float64:
		return true
	case reflect.Bool:
		return true
	case reflect.String:
		return true
	default:
		return false
	}
}

func cast(fieldKind reflect.Kind, valueStr string) (interface{}, error) {
	switch fieldKind {
	case reflect.Int:
		return strconv.Atoi(valueStr)
	case reflect.Int8:
		return strconv.ParseInt(valueStr, 10, 8)
	case reflect.Int16:
		return strconv.ParseInt(valueStr, 10, 16)
	case reflect.Int32:
		return strconv.ParseInt(valueStr, 10, 32)
	case reflect.Int64:
		return strconv.ParseInt(valueStr, 10, 64)
	case reflect.Uint:
		return strconv.ParseUint(valueStr, 10, 0)
	case reflect.Uint8:
		return strconv.ParseUint(valueStr, 10, 8)
	case reflect.Uint16:
		return strconv.ParseUint(valueStr, 10, 16)
	case reflect.Uint32:
		return strconv.ParseUint(valueStr, 10, 32)
	case reflect.Uint64:
		return strconv.ParseUint(valueStr, 10, 64)
	case reflect.Float32:
		return strconv.ParseFloat(valueStr, 32)
	case reflect.Float64:
		return strconv.ParseFloat(valueStr, 64)
	case reflect.Bool:
		return strconv.ParseBool(valueStr)
	}
	return valueStr, nil
}