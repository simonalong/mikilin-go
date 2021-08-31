package mikilin

import (
	"fmt"
	"reflect"
	"strings"
)

type Match interface {
	match(object interface{}, field reflect.StructField, fieldValue interface{}) bool
	//isEmpty() bool
	//whitMsg() string
	//blackMsg() string
}

type FieldMatcher struct {

	// 属性名
	fieldName string
	// 异常名字
	errMsg string
	// 是否禁用
	disable bool
	// 待转换的名字
	changeTo string
	// 匹配器列表
	Matchers []Match
}

type MatcherCollector func(objectTypeName string, objectFieldName string, subCondition string)

type MatcherUnit struct {
	name             string
	matcherCollector MatcherCollector
}

var matchers []MatcherUnit

/* key：类全名，value：key：属性名 */
var matcherMap = make(map[string]map[string]FieldMatcher)

func Check(object interface{}) (bool, error) {
	objType := reflect.TypeOf(object)
	objValue := reflect.ValueOf(object)
	fmt.Println(objType.String())
	for index, num := 0, objType.NumField(); index < num; index++ {
		field := objType.Field(index)

		matchJudge := field.Tag.Get(MATCHER)
		if len(matchJudge) == 0 {
			continue
		}

		if _, contain := matcherMap[objType.String()][field.Name]; !contain {
			collectMatcher(objType.String(), field.Name, matchJudge)
		}

		match(object, field, objValue.Field(index).Interface())
	}
	return true, nil
}

func collectMatcher(objectName string, fieldName string, matchJudge string) {

	subCondition := strings.Split(matchJudge, ";")
	for _, subStr := range subCondition {
		subStr = strings.TrimSpace(subStr)
		buildMatcher(objectName, fieldName, subStr)
	}
}

func buildMatcher(objectName string, fieldName string, subStr string) {
	for _, matcher := range matchers {
		matcher.matcherCollector(objectName, fieldName, subStr)
	}
}

func match(object interface{}, field reflect.StructField, fieldValue interface{}) {
	objectType := reflect.TypeOf(object)
	if fieldMatcher, contain := matcherMap[objectType.String()][field.Name]; contain {
		matchers := fieldMatcher.Matchers
		for _, matcher := range matchers {
			matcher.match(object, field, fieldValue)
		}
	}
}

type ValueMatch struct {
	values []interface{}
}

func (valueMatch *ValueMatch) match(object interface{}, field reflect.StructField, fieldValue interface{}) bool {
	values := valueMatch.values
	fmt.Println(fieldValue)
	fmt.Println(values)
	return true
}

/* 包的初始回调 */
func init() {
	/* 搜集匹配后的操作参数 */
	matchers = append(matchers, MatcherUnit{ERR_MSG, collectErrMsg})
	matchers = append(matchers, MatcherUnit{CHANGE_TO, collectChangeTo})
	matchers = append(matchers, MatcherUnit{ACCEPT, collectAccept})
	matchers = append(matchers, MatcherUnit{DISABLE, collectDisable})

	/* 构造匹配器 */
	matchers = append(matchers, MatcherUnit{VALUE, buildValuesMatcher})
	matchers = append(matchers, MatcherUnit{IS_NIL, buildIsNilMatcher})
	matchers = append(matchers, MatcherUnit{IS_BLANK, buildIsBlankMatcher})
	matchers = append(matchers, MatcherUnit{RANGE, buildRangeMatcher})
	matchers = append(matchers, MatcherUnit{MODEL, buildModelMatcher})
	matchers = append(matchers, MatcherUnit{ENUM_TYPE, buildEnumTypeMatcher})
	matchers = append(matchers, MatcherUnit{CONDITION, buildConditionMatcher})
	matchers = append(matchers, MatcherUnit{CUSTOMIZE, buildCustomizeMatcher})
	matchers = append(matchers, MatcherUnit{REGEX, buildRegexMatcher})
}

func collectErrMsg(objectTypeName string, objectFieldName string, subCondition string) {

}

func collectChangeTo(objectTypeName string, objectFieldName string, subCondition string) {

}

func collectAccept(objectTypeName string, objectFieldName string, subCondition string) {

}

func collectDisable(objectTypeName string, objectFieldName string, subCondition string) {

}

func buildValuesMatcher(objectTypeName string, objectFieldName string, subCondition string) {
	if !strings.Contains(subCondition, VALUE) || !strings.Contains(subCondition, EQUAL) {
		return
	}

	index := strings.Index(subCondition, "=")
	value := subCondition[index+1:]

	if strings.HasPrefix(value, "{") && strings.HasSuffix(value, "}") {
		value = value[1 : len(value)-1]
		valueMatchers := []interface{}{}
		for _, subValue := range strings.Split(value, ",") {
			subValue = strings.TrimSpace(subValue)
			valueMatchers = append(valueMatchers, subValue)
		}
		valueMatch := ValueMatch{values: valueMatchers}

		var matchers []Match
		matchers = append(matchers, &valueMatch)

		fieldMatcher, contain := matcherMap[objectTypeName][objectFieldName]
		if !contain {
			fieldMatcher = FieldMatcher{fieldName: objectFieldName, Matchers: matchers}
		} else {
			fieldMatcher.Matchers = append(fieldMatcher.Matchers, matchers...)
		}
		matcherMap[objectTypeName][objectFieldName] = fieldMatcher
	}
}

func buildIsNilMatcher(objectTypeName string, objectFieldName string, subCondition string) {

}

func buildIsBlankMatcher(objectTypeName string, objectFieldName string, subCondition string) {

}

func buildRangeMatcher(objectTypeName string, objectFieldName string, subCondition string) {

}

func buildModelMatcher(objectTypeName string, objectFieldName string, subCondition string) {

}

func buildEnumTypeMatcher(objectTypeName string, objectFieldName string, subCondition string) {

}

func buildConditionMatcher(objectTypeName string, objectFieldName string, subCondition string) {

}

func buildRegexMatcher(objectTypeName string, objectFieldName string, subCondition string) {

}

func buildCustomizeMatcher(objectTypeName string, objectFieldName string, subCondition string) {

}
