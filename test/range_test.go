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
	Age  int `match:"range=[1，]"`
}

// 整数类型3
type RangeIntEntity3 struct {
	Name string
	Age  int `match:"range=[1,)"`
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

// 浮点数类型
type RangeFloatEntity struct {
	Name  string
	money float32 `match:"range=[]"`
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

// 测试基本类型
func TestRangeBase(t *testing.T) {
	var value RangeIntEntity1
	var result bool
	var err string

	//测试 正常情况
	//value = RangeIntEntity1{Age: 1}
	//result, err = mikilin.Check(value, "age")
	//assert.TrueErr(t, result, err)

	//测试 正常情况
	value = RangeIntEntity1{Age: 3}
	result, err = mikilin.Check(value, "age")
	assert.Equal(t, "核查错误：属性 Age 的 3 没有命中只允许的范围 [1,2]", err, result, false)
}
