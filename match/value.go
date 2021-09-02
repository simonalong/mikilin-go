package matcher

import (
	"reflect"
)

type ValueMatch struct {
	BlackWhiteMatch
	Values []interface{}
}

func (valueMatch *ValueMatch) Match(object interface{}, field reflect.StructField, fieldValue interface{}) bool {
	values := valueMatch.Values

	for _, value := range values {
		if value == fieldValue {
			valueMatch.SetBlack("属性 %v 的值 %v 位于禁用值 %v 中", field.Name, fieldValue, values)
			return true
		}
	}
	valueMatch.SetWhite("属性 %v 的值 %v 不在只可用列表 %v 中", field.Name, fieldValue, values)
	return false
}

func (valueMatch *ValueMatch) IsEmpty() bool {
	return len(valueMatch.Values) == 0
}
