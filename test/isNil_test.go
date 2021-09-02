package main

import (
	mikilin "github.com/SimonAlong/Mikilin-go"
	"github.com/SimonAlong/Mikilin-go/test/assert"
	"testing"
)

type ValueIsNilEntity struct {
	Name string `match:"isNil=false"`
	Age  int
}

// 测试基本类型
func TestIsNilBase(t *testing.T) {
	var value ValueIsNilEntity
	var result bool
	var err string

	//测试 正常情况
	value = ValueIsNilEntity{Age: 12}
	result, err = mikilin.Check(value, "name")
	assert.FalseErr(t, result, err)
}
