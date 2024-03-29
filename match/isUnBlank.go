package matcher

import (
	"github.com/simonalong/mikilin-go/constant"
	log "github.com/sirupsen/logrus"
	"reflect"
	"strconv"
	"strings"
)

type IsUnBlankMatch struct {
	BlackWhiteMatch

	// 是否设置过isNil值
	HaveSet int8

	// 匹配非空的值
	IsUnBlank bool
}

func (isUnBlankMatch *IsUnBlankMatch) Match(parameterMap map[string]interface{}, object interface{}, field reflect.StructField, fieldValue interface{}) bool {
	if reflect.TypeOf(fieldValue).Kind() != field.Type.Kind() {
		isUnBlankMatch.SetBlackMsg("属性 %v 的值不是字符类型", field.Name)
		return false
	}

	if isUnBlankMatch.IsUnBlank {
		if fieldValue != "" {
			isUnBlankMatch.SetBlackMsg("属性 %v 的值为空字符", field.Name)
			return true
		} else {
			isUnBlankMatch.SetWhiteMsg("属性 %v 的值为非空字符", field.Name)
			return false
		}
	} else {
		if fieldValue == "" {
			isUnBlankMatch.SetBlackMsg("属性 %v 的值不为空", field.Name)
			return true
		} else {
			isUnBlankMatch.SetWhiteMsg("属性 %v 的值为空字符", field.Name)
			return false
		}
	}
}

func (isUnBlankMatch *IsUnBlankMatch) IsEmpty() bool {
	return isUnBlankMatch.HaveSet == 0
}

func BuildIsUnBlankMatcher(objectTypeFullName string, fieldKind reflect.Kind, objectFieldName string, tagName string, subCondition string, errMsg string) {
	if constant.MATCH != tagName {
		return
	}

	if fieldKind == reflect.Slice {
		return
	}

	if !strings.Contains(subCondition, constant.IsUnBlank) {
		return
	}

	value := "true"
	if strings.Contains(subCondition, constant.EQUAL) {
		index := strings.Index(subCondition, "=")
		value = strings.TrimSpace(subCondition[index+1:])
	}

	if strings.EqualFold("true", value) || strings.EqualFold("false", value) {
		var isUnBlank bool
		if chgValue, err := strconv.ParseBool(value); err == nil {
			isUnBlank = chgValue
		} else {
			log.Error(err.Error())
			return
		}
		addMatcher(objectTypeFullName, objectFieldName, &IsUnBlankMatch{IsUnBlank: isUnBlank, HaveSet: 1}, errMsg, true)
	}
}
