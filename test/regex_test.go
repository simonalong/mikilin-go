package main

import (
	mikilin "github.com/SimonAlong/Mikilin-go"
	"github.com/SimonAlong/Mikilin-go/test/assert"
	"testing"
)

type ValueRegexEntity struct {
	Name string `match:"regex=^zhou.*zhen$"`
	Age  int    `match:"regex=^\\d+$"`
}

func TestRegex(t *testing.T) {
	var value ValueRegexEntity
	var result bool
	var err string

	//测试 正常情况
	value = ValueRegexEntity{Name: "zhouOKzhen"}
	result, err = mikilin.Check(value, "name")
	assert.TrueErr(t, result, err)

	//// 测试 正常情况
	value = ValueRegexEntity{Age: 13}
	result, err = mikilin.Check(value, "age")
	assert.Equal(t, "核查错误：属性 Age 的值 13 不是String类型", err, result, false)
}
