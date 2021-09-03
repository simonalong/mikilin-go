package matcher

import (
	"fmt"
)

type BlackWhiteMatch struct {
	BlackMsg string
	WhiteMsg string
}

func (blackWhiteMatch *BlackWhiteMatch) SetBlackMsg(format string, a ...interface{}) {
	blackWhiteMatch.BlackMsg = fmt.Sprintf("核查错误："+format, a...)
}

func (blackWhiteMatch *BlackWhiteMatch) SetWhiteMsg(format string, a ...interface{}) {
	blackWhiteMatch.WhiteMsg = fmt.Sprintf("核查错误："+format, a...)
}

func (blackWhiteMatch *BlackWhiteMatch) GetWhitMsg() string {
	return blackWhiteMatch.WhiteMsg
}

func (blackWhiteMatch *BlackWhiteMatch) GetBlackMsg() string {
	return blackWhiteMatch.BlackMsg
}

func addMatcher(objectTypeFullName string, objectFieldName string, matcher Matcher) {
	// 添加匹配器到map
	fieldMatcherMap, c1 := MatchMap[objectTypeFullName]
	if !c1 {
		fieldMap := make(map[string]*FieldMatcher)

		var matchers []*Matcher
		matchers = append(matchers, &matcher)

		fieldMap[objectFieldName] = &FieldMatcher{FieldName: objectFieldName, Matchers: matchers, Accept: true}
		MatchMap[objectTypeFullName] = fieldMap
	} else {
		fieldMatcher, c2 := fieldMatcherMap[objectFieldName]
		if !c2 {
			var matchers []*Matcher
			matchers = append(matchers, &matcher)

			fieldMatcherMap[objectFieldName] = &FieldMatcher{FieldName: objectFieldName, Matchers: matchers, Accept: true}
		} else {
			fieldMatcher.Matchers = append(fieldMatcher.Matchers, &matcher)
		}
	}
}
