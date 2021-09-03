package matcher

import (
	"fmt"
	"reflect"
	"regexp"
)

type RegexMatch struct {
	BlackWhiteMatch

	Reg *regexp.Regexp
}

func (regexMatch *RegexMatch) Match(object interface{}, field reflect.StructField, fieldValue interface{}) bool {
	//if reflect.ValueOf(fieldValue).Kind() == reflect.String {
	//	if regexMatch.Reg.MatchString(fmt.Sprintf("%s", fieldValue)) {
	//		regexMatch.SetBlackString("属性 %v 的值 %v 命中禁用的正则表达式 %s ", field.Name, fieldValue, regexMatch.Reg.String())
	//		return true
	//	} else {
	//		regexMatch.SetWhiteString("属性 %s 的值 %v 没命中只允许的正则表达式 %s ", []byte(field.Name), fieldValue, regexMatch.Reg.String())
	//	}
	//} else {
	//	regexMatch.SetWhiteString("属性 %v 的值 %v 不是String类型", field.Name, fieldValue)
	//}

	if regexMatch.Reg.MatchString(fmt.Sprintf("%v", fieldValue)) {
		regexMatch.SetBlackString("属性 %v 的值 %v 命中禁用的正则表达式 %s ", field.Name, fieldValue, regexMatch.Reg.String())
		return true
	} else {
		regexMatch.SetWhiteString("属性 %s 的值 %v 没命中只允许的正则表达式 %s ", []byte(field.Name), fieldValue, regexMatch.Reg.String())
	}
	return false
}

func (regexMatch *RegexMatch) IsEmpty() bool {
	return regexMatch.Reg == nil
}
