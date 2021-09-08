package main

import (
	mikilin "github.com/SimonAlong/Mikilin-go"
	"github.com/SimonAlong/Mikilin-go/test/assert"
	"testing"
)

type AcceptEntity1 struct {
	Name string `match:"value=zhou" accept:"false"`
	Age  int
}

type AcceptEntity2 struct {
	Name string `match:"isBlank=false" accept:"false"`
	Age  int
}

type AcceptEntity3 struct {
	Name string `match:"isBlank=true value=zhou" accept:"true"`
	Age  int
}

func TestAccept1(t *testing.T) {
	var value AcceptEntity1
	var result bool
	var err string

	//测试 正常情况
	value = AcceptEntity1{Name: "宋江"}
	result, err = mikilin.Check(value, "name")
	assert.TrueErr(t, result, err)

	// 测试 正常情况
	value = AcceptEntity1{Name: "zhou"}
	result, err = mikilin.Check(value, "name")
	assert.Equal(t, err, "核查错误：属性 Name 的值 zhou 位于禁用值 [zhou] 中", result, false)
}

func TestAccept2(t *testing.T) {
	var value AcceptEntity2
	var result bool
	var err string

	//测试 正常情况
	value = AcceptEntity2{Name: ""}
	result, err = mikilin.Check(value, "name")
	assert.TrueErr(t, result, err)

	// 测试 正常情况
	value = AcceptEntity2{Name: "zhou"}
	result, err = mikilin.Check(value, "name")
	assert.Equal(t, err, "核查错误：属性 Name 的值不为空", result, false)
}

func TestAccept3(t *testing.T) {
	var value AcceptEntity3
	var result bool
	var err string

	//测试 正常情况
	value = AcceptEntity3{Name: ""}
	result, err = mikilin.Check(value, "name")
	assert.TrueErr(t, result, err)

	// 测试 正常情况
	value = AcceptEntity3{Name: "zhou"}
	result, err = mikilin.Check(value, "name")
	assert.TrueErr(t, result, err)

	// 测试 正常情况
	value = AcceptEntity3{Name: "宋江"}
	result, err = mikilin.Check(value, "name")
	assert.Equal(t, err, "[\"核查错误：属性 Name 的值为非空字符\",\"核查错误：属性 Name 的值 宋江 不在只可用列表 [zhou] 中\"]", result, false)
}