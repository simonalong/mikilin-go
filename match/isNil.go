package matcher

import (
	"reflect"
)

type NilMatch struct {
	BlackWhiteMatch

	// 是否设置过isNil值
	HaveSet int8

	// 修饰的值是否匹配isNil
	IsNil bool
}

func (isNilMatch *NilMatch) Match(object interface{}, field reflect.StructField, fieldValue interface{}) bool {
	if isNilMatch.IsNil {
		if fieldValue == nil {
			//isNilMatch.SetBlack("属性 %v 的值为null", field.Name)
			return true
		} else {
			//isNilMatch.SetWhite("属性 %v 的值为null", field.Name)
			return false
		}
	} else {
		if fieldValue != nil {
			//isNilMatch.SetBlack("属性 %v 的值不为null", field.Name)
			return true
		} else {
			//isNilMatch.SetWhite("属性 %v 的值不为null", field.Name)
			return false
		}
	}
}

func (isNilMatch *NilMatch) IsEmpty() bool {
	return isNilMatch.HaveSet != 0
}
