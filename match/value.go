package matcher

import (
	"fmt"
	"reflect"
)

type ValueMatch struct {
	BlackWhiteMatch
	Values []interface{}
}

func (valueMatch *ValueMatch) Match(object interface{}, field reflect.StructField, fieldValue interface{}) bool {
	values := valueMatch.Values
	fmt.Println(fieldValue)
	fmt.Println(values)

	for _, value := range values {
		if value == fieldValue {
			valueMatch.SetBlack("属性 %s 的值 %s 位于禁用值 %v 中", field.Name, fieldValue, values)
			return true
		}
	}
	valueMatch.SetBlack("属性 %s 的值 %s 不在只可用列表 %v 中", field, fieldValue, values)
	return false
}

func (valueMatch *ValueMatch) IsEmpty() bool {
	return len(valueMatch.Values) == 0
}

func (valueMatch *ValueMatch) GetWhitMsg() string {
	return valueMatch.WhiteMsg
}

func (valueMatch *ValueMatch) GetBlackMsg() string {
	return valueMatch.BlackMsg
}
