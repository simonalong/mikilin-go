package main

import (
	mikilin "github.com/SimonAlong/Mikilin-go"
	"github.com/SimonAlong/Mikilin-go/test/assert"
	"testing"
	"time"
)

// 整数类型1
type RangeIntEntity1 struct {
	Name string
	Age  int `match:"range=[1, 2]"`
}

// 整数类型2
type RangeIntEntity2 struct {
	Name string
	Age  int `match:"range=[3，]"`
}

// 整数类型3
type RangeIntEntity3 struct {
	Name string
	Age  int `match:"range=[3,)"`
}

// 整数类型4
type RangeIntEntity4 struct {
	Name string
	Age  int `match:"range=[2,1]"`
}

// 整数类型5
type RangeIntEntity5 struct {
	Name string
	Age  int `match:"range=(2, 7]"`
}

// 整数类型6
type RangeIntEntity6 struct {
	Name string
	Age  int `match:"range=(2, 7)"`
}

// 整数类型7
type RangeIntEntity7 struct {
	Name string
	Age  int `match:"range=(,7)"`
}

// 中文的逗号测试
type RangeIntEntityChina struct {
	Name string
	Age  int `match:"range=[1，10]"`
}

// 浮点数类型
type RangeFloatEntity struct {
	Name  string
	money float32 `match:"range=[12.3， 12.31]"`
}

// 字符类型
type RangeStringEntity struct {
	Name string `match:"range=[2, 12]"`
	Age  int
}

// 分片类型
type RangeSliceEntity struct {
	Name string
	Age  []int `match:"range=[2, 6]"`
}

// 时间类型
type RangeTimeEntity struct {
	createTime time.Time `match:"range=[]"`
}

// 时间计算
type RangeTimeCalEntity struct {
	Name       string
	createTime time.Time `match:"range=[]"`
}

// 测试基本类型1
func TestRangeBase1(t *testing.T) {
	var value RangeIntEntity1
	var result bool
	var err string

	//测试 正常情况
	value = RangeIntEntity1{Age: 1}
	result, err = mikilin.Check(value, "age")
	assert.TrueErr(t, result, err)

	//测试 正常情况
	value = RangeIntEntity1{Age: 3}
	result, err = mikilin.Check(value, "age")
	assert.Equal(t, "核查错误：属性 Age 的 3 没有命中只允许的范围 [1,2]", err, result, false)
}

// 测试基本类型2
func TestRangeBase2(t *testing.T) {
	var value RangeIntEntity2
	var result bool
	var err string

	//测试 正常情况
	value = RangeIntEntity2{Age: 3}
	result, err = mikilin.Check(value, "age")
	assert.TrueErr(t, result, err)

	//测试 正常情况
	value = RangeIntEntity2{Age: 5}
	result, err = mikilin.Check(value, "age")
	assert.TrueErr(t, result, err)

	//测试 异常情况
	value = RangeIntEntity2{Age: 2}
	result, err = mikilin.Check(value, "age")
	assert.Equal(t, "核查错误：属性 Age 的 2 没有命中只允许的范围 [3，]", err, result, false)
}

// 测试基本类型3
func TestRangeBase3(t *testing.T) {
	var value RangeIntEntity3
	var result bool
	var err string

	//测试 正常情况
	value = RangeIntEntity3{Age: 3}
	result, err = mikilin.Check(value, "age")
	assert.TrueErr(t, result, err)

	//测试 正常情况
	value = RangeIntEntity3{Age: 5}
	result, err = mikilin.Check(value, "age")
	assert.TrueErr(t, result, err)

	//测试 异常情况
	value = RangeIntEntity3{Age: 2}
	result, err = mikilin.Check(value, "age")
	assert.Equal(t, "核查错误：属性 Age 的 2 没有命中只允许的范围 [3,)", err, result, false)
}

// 测试基本类型4
func TestRangeBase4(t *testing.T) {

	// todo
	//测试 正常情况
	//value = RangeIntEntity4{Age: 3}
	//result, err = mikilin.Check(value, "age")
	//assert.TrueErr(t, result, err)
}

// 测试基本类型5
func TestRangeBase5(t *testing.T) {
	var value RangeIntEntity5
	var result bool
	var err string

	//测试 正常情况
	value = RangeIntEntity5{Age: 3}
	result, err = mikilin.Check(value, "age")
	assert.TrueErr(t, result, err)

	//测试 正常情况
	value = RangeIntEntity5{Age: 7}
	result, err = mikilin.Check(value, "age")
	assert.TrueErr(t, result, err)

	//测试 异常情况
	value = RangeIntEntity5{Age: 8}
	result, err = mikilin.Check(value, "age")
	assert.Equal(t, "核查错误：属性 Age 的 8 没有命中只允许的范围 (2,7]", err, result, false)

	//测试 异常情况
	value = RangeIntEntity5{Age: 2}
	result, err = mikilin.Check(value, "age")
	assert.Equal(t, "核查错误：属性 Age 的 2 没有命中只允许的范围 (2,7]", err, result, false)
}

// 测试基本类型6
func TestRangeBase6(t *testing.T) {
	var value RangeIntEntity6
	var result bool
	var err string

	//测试 正常情况
	value = RangeIntEntity6{Age: 3}
	result, err = mikilin.Check(value, "age")
	assert.TrueErr(t, result, err)

	//测试 正常情况
	value = RangeIntEntity6{Age: 7}
	result, err = mikilin.Check(value, "age")
	assert.Equal(t, "核查错误：属性 Age 的 7 没有命中只允许的范围 (2,7)", err, result, false)

	//测试 异常情况
	value = RangeIntEntity6{Age: 8}
	result, err = mikilin.Check(value, "age")
	assert.Equal(t, "核查错误：属性 Age 的 8 没有命中只允许的范围 (2,7)", err, result, false)

	//测试 异常情况
	value = RangeIntEntity6{Age: 2}
	result, err = mikilin.Check(value, "age")
	assert.Equal(t, "核查错误：属性 Age 的 2 没有命中只允许的范围 (2,7)", err, result, false)
}

// 测试基本类型7
func TestRangeBase7(t *testing.T) {
	var value RangeIntEntity7
	var result bool
	var err string

	//测试 正常情况
	value = RangeIntEntity7{Age: 3}
	result, err = mikilin.Check(value, "age")
	assert.TrueErr(t, result, err)

	//测试 正常情况
	value = RangeIntEntity7{Age: -1}
	result, err = mikilin.Check(value, "age")
	assert.TrueErr(t, result, err)

	//测试 异常情况
	value = RangeIntEntity7{Age: 7}
	result, err = mikilin.Check(value, "age")
	assert.Equal(t, "核查错误：属性 Age 的 7 没有命中只允许的范围 (,7)", err, result, false)

	//测试 异常情况
	value = RangeIntEntity7{Age: 8}
	result, err = mikilin.Check(value, "age")
	assert.Equal(t, "核查错误：属性 Age 的 8 没有命中只允许的范围 (,7)", err, result, false)
}

// 测试中文逗号表示
func TestRangeBaseChinaComma(t *testing.T) {
	var value RangeIntEntityChina
	var result bool
	var err string

	//测试 正常情况
	value = RangeIntEntityChina{Age: 3}
	result, err = mikilin.Check(value, "age")
	assert.TrueErr(t, result, err)

	//测试 正常情况
	value = RangeIntEntityChina{Age: 5}
	result, err = mikilin.Check(value, "age")
	assert.TrueErr(t, result, err)

	//测试 异常情况
	value = RangeIntEntityChina{Age: 0}
	result, err = mikilin.Check(value, "age")
	assert.Equal(t, "核查错误：属性 Age 的 0 没有命中只允许的范围 [1，10]", err, result, false)

	//测试 异常情况
	value = RangeIntEntityChina{Age: 12}
	result, err = mikilin.Check(value, "age")
	assert.Equal(t, "核查错误：属性 Age 的 12 没有命中只允许的范围 [1，10]", err, result, false)
}
