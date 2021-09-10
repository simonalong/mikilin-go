package matcher

import (
	"fmt"
	"github.com/antonmedv/expr/compiler"
	"github.com/antonmedv/expr/parser"
	log "github.com/sirupsen/logrus"
	"strings"
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

func addMatcher(objectTypeFullName string, objectFieldName string, matcher Matcher, errMsg string, accept bool) {
	fieldMatcherMap, c1 := MatchMap[objectTypeFullName]

	if !c1 {
		fieldMap := make(map[string]*FieldMatcher)
		var matchers []*Matcher
		if matcher != nil {
			matchers = append(matchers, &matcher)
		}

		fieldMap[objectFieldName] = &FieldMatcher{FieldName: objectFieldName, ErrMsg: errMsg, Matchers: matchers, Accept: accept}
		MatchMap[objectTypeFullName] = fieldMap
	} else {
		fieldMatcher, c2 := fieldMatcherMap[objectFieldName]
		if !c2 {
			var matchers []*Matcher
			if matcher != nil {
				matchers = append(matchers, &matcher)
			}
			fieldMatcherMap[objectFieldName] = &FieldMatcher{FieldName: objectFieldName, ErrMsg: errMsg, Matchers: matchers, Accept: accept}
		} else {
			if matcher != nil {
				fieldMatcher.Matchers = append(fieldMatcher.Matchers, &matcher)
			}
			fieldMatcher.Accept = accept
		}
	}
}

func parseErrMsg(originalErrMsg string, object interface{}, fieldValue interface{}) string {
	// todo
}

// 将其中的root.xx和current生成对应的占位符和sprintf字段，比如：数据#root.Age的名字#current不合法，转换为：sprintf("数据%v的名字%v不合法", root.Age, current)
func errMsgToTemplate(errMsg string) string {

}

// 将#root和#current转换为root和#current，相当于移除井号
func rmvWell(expression string) string {
	if strings.Contains(expression, "#root.") {
		expression = strings.ReplaceAll(expression, "#root.", "root.")
	}

	if strings.Contains(expression, "#current") {
		expression = strings.ReplaceAll(expression, "#current", "current")
	}
	return expression
}
