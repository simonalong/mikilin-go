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

	// 测试 正常情况
	value = ValueRegexEntity{Name: "zhouOKzhen"}
	result, err = mikilin.Check(value, "name")
	assert.TrueErr(t, result, err)

	// 测试 正常情况
	value = ValueRegexEntity{Age: 13}
	result, err = mikilin.Check(value, "age")
	assert.TrueErr(t, result, err)

	// 测试 异常情况
	value = ValueRegexEntity{Name: "chenzhen"}
	result, err = mikilin.Check(value, "name")
	assert.Equal(t, "核查错误：属性 Name 的值 chenzhen 没命中只允许的正则表达式 ^zhou.*zhen$ ", err, result, false)
}
