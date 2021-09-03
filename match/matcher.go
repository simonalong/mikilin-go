package matcher

import "reflect"

type Matcher interface {
	Match(object interface{}, field reflect.StructField, fieldValue interface{}) bool
	IsEmpty() bool
	GetWhitMsg() string
	GetBlackMsg() string
}

type FieldMatcher struct {

	// 属性名
	FieldName string
	// 异常名字
	ErrMsg string
	// 是否接受：true，则表示白名单，false，则表示黑名单
	Accept bool
	// 是否禁用
	Disable bool
	// 待转换的名字
	ChangeTo string
	// 匹配器列表
	Matchers []*Matcher
}

/* key：类全名，value：key：属性名 */
var MatchMap = make(map[string]map[string]*FieldMatcher)