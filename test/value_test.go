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
	Name string `match:"value={inner_zhou, inner_宋江}"`
	Age  int    `match:"value={2212, 2213}"`
}

type ValueStructEntity struct {
	Name string `match:"value={zhou, 宋江}"`
	Age  int    `match:"value={12, 13}"`

	InnerA ValueInnerEntity  `match:"check"`
	InnerB *ValueInnerEntity `match:"check"`
}

type ValueMapValueEntity struct {
	Name string `match:"value={zhou, 宋江}"`
	Age  int    `match:"value={12, 13}"`

	InnerMap map[string]ValueInnerEntity `match:"check"`
}

type ValueMapKeyEntity struct {
	Name string `match:"value={zhou, 宋江}"`
	Age  int    `match:"value={12, 13}"`

	InnerMap map[ValueInnerEntity]string `match:"check"`
}

// 测试基本类型
func TestValueBase(t *testing.T) {
	var value ValueBaseEntity
	var result bool
	var err string

	//测试 正常情况
	value = ValueBaseEntity{Age: 12}
	result, err = mikilin.Check(value, "age")
	assert.TrueErr(t, result, err)

	// 测试 正常情况
	value = ValueBaseEntity{Age: 13}
	result, err = mikilin.Check(value, "age")
	assert.TrueErr(t, result, err)

	// 测试 异常情况
	value = ValueBaseEntity{Age: 14}
	result, err = mikilin.Check(value, "age")
	assert.Equal(t, false, result, "核查错误：属性 \"Age\" 的值 14 不在只可用列表 [12,13] 中", err)

	// 测试 正常情况
	value = ValueBaseEntity{Name: "zhou"}
	result, err = mikilin.Check(value, "name")
	assert.TrueErr(t, result, err)

	// 测试 正常情况
	value = ValueBaseEntity{Name: "宋江"}
	result, err = mikilin.Check(value, "name")
	assert.TrueErr(t, result, err)

	// 测试 异常情况
	value = ValueBaseEntity{Name: "陈真"}
	result, err = mikilin.Check(value, "name")
	assert.Equal(t, false, result, "核查错误：属性 \"Name\" 的值 \"陈真\" 不在只可用列表 [\"zhou\",\"宋江\"] 中", err)
}

// 测试Struct类型
func TestValueStruct(t *testing.T) {
	var value ValueStructEntity
	var result bool
	var err string
	// 测试 正常情况
	value = ValueStructEntity{InnerA: ValueInnerEntity{
		Age:  2212,
		Name: "inner_宋江",
	}}
	result, err = mikilin.Check(value, "innerA")
	assert.TrueErr(t, result, err)

	// 测试 正常情况
	value = ValueStructEntity{InnerA: ValueInnerEntity{
		Age:  2213,
		Name: "inner_宋江",
	}}
	result, err = mikilin.Check(value, "innerA")
	assert.TrueErr(t, result, err)

	// 测试 异常情况
	value = ValueStructEntity{InnerA: ValueInnerEntity{
		Age: 2211,
	}, InnerB: &ValueInnerEntity{
		Age: 2214,
	}}
	result, err = mikilin.Check(value, "innerA", "innerB")
	assert.Equal(t, false, result, "核查错误：属性 \"Name\" 的值 \"\" 不在只可用列表 [\"inner_zhou\",\"inner_宋江\"] 中", err)
}

// 测试Map：value的验证
func TestValueMapValue(t *testing.T) {
	var value ValueMapValueEntity
	var result bool
	var err string
	var innerMap map[string]ValueInnerEntity

	// 测试 正常情况
	value = ValueMapValueEntity{}
	innerMap = make(map[string]ValueInnerEntity)
	innerMap["a"] = ValueInnerEntity{Age: 2212, Name: "inner_zhou"}
	value.InnerMap = innerMap
	result, err = mikilin.Check(value, "InnerMap")
	assert.TrueErr(t, result, err)

	// 测试 正常情况
	value = ValueMapValueEntity{}
	innerMap = make(map[string]ValueInnerEntity)
	innerMap["a"] = ValueInnerEntity{Age: 2213, Name: "inner_宋江"}
	value.InnerMap = innerMap
	result, err = mikilin.Check(value, "InnerMap")
	assert.TrueErr(t, result, err)

	// 测试 异常情况
	value = ValueMapValueEntity{}
	innerMap = make(map[string]ValueInnerEntity)
	innerMap["a"] = ValueInnerEntity{Age: 2213}
	value.InnerMap = innerMap
	result, err = mikilin.Check(value, "InnerMap")
	assert.Equal(t, false, result, "核查错误：属性 \"Name\" 的值 \"\" 不在只可用列表 [\"inner_zhou\",\"inner_宋江\"] 中", err)

	// 测试 异常情况
	value = ValueMapValueEntity{}
	innerMap = make(map[string]ValueInnerEntity)
	innerMap["a"] = ValueInnerEntity{Age: 2213, Name: "inner_陈"}
	value.InnerMap = innerMap
	result, err = mikilin.Check(value, "InnerMap")
	assert.Equal(t, false, result, "核查错误：属性 \"Name\" 的值 \"inner_陈\" 不在只可用列表 [\"inner_zhou\",\"inner_宋江\"] 中", err)
}

// 测试Map：key的验证
func TestValueMapKey(t *testing.T) {
	var value ValueMapKeyEntity
	var result bool
	var err string
	var innerMap map[ValueInnerEntity]string

	// 测试 正常情况
	value = ValueMapKeyEntity{}
	innerMap = make(map[ValueInnerEntity]string)
	innerMap[ValueInnerEntity{Age: 2212, Name: "inner_zhou"}] = "a"
	value.InnerMap = innerMap
	result, err = mikilin.Check(value, "InnerMap")
	assert.TrueErr(t, result, err)

	// 测试 正常情况
	value = ValueMapKeyEntity{}
	innerMap = make(map[ValueInnerEntity]string)
	innerMap[ValueInnerEntity{Age: 2213, Name: "inner_zhou"}] = "a"
	value.InnerMap = innerMap
	result, err = mikilin.Check(value, "InnerMap")
	assert.TrueErr(t, result, err)

	// 测试 异常情况
	value = ValueMapKeyEntity{}
	innerMap = make(map[ValueInnerEntity]string)
	innerMap[ValueInnerEntity{Age: 2214, Name: "inner_zhou"}] = "a"
	value.InnerMap = innerMap
	result, err = mikilin.Check(value, "InnerMap")
	assert.Equal(t, false, result, "核查错误：属性 \"Age\" 的值 2214 不在只可用列表 [2212,2213] 中", err)

}

// 测试Map：value的指针验证
func TestValueMapValuePtr(t *testing.T) {
	var value ValueMapValueEntity
	var result bool
	var err string
	var innerMap map[string]ValueInnerEntity

	// 测试 正常情况
	value = ValueMapValueEntity{}
	innerMap = make(map[string]ValueInnerEntity)
	innerMap["a"] = ValueInnerEntity{Age: 2212, Name: "inner_zhou"}
	value.InnerMap = innerMap
	result, err = mikilin.Check(value, "InnerMap")
	assert.TrueErr(t, result, err)

	// 测试 正常情况
	value = ValueMapValueEntity{}
	innerMap = make(map[string]ValueInnerEntity)
	innerMap["a"] = ValueInnerEntity{Age: 2213, Name: "inner_宋江"}
	value.InnerMap = innerMap
	result, err = mikilin.Check(value, "InnerMap")
	assert.TrueErr(t, result, err)

	// 测试 异常情况
	value = ValueMapValueEntity{}
	innerMap = make(map[string]ValueInnerEntity)
	innerMap["a"] = ValueInnerEntity{Age: 2213}
	value.InnerMap = innerMap
	result, err = mikilin.Check(value, "InnerMap")
	assert.Equal(t, false, result, "核查错误：属性 \"Name\" 的值 \"\" 不在只可用列表 [\"inner_zhou\",\"inner_宋江\"] 中", err)

	// 测试 异常情况
	value = ValueMapValueEntity{}
	innerMap = make(map[string]ValueInnerEntity)
	innerMap["a"] = ValueInnerEntity{Age: 2213, Name: "inner_陈"}
	value.InnerMap = innerMap
	result, err = mikilin.Check(value, "InnerMap")
	assert.Equal(t, false, result, "核查错误：属性 \"Name\" 的值 \"inner_陈\" 不在只可用列表 [\"inner_zhou\",\"inner_宋江\"] 中", err)
}
