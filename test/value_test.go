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

type ValueBaseEntityOne struct {
	Name string `match:"value={zhou, 宋江}"`
	Age  int    `match:"value=12"`
}

type ValueBasePtrEntity struct {
	Name *string `match:"value={zhou, 宋江}"`
	Age  *int    `match:"value={12, 13}"`
}

type ValueInnerEntity struct {
	Name string `match:"value={inner_zhou, inner_宋江}"`
	Age  int    `match:"value={2212, 2213}"`
}

type ValueStructEntity struct {
	Name string `match:"value={zhou, 宋江}"`
	Age  int    `match:"value={12, 13}"`

	Inner ValueInnerEntity `match:"check"`
}

type ValueStructPtrEntity struct {
	Name string `match:"value={zhou, 宋江}"`
	Age  int    `match:"value={12, 13}"`

	Inner *ValueInnerEntity `match:"check"`
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

type ValueMapValuePtrEntity struct {
	Name string `match:"value={zhou, 宋江}"`
	Age  int    `match:"value={12, 13}"`

	InnerMap map[string]*ValueInnerEntity `match:"check"`
}

type ValueArrayEntity struct {
	Inner [3]ValueInnerEntity `match:"check"`
}

type ValueArrayPtrEntity struct {
	Inner [3]*ValueInnerEntity `match:"check"`
}

type ValueSliceEntity struct {
	Inner []ValueInnerEntity `match:"check"`
}

type ValueSlicePtrEntity struct {
	Inner []*ValueInnerEntity `match:"check"`
}

// 测试基本类型：一个值的情况
func TestValueBase2(t *testing.T) {
	var value ValueBaseEntityOne
	var result bool
	var err string

	//测试 正常情况
	value = ValueBaseEntityOne{Age: 12}
	result, err = mikilin.Check(value, "age")
	assert.TrueErr(t, result, err)

	// 测试 正常情况
	value = ValueBaseEntityOne{Age: 13}
	result, err = mikilin.Check(value, "age")
	assert.Equal(t, "核查错误：属性 Age 的值 13 不在只可用列表 [12] 中", err, result, false)
}

// 测试基本类型：多个值的情况
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
	assert.Equal(t, "核查错误：属性 Age 的值 14 不在只可用列表 [12 13] 中", err, false, result)

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
	assert.Equal(t, "核查错误：属性 Name 的值 陈真 不在只可用列表 [zhou 宋江] 中", err, false, result)
}

// 测试基本类型：指针类型
func TestValueBasePtr(t *testing.T) {
	var value *ValueBasePtrEntity
	var result bool
	var err string
	var age int
	var name string

	//测试 正常情况
	value = &ValueBasePtrEntity{}
	age = 12
	value.Age = &age
	result, err = mikilin.Check(value, "age")
	assert.TrueErr(t, result, err)

	//测试 正常情况
	value = &ValueBasePtrEntity{}
	age = 13
	value.Age = &age
	result, err = mikilin.Check(value, "age")
	assert.TrueErr(t, result, err)

	// 测试 异常情况
	value = &ValueBasePtrEntity{}
	age = 14
	value.Age = &age
	result, err = mikilin.Check(value, "age")
	assert.Equal(t, "核查错误：属性 Age 的值 14 不在只可用列表 [12 13] 中", err, false, result)

	// 测试 正常情况
	value = &ValueBasePtrEntity{}
	name = "zhou"
	value.Name = &name
	result, err = mikilin.Check(value, "name")
	assert.TrueErr(t, result, err)

	// 测试 正常情况
	value = &ValueBasePtrEntity{}
	name = "宋江"
	value.Name = &name
	result, err = mikilin.Check(value, "name")
	assert.TrueErr(t, result, err)

	// 测试 异常情况
	value = &ValueBasePtrEntity{}
	name = "陈真"
	value.Name = &name
	result, err = mikilin.Check(value, "name")
	assert.Equal(t, "核查错误：属性 Name 的值 陈真 不在只可用列表 [zhou 宋江] 中", err, false, result)
}

// 测试Struct类型
func TestValueStruct(t *testing.T) {
	var value ValueStructEntity
	var result bool
	var err string
	//测试 正常情况
	value = ValueStructEntity{Inner: ValueInnerEntity{
		Age:  2212,
		Name: "inner_宋江",
	}}
	result, err = mikilin.Check(value, "inner")
	assert.TrueErr(t, result, err)

	//测试 正常情况
	value = ValueStructEntity{Inner: ValueInnerEntity{
		Age:  2213,
		Name: "inner_宋江",
	}}
	result, err = mikilin.Check(value, "inner")
	assert.TrueErr(t, result, err)

	//测试 异常情况
	value = ValueStructEntity{Inner: ValueInnerEntity{
		Age:  2214,
		Name: "inner_宋江",
	}}
	result, err = mikilin.Check(value, "inner")
	assert.Equal(t, "核查错误：属性 Age 的值 2214 不在只可用列表 [2212 2213] 中", err, false, result)
}

// 测试Struct类型：指针类型
func TestValueStructPtr(t *testing.T) {
	var value ValueStructPtrEntity
	var result bool
	var err string
	// 测试 正常情况
	value = ValueStructPtrEntity{Inner: &ValueInnerEntity{
		Age:  2212,
		Name: "inner_宋江",
	}}
	result, err = mikilin.Check(value, "inner")
	assert.TrueErr(t, result, err)

	//测试 正常情况
	value = ValueStructPtrEntity{Inner: &ValueInnerEntity{
		Age:  2213,
		Name: "inner_宋江",
	}}
	result, err = mikilin.Check(value, "inner")
	assert.TrueErr(t, result, err)

	// 测试 核查其他情况
	value = ValueStructPtrEntity{Age: 12}
	result, err = mikilin.Check(value, "age")
	assert.TrueErr(t, result, err)

	// 测试 核查其他情况
	value = ValueStructPtrEntity{Age: 12, Inner: &ValueInnerEntity{
		Age:  2213,
		Name: "inner_宋江",
	}}
	result, err = mikilin.Check(value, "age", "inner")
	assert.TrueErr(t, result, err)

	// 测试 核查其他情况
	value = ValueStructPtrEntity{Age: 14, Inner: &ValueInnerEntity{
		Age:  2213,
		Name: "inner_宋江",
	}}
	result, err = mikilin.Check(value, "age", "inner")
	assert.Equal(t, false, result, "核查错误：属性 Age 的值 14 不在只可用列表 [12 13] 中", err)

	//测试 异常情况
	value = ValueStructPtrEntity{Inner: &ValueInnerEntity{
		Age:  2214,
		Name: "inner_宋江",
	}}
	result, err = mikilin.Check(value, "inner")
	assert.Equal(t, false, result, "核查错误：属性 Age 的值 2214 不在只可用列表 [2212 2213] 中", err)
}

// 测试Map：value的验证
func TestValueMapValue(t *testing.T) {
	var value ValueMapValueEntity
	var result bool
	var err string
	var innerMap map[string]ValueInnerEntity

	// 测试 正常情况
	value = ValueMapValueEntity{Age: 12, Name: "宋江"}
	result, err = mikilin.Check(value)
	assert.TrueErr(t, result, err)

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
	assert.Equal(t, false, result, "核查错误：属性 Name 的值  不在只可用列表 [inner_zhou inner_宋江] 中", err)

	// 测试 异常情况
	value = ValueMapValueEntity{}
	innerMap = make(map[string]ValueInnerEntity)
	innerMap["a"] = ValueInnerEntity{Age: 2213, Name: "inner_陈"}
	value.InnerMap = innerMap
	result, err = mikilin.Check(value, "InnerMap")
	assert.Equal(t, false, result, "核查错误：属性 Name 的值 inner_陈 不在只可用列表 [inner_zhou inner_宋江] 中", err)
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
	assert.Equal(t, false, result, "核查错误：属性 Age 的值 2214 不在只可用列表 [2212 2213] 中", err)
}

// 测试Map：value的指针验证
func TestValueMapValuePtr(t *testing.T) {
	var value ValueMapValuePtrEntity
	var result bool
	var err string
	var innerMap map[string]*ValueInnerEntity

	// 测试 正常情况
	value = ValueMapValuePtrEntity{}
	innerMap = make(map[string]*ValueInnerEntity)
	innerMap["a"] = &ValueInnerEntity{Age: 2212, Name: "inner_zhou"}
	value.InnerMap = innerMap
	result, err = mikilin.Check(value, "InnerMap")
	assert.TrueErr(t, result, err)

	// 测试 正常情况
	value = ValueMapValuePtrEntity{}
	innerMap = make(map[string]*ValueInnerEntity)
	innerMap["a"] = &ValueInnerEntity{Age: 2213, Name: "inner_宋江"}
	value.InnerMap = innerMap
	result, err = mikilin.Check(value, "InnerMap")
	assert.TrueErr(t, result, err)

	// 测试 异常情况
	value = ValueMapValuePtrEntity{}
	innerMap = make(map[string]*ValueInnerEntity)
	innerMap["a"] = &ValueInnerEntity{Age: 2213}
	value.InnerMap = innerMap
	result, err = mikilin.Check(value, "InnerMap")
	assert.Equal(t, false, result, "核查错误：属性 Name 的值  不在只可用列表 [inner_zhou inner_宋江] 中", err)

	// 测试 异常情况
	value = ValueMapValuePtrEntity{}
	innerMap = make(map[string]*ValueInnerEntity)
	innerMap["a"] = &ValueInnerEntity{Age: 2213, Name: "inner_陈"}
	value.InnerMap = innerMap
	result, err = mikilin.Check(value, "InnerMap")
	assert.Equal(t, false, result, "核查错误：属性 Name 的值 inner_陈 不在只可用列表 [inner_zhou inner_宋江] 中", err)
}

// 测试Array
func TestValueArray(t *testing.T) {
	var value ValueArrayEntity
	var result bool
	var err string
	innerArray := [3]ValueInnerEntity{}

	// 正常
	value = ValueArrayEntity{}
	innerArray[0] = ValueInnerEntity{Age: 2212, Name: "inner_zhou"}
	innerArray[1] = ValueInnerEntity{Age: 2213, Name: "inner_zhou"}
	innerArray[2] = ValueInnerEntity{Age: 2212, Name: "inner_宋江"}
	value.Inner = innerArray

	result, err = mikilin.Check(value, "inner")
	assert.TrueErr(t, result, err)

	// 异常
	value = ValueArrayEntity{}
	innerArray[0] = ValueInnerEntity{Age: 2212, Name: "inner_zhou"}
	innerArray[1] = ValueInnerEntity{Age: 2213, Name: "inner_zhou"}
	innerArray[2] = ValueInnerEntity{Age: 2214, Name: "inner_宋江"}
	value.Inner = innerArray
	result, err = mikilin.Check(value, "inner")
	assert.Equal(t, "核查错误：属性 Age 的值 2214 不在只可用列表 [2212 2213] 中", err, false, result)
}

// 测试Array：指针类型
func TestValueArrayPtr(t *testing.T) {
	var value ValueArrayPtrEntity
	var result bool
	var err string
	innerArray := [3]*ValueInnerEntity{}

	// 正常
	value = ValueArrayPtrEntity{}
	innerArray[0] = &ValueInnerEntity{Age: 2212, Name: "inner_zhou"}
	innerArray[1] = &ValueInnerEntity{Age: 2213, Name: "inner_zhou"}
	innerArray[2] = &ValueInnerEntity{Age: 2212, Name: "inner_宋江"}
	value.Inner = innerArray

	result, err = mikilin.Check(value, "inner")
	assert.TrueErr(t, result, err)

	// 异常
	value = ValueArrayPtrEntity{}
	innerArray[0] = &ValueInnerEntity{Age: 2212, Name: "inner_zhou"}
	innerArray[1] = &ValueInnerEntity{Age: 2213, Name: "inner_zhou"}
	innerArray[2] = &ValueInnerEntity{Age: 2214, Name: "inner_宋江"}
	value.Inner = innerArray
	result, err = mikilin.Check(value, "inner")
	assert.Equal(t, "核查错误：属性 Age 的值 2214 不在只可用列表 [2212 2213] 中", err, false, result)
}

// 测试 Slice
func TestValueSlice(t *testing.T) {
	var value ValueSliceEntity
	var result bool
	var err string
	innerSlice := []ValueInnerEntity{}

	// 正常
	//value = ValueSliceEntity{}
	//innerSlice = append(innerSlice, ValueInnerEntity{Age: 2212, Name: "inner_zhou"})
	//innerSlice = append(innerSlice, ValueInnerEntity{Age: 2213, Name: "inner_宋江"})
	//innerSlice = append(innerSlice, ValueInnerEntity{Age: 2212, Name: "inner_宋江"})
	//value.Inner = innerSlice
	//
	//result, err = mikilin.Check(value, "inner")
	//assert.TrueErr(t, result, err)

	// 异常
	value = ValueSliceEntity{}
	innerSlice = append(innerSlice, ValueInnerEntity{Age: 2212, Name: "inner_zhou"})
	innerSlice = append(innerSlice, ValueInnerEntity{Age: 2213, Name: "inner_zhou"})
	innerSlice = append(innerSlice, ValueInnerEntity{Age: 2214, Name: "inner_宋江"})
	value.Inner = innerSlice

	result, err = mikilin.Check(value, "inner")
	assert.Equal(t, "核查错误：属性 Age 的值 2214 不在只可用列表 [2212 2213] 中", err, false, result)
}

// 测试 Slice：指针类型
func TestValueSlicePtr(t *testing.T) {
	var value ValueSlicePtrEntity
	var result bool
	var err string
	innerSlice := []*ValueInnerEntity{}

	// 正常
	value = ValueSlicePtrEntity{}
	innerSlice = append(innerSlice, &ValueInnerEntity{Age: 2212, Name: "inner_zhou"})
	innerSlice = append(innerSlice, &ValueInnerEntity{Age: 2213, Name: "inner_zhou"})
	innerSlice = append(innerSlice, &ValueInnerEntity{Age: 2212, Name: "inner_宋江"})
	value.Inner = innerSlice

	result, err = mikilin.Check(value, "inner")
	assert.TrueErr(t, result, err)

	// 异常
	value = ValueSlicePtrEntity{}
	innerSlice = append(innerSlice, &ValueInnerEntity{Age: 2212, Name: "inner_zhou"})
	innerSlice = append(innerSlice, &ValueInnerEntity{Age: 2213, Name: "inner_zhou"})
	innerSlice = append(innerSlice, &ValueInnerEntity{Age: 2214, Name: "inner_宋江"})
	value.Inner = innerSlice
	result, err = mikilin.Check(value, "inner")
	assert.Equal(t, "核查错误：属性 Age 的值 2214 不在只可用列表 [2212 2213] 中", err, false, result)
}

// value的基准测试
func Benchmark_Value(b *testing.B) {
	var value ValueBaseEntityOne
	for i := 0; i < b.N; i++ {
		value = ValueBaseEntityOne{Age: 12}
		mikilin.Check(value, "age")
	}
}
