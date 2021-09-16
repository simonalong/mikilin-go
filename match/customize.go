package matcher

//
//import (
//	"fmt"
//	"github.com/SimonAlong/Mikilin-go/constant"
//	"github.com/SimonAlong/Mikilin-go/util"
//	"github.com/antonmedv/expr"
//	"github.com/antonmedv/expr/compiler"
//	"github.com/antonmedv/expr/parser"
//	"github.com/antonmedv/expr/vm"
//	log "github.com/sirupsen/logrus"
//	"reflect"
//	"strings"
//)
//
//type CustomizeMatch struct {
//	BlackWhiteMatch
//
//	expression string
//	Program    *vm.Program
//}
//
//
//type MatchJudge func(interface{}) bool
//
//func (customizeMatch *CustomizeMatch) Match(object interface{}, field reflect.StructField, fieldValue interface{}) bool {
//	env := map[string]interface{}{
//		"root":    object,
//		"current": fieldValue,
//		// todo
//		customizeMatch.expression:
//	}
//
//	output, err := expr.Run(customizeMatch.Program, env)
//	if err != nil {
//		log.Errorf("函数 %v 执行失败: %v", customizeMatch.expression, err.Error())
//		return false
//	}
//
//	result, err := util.CastBool(fmt.Sprintf("%v", output))
//	if err != nil {
//		return false
//	}
//
//	if result {
//		customizeMatch.SetBlackMsg("属性 %v 的值 %v 命中禁用条件回调 [%v] ", field.Name, fieldValue, customizeMatch.expression)
//	} else {
//		customizeMatch.SetWhiteMsg("属性 %v 的值 %v 没命中只允许条件回调 [%v] ", field.Name, fieldValue, customizeMatch.expression)
//	}
//	return result;
//}
//
//func (customizeMatch *CustomizeMatch) IsEmpty() bool {
//	return true
//}
//
//func BuildCustomizeMatcher(objectTypeFullName string, fieldKind reflect.Kind, objectFieldName string, tagName string, subCondition string, errMsg string) {
//	if constant.MATCH != tagName {
//		return
//	}
//
//	if !strings.Contains(subCondition, constant.Customize) {
//		return
//	}
//
//	index := strings.Index(subCondition, "=")
//	expression := subCondition[index+1:]
//
//	if expression == "" {
//		return
//	}
//
//	// 替换"."为"_"
//	expression = strings.ReplaceAll(strings.TrimSpace(expression), ".", "_")
//
//	tree, err := parser.Parse(expression)
//	if err != nil {
//		log.Errorf("函数：%v 解析异常：%v", expression, err.Error())
//		return
//	}
//
//	program, err := compiler.Compile(tree, nil)
//	if err != nil {
//		log.Errorf("函数: %v 编译异常：%v", expression, err.Error())
//		return
//	}
//	addMatcher(objectTypeFullName, objectFieldName, &CustomizeMatch{Program: program, expression: expression}, errMsg, true)
//}
