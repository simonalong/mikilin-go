package main

import (
	mikilin "github.com/simonalong/mikilin-go"
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
	True(t, result)

	// 测试 正常情况
	value = fun.CustomizeEntity1{Name: "宋江"}
	result, _ = mikilin.Check(value, "name")
	True(t, result)

	// 测试 异常情况
	value = fun.CustomizeEntity1{Name: "陈真"}
	result, err = mikilin.Check(value)
	Equal(t, err, "属性 Name 的值 陈真 没命中只允许条件回调 [judge1Name] ", result, false)
}

func TestCustomize2(t *testing.T) {
	var value fun.CustomizeEntity2
	var result bool
	var err string

	// 测试 正常情况
	value = fun.CustomizeEntity2{Name: "zhou"}
	result, err = mikilin.Check(value, "name")
	True(t, result)

	// 测试 正常情况
	value = fun.CustomizeEntity2{Name: "宋江"}
	result, err = mikilin.Check(value, "name")
	True(t, result)

	// 测试 异常情况
	value = fun.CustomizeEntity2{Name: "陈真"}
	result, err = mikilin.Check(value)
	Equal(t, err, "没有命中可用的值'zhou'和'宋江'", result, false)
}

func TestCustomize3(t *testing.T) {
	var value fun.CustomizeEntity3
	var result bool
	var err string

	// 测试 正常情况
	value = fun.CustomizeEntity3{Name: "zhou", Age: 20}
	result, err = mikilin.Check(value, "name")
	True(t, result)

	// 测试 正常情况
	value = fun.CustomizeEntity3{Name: "宋江", Age: 20}
	result, _ = mikilin.Check(value, "name")
	True(t, result)

	// 测试 异常情况
	value = fun.CustomizeEntity3{Name: "陈真"}
	result, err = mikilin.Check(value)
	Equal(t, err, "没有命中可用的值'zhou'和'宋江'", result, false)

	// 测试 正常情况
	value = fun.CustomizeEntity3{Name: "zhou", Age: 13}
	result, _ = mikilin.Check(value)
	True(t, result)

	// 测试 异常情况
	value = fun.CustomizeEntity3{Name: "zhou", Age: 10}
	result, err = mikilin.Check(value)
	Equal(t, err, "用户[zhou]没有满足年龄age > 12，当前年龄为：10", result, false)
}

func TestCustomize4(t *testing.T) {
	var value fun.CustomizeEntity4
	var result bool
	var err string

	// 测试 正常情况
	value = fun.CustomizeEntity4{Name: "zhou", Age: 20}
	result, err = mikilin.Check(value, "name")
	True(t, result)

	// 测试 正常情况
	value = fun.CustomizeEntity4{Name: "宋江", Age: 20}
	result, _ = mikilin.Check(value, "name")
	True(t, result)

	// 测试 异常情况
	value = fun.CustomizeEntity4{Name: "陈真"}
	result, err = mikilin.Check(value)
	Equal(t, err, "没有命中可用的值'zhou'和'宋江'", result, false)

	// 测试 正常情况
	value = fun.CustomizeEntity4{Name: "zhou", Age: 13}
	result, _ = mikilin.Check(value)
	True(t, result)

	// 测试 异常情况
	value = fun.CustomizeEntity4{Name: "zhou", Age: 10}
	result, err = mikilin.Check(value)
	Equal(t, err, "用户[zhou]没有满足年龄age > 12，当前年龄为：10", result, false)
}
