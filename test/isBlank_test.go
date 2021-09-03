package main

import (
	mikilin "github.com/SimonAlong/Mikilin-go"
	"github.com/SimonAlong/Mikilin-go/test/assert"
	"testing"
)

type IsBlankBaseEntity struct {
	Name string `match:"isBlank=false"`
	Age  int
}

// 测试基本类型
func TestIsBlankBase(t *testing.T) {
	var value IsBlankBaseEntity
	var result bool
	var err string

	//测试 正常情况
	value = IsBlankBaseEntity{Name: "zhou"}
	result, err = mikilin.Check(value, "name")
	assert.TrueErr(t, result, err)

	// 测试 正常情况
	value = IsBlankBaseEntity{Age: 13}
	result, err = mikilin.Check(value, "name")
	assert.Equal(t, "核查错误：属性 \"Name\" 的值不为空", err, result, false)
}
