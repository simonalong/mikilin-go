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
	if regexMatch.Reg.MatchString(fmt.Sprintf("%v", fieldValue)) {
		regexMatch.SetBlackMsg("属性 %v 的值 %v 命中禁用的正则表达式 %v ", field.Name, fieldValue, regexMatch.Reg.String())
		return true
	} else {
		regexMatch.SetWhiteMsg("属性 %v 的值 %v 没命中只允许的正则表达式 %v ", field.Name, fieldValue, regexMatch.Reg.String())
	}
	return false
}

func (regexMatch *RegexMatch) IsEmpty() bool {
	return regexMatch.Reg == nil
}
