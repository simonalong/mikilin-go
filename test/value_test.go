package main

import (
	mikilin "github.com/SimonAlong/Mikilin-go"
	"github.com/SimonAlong/Mikilin-go/test/assert"
	"testing"
)

type ValueBaseEntity struct {
	Name string `match:"value={zhou, 宋江}"`
	Age  int    `match:"value={12, 13}"`
}

type ValueInnerEntity struct {
	Name string `match:"value={innser_zhou, innser_宋江}"`
	Age  int    `match:"value={2212, 2213}"`
}

type ValueStructEntity struct {
	Name string `match:"value={zhou, 宋江}"`
	Age  int    `match:"value={12, 13}"`

	InnerA ValueInnerEntity  `match:"check"`
	InnerB *ValueInnerEntity `match:"check"`
}

type ValueMapEntity struct {
	Name string `match:"value={zhou, 宋江}"`
	Age  int    `match:"value={12, 13}"`

	InnerMap map[string]ValueInnerEntity `match:"check"`
}

// 测试基本类型
func TestValueBase(t *testing.T) {
	var value ValueBaseEntity
	var result bool
	var err string

	//测试 正常情况
	value = ValueBaseEntity{Age: 12}
	result, err = mikilin.Check(value, "age")
	assert.AssertTrueErr(t, result, err)

	// 测试 正常情况
	value = ValueBaseEntity{Age: 13}
	result, err = mikilin.Check(value, "age")
	assert.AssertTrueErr(t, result, err)

	// 测试 异常情况
	value = ValueBaseEntity{Age: 14}
	result, err = mikilin.Check(value, "age")
	assert.AssertFalseErr(t, result, err)

	// 测试 正常情况
	value = ValueBaseEntity{Name: "zhou"}
	result, err = mikilin.Check(value, "name")
	assert.AssertTrueErr(t, result, err)

	// 测试 正常情况
	value = ValueBaseEntity{Name: "宋江"}
	result, err = mikilin.Check(value, "name")
	assert.AssertTrueErr(t, result, err)

	// 测试 异常情况
	value = ValueBaseEntity{Name: "陈真"}
	result, err = mikilin.Check(value, "name")
	assert.AssertFalseErr(t, result, err)
}

// 测试Struct类型
func TestValueStruct(t *testing.T) {
	var value ValueStructEntity
	var result bool
	var err string
	// 测试 正常情况
	value = ValueStructEntity{InnerA: ValueInnerEntity{
		Age: 2212,
	}}
	result, err = mikilin.Check(value, "innerA")
	assert.AssertTrueErr(t, result, err)

	// 测试 正常情况
	value = ValueStructEntity{InnerA: ValueInnerEntity{
		Age: 2213,
	}}
	result, err = mikilin.Check(value, "innerA")
	assert.AssertTrueErr(t, result, err)

	// 测试 异常情况
	value = ValueStructEntity{InnerA: ValueInnerEntity{
		Age: 2211,
	}, InnerB: &ValueInnerEntity{
		Age: 2214,
	}}
	result, err = mikilin.Check(value, "innerA", "innerB")
	assert.AssertFalseErr(t, result, err)
}

// 测试Map类型
func TestValueMap(t *testing.T) {
	var value ValueMapEntity
	var result bool
	var err string
	var innerMap map[string]ValueInnerEntity

	// 测试 正常情况
	value = ValueMapEntity{}
	innerMap = make(map[string]ValueInnerEntity)
	innerMap["a"] = ValueInnerEntity{Age: 2212}
	value.InnerMap = innerMap
	result, err = mikilin.Check(value, "InnerMap")
	assert.AssertTrueErr(t, result, err)

	// 测试 正常情况
	value = ValueMapEntity{}
	innerMap = make(map[string]ValueInnerEntity)
	innerMap["a"] = ValueInnerEntity{Age: 2213}
	value.InnerMap = innerMap
	result, err = mikilin.Check(value, "InnerMap")
	assert.AssertTrueErr(t, result, err)

	// 测试 异常情况
	value = ValueMapEntity{}
	innerMap = make(map[string]ValueInnerEntity)
	innerMap["a"] = ValueInnerEntity{Age: 2214}
	value.InnerMap = innerMap
	result, err = mikilin.Check(value, "InnerMap")
	assert.AssertFalseErr(t, result, err)
}
