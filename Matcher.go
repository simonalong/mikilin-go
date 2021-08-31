package mikilin

import (
	"fmt"
	matcher "github.com/SimonAlong/Mikilin-go/match"
	"reflect"
	"strings"
)

type Matcher interface {
	Match(object interface{}, field reflect.StructField, fieldValue interface{}) bool
	IsEmpty() bool
	WhitMsg() string
	BlackMsg() string
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
	Matchers []Matcher
}

type InfoCollector func(objectTypeName string, objectFieldName string, subCondition string)

type CollectorEntity struct {
	name         string
	infCollector InfoCollector
}

var checkerEntities []CollectorEntity

/* key：类全名，value：key：属性名 */
var matcherMap = make(map[string]map[string]FieldMatcher)

func Check(object interface{}) bool {
	objType := reflect.TypeOf(object)
	objValue := reflect.ValueOf(object)
	fmt.Println(objType.String())
	for index, num := 0, objType.NumField(); index < num; index++ {
		field := objType.Field(index)

		tagJudge := field.Tag.Get(CHECK)
		if len(tagJudge) == 0 {
			continue
		}

		// 搜集核查器
		if _, contain := matcherMap[objType.String()][field.Name]; !contain {
			collectChecker(objType.String(), field.Name, tagJudge)
		}

		// 核查结果：任何一个属性失败，则返回失败
		matchResult := check(object, field, objValue.Field(index).Interface())
		if !matchResult {
			return false
		}
	}
	return true
}

func ErrMsg() string {
	// todo errMsg
	return "出错啦"
}

func ErrMsgMap() map[string]interface{} {
	// todo
	return nil
}

func collectChecker(objectName string, fieldName string, matchJudge string) {
	subCondition := strings.Split(matchJudge, ";")
	for _, subStr := range subCondition {
		subStr = strings.TrimSpace(subStr)
		buildChecker(objectName, fieldName, subStr)
	}
}

func buildChecker(objectName string, fieldName string, subStr string) {
	for _, entity := range checkerEntities {
		entity.infCollector(objectName, fieldName, subStr)
	}
}

func check(object interface{}, field reflect.StructField, fieldValue interface{}) bool {
	objectType := reflect.TypeOf(object)
	if fieldMatcher, contain := matcherMap[objectType.String()][field.Name]; contain {
		accept := fieldMatcher.accept
		matchers := fieldMatcher.Matchers
		for _, entity := range matchers {
			if entity.IsEmpty() {
				continue
			}

			matchResult := entity.Match(object, field, fieldValue)
			if accept {
				if !matchResult {
					// 白名单，没有匹配上则返回false
					return false
				}
			} else {
				if matchResult {
					// 黑名单，匹配上则返回false
					return false
				}
			}
		}
	}
	return true
}

// 包的初始回调
func init() {
	/* 搜集匹配后的操作参数 */
	checkerEntities = append(checkerEntities, CollectorEntity{ERR_MSG, collectErrMsg})
	checkerEntities = append(checkerEntities, CollectorEntity{CHANGE_TO, collectChangeTo})
	checkerEntities = append(checkerEntities, CollectorEntity{ACCEPT, collectAccept})
	checkerEntities = append(checkerEntities, CollectorEntity{DISABLE, collectDisable})

	/* 搜集匹配器 */
	checkerEntities = append(checkerEntities, CollectorEntity{VALUE, buildValuesMatcher})
	checkerEntities = append(checkerEntities, CollectorEntity{IS_NIL, buildIsNilMatcher})
	checkerEntities = append(checkerEntities, CollectorEntity{IS_BLANK, buildIsBlankMatcher})
	checkerEntities = append(checkerEntities, CollectorEntity{RANGE, buildRangeMatcher})
	checkerEntities = append(checkerEntities, CollectorEntity{MODEL, buildModelMatcher})
	checkerEntities = append(checkerEntities, CollectorEntity{ENUM_TYPE, buildEnumTypeMatcher})
	checkerEntities = append(checkerEntities, CollectorEntity{CONDITION, buildConditionMatcher})
	checkerEntities = append(checkerEntities, CollectorEntity{CUSTOMIZE, buildCustomizeMatcher})
	checkerEntities = append(checkerEntities, CollectorEntity{REGEX, buildRegexMatcher})
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
		var valueMatchers []interface{}
		for _, subValue := range strings.Split(value, ",") {
			subValue = strings.TrimSpace(subValue)
			valueMatchers = append(valueMatchers, subValue)
		}
		valueMatch := matcher.ValueMatch{Values: valueMatchers}

		var matchers []Matcher
		matchers = append(matchers, &valueMatch)

		/* 添加匹配器到map */
		fieldMatcher, contain := matcherMap[objectTypeName][objectFieldName]
		if !contain {
			matcherMap[objectTypeName] = make(map[string]FieldMatcher)
			fieldMatcher = FieldMatcher{fieldName: objectFieldName, Matchers: matchers, accept: true}
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
