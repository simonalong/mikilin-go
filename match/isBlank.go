package matcher

import "reflect"

type IsBlankMatch struct {
	BlackWhiteMatch

	// 是否设置过isNil值
	HaveSet int8

	// 修饰的值是否匹配isNil
	IsBlank bool
}

func (isBlankMatch *IsBlankMatch) Match(object interface{}, field reflect.StructField, fieldValue interface{}) bool {
	if reflect.TypeOf(fieldValue).Kind() != field.Type.Kind() {
		isBlankMatch.SetBlackStruct("属性 %v 的值不是字符类型", field.Name)
		return false
	}

	if isBlankMatch.IsBlank {
		if fieldValue == "" {
			isBlankMatch.SetBlackString("属性 %v 的值为空字符", field.Name)
			return true
		} else {
			isBlankMatch.SetWhiteString("属性 %v 的值为非空字符", field.Name)
			return false
		}
	} else {
		if fieldValue != "" {
			isBlankMatch.SetBlackString("属性 %v 的值不为空", field.Name)
			return true
		} else {
			isBlankMatch.SetWhiteString("属性 %v 的值为空字符", field.Name)
			return false
		}
	}
}

func (isBlankMatch *IsBlankMatch) IsEmpty() bool {
	return isBlankMatch.HaveSet == 0
}
