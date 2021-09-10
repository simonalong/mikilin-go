package matcher

import (
	"fmt"
	"github.com/SimonAlong/Mikilin-go/constant"
	"github.com/SimonAlong/Mikilin-go/util"
	"github.com/antonmedv/expr"
	"github.com/antonmedv/expr/compiler"
	"github.com/antonmedv/expr/parser"
	"github.com/antonmedv/expr/vm"
	log "github.com/sirupsen/logrus"
	"reflect"
	"strings"
)

type ConditionMatch struct {
	BlackWhiteMatch

	expression string
	Program    *vm.Program
}

func (conditionMatch *ConditionMatch) Match(object interface{}, field reflect.StructField, fieldValue interface{}) bool {
	env := map[string]interface{}{
		"root":    object,
		"current": fieldValue,
	}

	output, err := expr.Run(conditionMatch.Program, env)
	if err != nil {
		log.Errorf("表达式 %v 执行失败: %v", conditionMatch.expression, err.Error())
		return false
	}

	result, err := util.CastBool(fmt.Sprintf("%v", output))
	if err != nil {
		return false
	}

	if result {
		conditionMatch.SetBlackMsg("属性 %v 的值 %v 命中禁用条件 [%v] ", field.Name, fieldValue, conditionMatch.expression)
		return true
	} else {
		conditionMatch.SetWhiteMsg("属性 %v 的值 %v 不符合条件 [%v] ", field.Name, fieldValue, conditionMatch.expression)
		return false
	}
}

func (conditionMatch *ConditionMatch) IsEmpty() bool {
	return conditionMatch.Program == nil
}

func BuildConditionMatcher(objectTypeFullName string, fieldKind reflect.Kind, objectFieldName string, tagName string, subCondition string, errMsg string) {
	if constant.MATCH != tagName {
		return
	}

	if fieldKind == reflect.Slice {
		return
	}
	if !strings.Contains(subCondition, constant.Condition) || !strings.Contains(subCondition, constant.EQUAL) {
		return
	}

	index := strings.Index(subCondition, "=")
	expression := subCondition[index+1:]

	if expression == "" {
		return
	}

	tree, err := parser.Parse(rmvWell(expression))
	if err != nil {
		log.Errorf("脚本：%v 解析异常：%v", expression, err.Error())
		return
	}

	program, err := compiler.Compile(tree, nil)
	if err != nil {
		log.Errorf("脚本: %v 编译异常：%v", expression, err.Error())
		return
	}
	addMatcher(objectTypeFullName, objectFieldName, &ConditionMatch{Program: program, expression: expression}, errMsg, true)
}
