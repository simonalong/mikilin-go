package main

import (
	mikilin "github.com/simonalong/mikilin-go"
	"github.com/simonalong/mikilin-go/test/assert"
	"testing"
)

type IsUnBlankEntity1 struct {
	Name string `match:"isUnBlank"`
	Age  int
}

type IsUnBlankEntity2 struct {
	Name string `match:"isUnBlank=true"`
	Age  int
}

type IsUnBlankEntity3 struct {
	Name string `match:"isUnBlank=false"`
	Age  int
}

// 测试基本类型：简化版
func TestIsUnBlank1(t *testing.T) {
	var value IsUnBlankEntity1
	var result bool
	var err string

	//测试 正常情况
	value = IsUnBlankEntity1{Name: "zhou"}
	result, err = mikilin.Check(value, "name")
	assert.TrueErr(t, result, err)

	// 测试 正常情况
	value = IsUnBlankEntity1{Name: ""}
	result, err = mikilin.Check(value, "name")
	assert.Equal(t, err, "核查错误：属性 Name 的值为非空字符", result, false)
}

func TestIsUnBlank2(t *testing.T) {
	var value IsUnBlankEntity1
	var result bool
	var err string

	//测试 正常情况
	value = IsUnBlankEntity1{Name: "zhou"}
	result, err = mikilin.Check(value, "name")
	assert.TrueErr(t, result, err)

	// 测试 正常情况
	value = IsUnBlankEntity1{Name: ""}
	result, err = mikilin.Check(value, "name")
	assert.Equal(t, err, "核查错误：属性 Name 的值为非空字符", result, false)
}

// 测试基本类型
func TestIsUnBlank3(t *testing.T) {
	var value IsUnBlankEntity3
	var result bool
	var err string

	//测试 正常情况
	value = IsUnBlankEntity3{Name: ""}
	result, err = mikilin.Check(value, "name")
	assert.TrueErr(t, result, err)

	// 测试 正常情况
	value = IsUnBlankEntity3{Name: "zhou"}
	result, err = mikilin.Check(value, "name")
	assert.Equal(t, err, "核查错误：属性 Name 的值为空字符", result, false)
}
