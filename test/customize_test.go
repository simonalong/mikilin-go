package main

import (
	mikilin "github.com/simonalong/mikilin-go"
	"github.com/simonalong/mikilin-go/test/assert"
	"github.com/simonalong/mikilin-go/test/fun"
	"testing"
)

func TestCustomize1(t *testing.T) {
	var value fun.CustomizeEntity1
	var result bool
	var err string

	// 测试 正常情况
	value = fun.CustomizeEntity1{Name: "zhou"}
	result, _ = mikilin.Check(value, "name")
	assert.True(t, result)

	// 测试 正常情况
	value = fun.CustomizeEntity1{Name: "宋江"}
	result, _ = mikilin.Check(value, "name")
	assert.True(t, result)

	// 测试 异常情况
	value = fun.CustomizeEntity1{Name: "陈真"}
	result, err = mikilin.Check(value)
	assert.Equal(t, err, "核查错误：属性 Name 的值 陈真 没命中只允许条件回调 [fun.Judge1] ", result, false)
}

func TestCustomize2(t *testing.T) {
	var value fun.CustomizeEntity2
	var result bool
	var err string

	// 测试 正常情况
	value = fun.CustomizeEntity2{Name: "zhou"}
	result, err = mikilin.Check(value, "name")
	assert.True(t, result)

	// 测试 正常情况
	value = fun.CustomizeEntity2{Name: "宋江"}
	result, err = mikilin.Check(value, "name")
	assert.True(t, result)

	// 测试 异常情况
	value = fun.CustomizeEntity2{Name: "陈真"}
	result, err = mikilin.Check(value)
	assert.Equal(t, err, "核查错误：没有命中可用的值'zhou'和'宋江'", result, false)
}

func TestCustomize3(t *testing.T) {
	var value fun.CustomizeEntity3
	var result bool
	var err string

	// 测试 正常情况
	value = fun.CustomizeEntity3{Name: "zhou", Age: 20}
	result, err = mikilin.Check(value, "name")
	assert.True(t, result)

	// 测试 正常情况
	value = fun.CustomizeEntity3{Name: "宋江", Age: 20}
	result, _ = mikilin.Check(value, "name")
	assert.True(t, result)

	// 测试 异常情况
	value = fun.CustomizeEntity3{Name: "陈真"}
	result, err = mikilin.Check(value)
	assert.Equal(t, err, "核查错误：没有命中可用的值'zhou'和'宋江'", result, false)

	// 测试 正常情况
	value = fun.CustomizeEntity3{Name: "zhou", Age: 13}
	result, _ = mikilin.Check(value)
	assert.True(t, result)

	// 测试 异常情况
	value = fun.CustomizeEntity3{Name: "zhou", Age: 10}
	result, err = mikilin.Check(value)
	assert.Equal(t, err, "核查错误：用户[zhou]没有满足年龄age > 12，当前年龄为：10", result, false)
}

func init() {
	mikilin.RegisterFun("fun.Judge1", fun.JudgeString1)
	mikilin.RegisterFun("fun.Judge2", fun.JudgeString2)
	mikilin.RegisterFun("fun.Judge3", fun.JudgeString3)
}
