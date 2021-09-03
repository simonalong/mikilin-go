package matcher

import (
	"github.com/SimonAlong/Mikilin-go/util"
	log "github.com/sirupsen/logrus"
	"reflect"
	"strings"
)

type ValueMatch struct {
	BlackWhiteMatch
	Values []interface{}
}

func (valueMatch *ValueMatch) Match(object interface{}, field reflect.StructField, fieldValue interface{}) bool {
	values := valueMatch.Values

	for _, value := range values {
		if value == fieldValue {
			valueMatch.SetBlackMsg("属性 %v 的值 %v 位于禁用值 %v 中", field.Name, fieldValue, values)
			return true
		}
	}
	valueMatch.SetWhiteMsg("属性 %v 的值 %v 不在只可用列表 %v 中", field.Name, fieldValue, values)
	return false
}

func (valueMatch *ValueMatch) IsEmpty() bool {
	return len(valueMatch.Values) == 0
}

func BuildValuesMatcher(objectTypeFullName string, fieldKind reflect.Kind, objectFieldName string, subCondition string) {
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
			if chgValue, err := util.Cast(fieldKind, subValue); err == nil {
				availableValues = append(availableValues, chgValue)
			} else {
				log.Error(err.Error())
				continue
			}
		}
	} else {
		value = strings.TrimSpace(value)
		if chgValue, err := util.Cast(fieldKind, value); err == nil {
			availableValues = append(availableValues, chgValue)
		} else {
			log.Error(err.Error())
			return
		}
	}
	addMatcher(objectTypeFullName, objectFieldName, &ValueMatch{Values: availableValues})
}
