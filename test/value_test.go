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

func TestValue1(t *testing.T) {
	// 测试 正常情况
	value := ValueBaseEntity{Age: 12}
	result, err := mikilin.Check(value, "age")
	assert.AssertTrueErr(t, result, err)

	// 测试 正常情况
	value = ValueBaseEntity{Age: 13}
	result, err = mikilin.Check(value, "age")
	assert.AssertTrueErr(t, result, err)

	// 测试 异常情况
	value = ValueBaseEntity{Age: 14}
	result, err = mikilin.Check(value, "age")
	assert.AssertFalseErr(t, result, err)

	// 测试 正常情况
	value = ValueBaseEntity{Name: "zhou"}
	result, err = mikilin.Check(value, "name")
	assert.AssertTrueErr(t, result, err)

	// 测试 正常情况
	value = ValueBaseEntity{Name: "宋江"}
	result, err = mikilin.Check(value, "name")
	assert.AssertTrueErr(t, result, err)

	// 测试 异常情况
	value = ValueBaseEntity{Name: "陈真"}
	result, err = mikilin.Check(value, "name")
	assert.AssertFalseErr(t, result, err)
}
