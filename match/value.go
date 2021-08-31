package matcher

import (
	"fmt"
	"reflect"
)

type ValueMatch struct {
	Values []interface{}
}

func (valueMatch *ValueMatch) Match(object interface{}, field reflect.StructField, fieldValue interface{}) bool {
	values := valueMatch.Values
	fmt.Println(fieldValue)
	fmt.Println(values)

	for _, value := range values {
		if value == fieldValue {
			return true
		}
	}
	return false
}

func (valueMatch *ValueMatch) IsEmpty() bool {
	return len(valueMatch.Values) == 0
}

func (valueMatch *ValueMatch) WhitMsg() string {
	return ""
}

func (valueMatch *ValueMatch) BlackMsg() string {
	return ""
}
