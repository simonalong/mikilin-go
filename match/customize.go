package matcher

import (
	"github.com/simonalong/mikilin-go/constant"
	log "github.com/sirupsen/logrus"
	"reflect"
	"strings"
)

type CustomizeMatch struct {
	BlackWhiteMatch

	expression string
	funValue   reflect.Value
}

var funMap = make(map[string]interface{})

type MatchJudge func(interface{}) bool

func (customizeMatch *CustomizeMatch) Match(object interface{}, field reflect.StructField, fieldValue interface{}) bool {
	var in []reflect.Value
	if customizeMatch.funValue.Type().NumIn() == 1 {
		in = make([]reflect.Value, 1)
		in[0] = reflect.ValueOf(fieldValue)
	} else {
		in = make([]reflect.Value, 2)
		inKind0 := customizeMatch.funValue.Type().In(0).Kind()
		inKind1 := customizeMatch.funValue.Type().In(1).Kind()

		if inKind0 == reflect.ValueOf(object).Kind() && inKind1 == reflect.ValueOf(fieldValue).Kind() {
			in[0] = reflect.ValueOf(object)
			in[1] = reflect.ValueOf(fieldValue)
		} else if inKind0 == reflect.ValueOf(fieldValue).Kind() && inKind1 == reflect.ValueOf(object).Kind() {
			in[0] = reflect.ValueOf(fieldValue)
			in[1] = reflect.ValueOf(object)
		} else {
			log.Errorf("the value don't match parameter of fun")
			return false
		}
	}

	retValues := customizeMatch.funValue.Call(in)
	if len(retValues) == 1 {
		if retValues[0].Bool() {
			customizeMatch.SetBlackMsg("属性 %v 的值 %v 命中禁用条件回调 [%v] ", field.Name, fieldValue, customizeMatch.expression)
		} else {
			customizeMatch.SetWhiteMsg("属性 %v 的值 %v 没命中只允许条件回调 [%v] ", field.Name, fieldValue, customizeMatch.expression)
		}
		return retValues[0].Bool()
	} else {
		kind0 := retValues[0].Kind()
		kind1 := retValues[1].Kind()

		if kind0 == reflect.Bool {
			if retValues[0].Bool() {
				customizeMatch.SetBlackMsg(retValues[1].String())
			} else {
				customizeMatch.SetWhiteMsg(retValues[1].String())
			}
			return retValues[0].Bool()
		} else if kind1 == reflect.Bool {
			if retValues[1].Bool() {
				customizeMatch.SetBlackMsg(retValues[0].String())
			} else {
				customizeMatch.SetWhiteMsg(retValues[0].String())
			}
			return retValues[1].Bool()
		} else {
			return retValues[0].Bool()
		}
	}
}

func (customizeMatch *CustomizeMatch) IsEmpty() bool {
	return customizeMatch.expression == ""
}

func BuildCustomizeMatcher(objectTypeFullName string, fieldKind reflect.Kind, objectFieldName string, tagName string, subCondition string, errMsg string) {
	if constant.MATCH != tagName {
		return
	}

	if !strings.Contains(subCondition, constant.Customize) {
		return
	}

	index := strings.Index(subCondition, "=")
	expression := subCondition[index+1:]

	if expression == "" {
		return
	}

	fun, contain := funMap[expression]
	if !contain {
		log.Errorf("the name of fun not find, funName is [%v]", expression)
		return
	}
	addMatcher(objectTypeFullName, objectFieldName, &CustomizeMatch{funValue: reflect.ValueOf(fun), expression: expression}, errMsg, true)
}

func RegisterCustomize(funName string, fun interface{}) {
	funValue := reflect.ValueOf(fun)
	if funValue.Kind() != reflect.Func {
		log.Errorf("fun is not fun type")
		return
	}

	if funValue.Type().NumIn() > 2 {
		log.Errorf("the num of argument need to be less than or equal to 2")
		return
	}

	if funValue.Type().NumOut() > 2 {
		log.Errorf("the num of return need to be less than or equal to 2")
		return
	}

	if funValue.Type().NumOut() == 1 {
		if funValue.Type().Out(0).Kind() != reflect.Bool {
			log.Errorf("the type of return must be bool")
			return
		}
	} else {
		kind0 := funValue.Type().Out(0).Kind()
		kind1 := funValue.Type().Out(1).Kind()

		if kind0 != reflect.Bool && kind0 != reflect.String {
			log.Errorf("return type of customize's fun return must be bool or string")
			return
		}

		if kind1 != reflect.Bool && kind1 != reflect.String {
			log.Errorf("return type of customize's fun return must be bool or string")
			return
		}
	}

	funMap[funName] = fun
}
