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
	Money float32 `match:"range=[10.37， 20.31]"`
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

// 测试整数类型1
func TestRangeInt1(t *testing.T) {
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
	assert.Equal(t, err, "核查错误：属性 Age 值 3 没有命中只允许的范围 [1,2]", result, false)
}

// 测试整数类型2
func TestRangeInt2(t *testing.T) {
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
	assert.Equal(t, err, "核查错误：属性 Age 值 2 没有命中只允许的范围 [3，]", result, false)
}

// 测试整数类型3
func TestRangeInt3(t *testing.T) {
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
	assert.Equal(t, err, "核查错误：属性 Age 值 2 没有命中只允许的范围 [3,)", result, false)
}

// 测试整数类型4
func TestRangeInt4(t *testing.T) {

	// todo 测试数据的判断异常情况
	//测试 正常情况
	//value = RangeIntEntity4{Age: 3}
	//result, err = mikilin.Check(value, "age")
	//assert.TrueErr(t, result, err)
}

// 测试整数类型5
func TestRangeInt5(t *testing.T) {
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
	assert.Equal(t, err, "核查错误：属性 Age 值 8 没有命中只允许的范围 (2,7]", result, false)

	//测试 异常情况
	value = RangeIntEntity5{Age: 2}
	result, err = mikilin.Check(value, "age")
	assert.Equal(t, err, "核查错误：属性 Age 值 2 没有命中只允许的范围 (2,7]", result, false)
}

// 测试整数类型6
func TestRangeInt6(t *testing.T) {
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
	assert.Equal(t, err, "核查错误：属性 Age 值 7 没有命中只允许的范围 (2,7)", result, false)

	//测试 异常情况
	value = RangeIntEntity6{Age: 8}
	result, err = mikilin.Check(value, "age")
	assert.Equal(t, err, "核查错误：属性 Age 值 8 没有命中只允许的范围 (2,7)", result, false)

	//测试 异常情况
	value = RangeIntEntity6{Age: 2}
	result, err = mikilin.Check(value, "age")
	assert.Equal(t, err, "核查错误：属性 Age 值 2 没有命中只允许的范围 (2,7)", result, false)
}

// 测试整数类型7
func TestRangeInt7(t *testing.T) {
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
	assert.Equal(t, err, "核查错误：属性 Age 值 7 没有命中只允许的范围 (,7)", result, false)

	//测试 异常情况
	value = RangeIntEntity7{Age: 8}
	result, err = mikilin.Check(value, "age")
	assert.Equal(t, err, "核查错误：属性 Age 值 8 没有命中只允许的范围 (,7)", result, false)
}

// 测试中文逗号表示
func TestRangeIntChinaComma(t *testing.T) {
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
	assert.Equal(t, err, "核查错误：属性 Age 值 0 没有命中只允许的范围 [1，10]", result, false)

	//测试 异常情况
	value = RangeIntEntityChina{Age: 12}
	result, err = mikilin.Check(value, "age")
	assert.Equal(t, err, "核查错误：属性 Age 值 12 没有命中只允许的范围 [1，10]", result, false)
}

// 测试浮点数类型1
func TestRangeFloat1(t *testing.T) {
	var value RangeFloatEntity
	var result bool
	var err string

	//测试 正常情况
	value = RangeFloatEntity{Money: 10.37}
	result, err = mikilin.Check(value, "money")
	assert.TrueErr(t, result, err)

	//测试 正常情况
	value = RangeFloatEntity{Money: 15.0}
	result, err = mikilin.Check(value, "money")
	assert.TrueErr(t, result, err)

	//测试 正常情况
	value = RangeFloatEntity{Money: 20.31}
	result, err = mikilin.Check(value, "money")
	assert.TrueErr(t, result, err)

	//测试 异常情况
	value = RangeFloatEntity{Money: 10.01}
	result, err = mikilin.Check(value, "money")
	assert.Equal(t, err, "核查错误：属性 Money 值 10.01 没有命中只允许的范围 [10.37，20.31]", result, false)

	//测试 异常情况
	value = RangeFloatEntity{Money: 20.32}
	result, err = mikilin.Check(value, "money")
	assert.Equal(t, err, "核查错误：属性 Money 值 20.32 没有命中只允许的范围 [10.37，20.31]", result, false)
}

// 测试字符类型1
func TestRangeString(t *testing.T) {
	var value RangeStringEntity
	var result bool
	var err string

	//测试 正常情况
	value = RangeStringEntity{Name: "zh"}
	result, err = mikilin.Check(value, "name")
	assert.TrueErr(t, result, err)

	//测试 正常情况
	value = RangeStringEntity{Name: "zhou"}
	result, err = mikilin.Check(value, "name")
	assert.TrueErr(t, result, err)

	//测试 正常情况
	value = RangeStringEntity{Name: "zhou zhen yo"}
	result, err = mikilin.Check(value, "name")
	assert.TrueErr(t, result, err)

	//测试 异常情况
	value = RangeStringEntity{Name: "zhou zhen yong"}
	result, err = mikilin.Check(value, "name")
	assert.Equal(t, err, "核查错误：属性 Name 值 zhou zhen yong 的长度没有命中只允许的范围 [2,12]", result, false)

	//测试 异常情况
	value = RangeStringEntity{Name: "z"}
	result, err = mikilin.Check(value, "name")
	assert.Equal(t, err, "核查错误：属性 Name 值 z 的长度没有命中只允许的范围 [2,12]", result, false)
}

// 测试分片类型1
func TestRangeSlice(t *testing.T) {
	var value RangeSliceEntity
	var result bool
	var err string

	//测试 正常情况
	value = RangeSliceEntity{Age: []int{1, 2}}
	result, err = mikilin.Check(value, "age")
	assert.TrueErr(t, result, err)

	//测试 正常情况
	value = RangeSliceEntity{Age: []int{1, 2, 3, 4, 5}}
	result, err = mikilin.Check(value, "age")
	assert.TrueErr(t, result, err)

	//测试 正常情况
	value = RangeSliceEntity{Age: []int{1, 2, 3, 4, 5, 6}}
	result, err = mikilin.Check(value, "age")
	assert.TrueErr(t, result, err)

	////测试 异常情况
	value = RangeSliceEntity{Age: []int{1, 2, 3, 4, 5, 6, 7}}
	result, err = mikilin.Check(value, "age")
	assert.Equal(t, err, "核查错误：属性 Age 值 [1 2 3 4 5 6 7] 的数组长度没有命中只允许的范围 [2,6]", result, false)

	//测试 异常情况
	value = RangeSliceEntity{Age: []int{1}}
	result, err = mikilin.Check(value, "age")
	assert.Equal(t, err, "核查错误：属性 Age 值 [1] 的数组长度没有命中只允许的范围 [2,6]", result, false)
}

// 压测进行基准测试
func Benchmark_Range(b *testing.B) {
	var value RangeSliceEntity
	for i := 0; i < b.N; i++ {
		value = RangeSliceEntity{Age: []int{1}}
		mikilin.Check(value, "age")
	}
}
