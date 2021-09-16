package main

//
//import (
//	mikilin "github.com/SimonAlong/Mikilin-go"
//	"github.com/SimonAlong/Mikilin-go/test/assert"
//	"testing"
//)
//
//type CustomizeEntity1 struct {
//	Name string `match:"customize=fun.Judge1"`
//}
//
//type CustomizeEntity2 struct {
//	Name string `match:"customize=fun.Judge2"`
//}
//
//// 身份证号
//func TestCustomize1(t *testing.T) {
//	var value CustomizeEntity1
//	var result bool
//	var err string
//
//	// 测试 正常情况
//	value = CustomizeEntity1{Name: "zhou"}
//	result, err = mikilin.Check(value, "name")
//	assert.True(t, result)
//
//	// 测试 正常情况
//	value = CustomizeEntity1{Name: "宋江"}
//	result, err = mikilin.Check(value, "name")
//	assert.True(t, result)
//
//	// 测试 异常情况
//	value = CustomizeEntity1{Name: "陈真"}
//	result, err = mikilin.Check(value)
//	assert.Equal(t, err, "核查错误：属性 Data1 的值 90 不符合条件 [#current + #root.Data2 > 100] ", result, false)
//}
//
//// 身份证号
//func TestCustomize2(t *testing.T) {
//	var value CustomizeEntity2
//	var result bool
//	var err string
//
//	// 测试 正常情况
//	value = CustomizeEntity2{Name: "zhou"}
//	result, err = mikilin.Check(value, "name")
//	assert.True(t, result)
//
//	// 测试 正常情况
//	value = CustomizeEntity2{Name: "宋江"}
//	result, err = mikilin.Check(value, "name")
//	assert.True(t, result)
//
//	// 测试 异常情况
//	value = CustomizeEntity2{Name: "陈真"}
//	result, err = mikilin.Check(value)
//	assert.Equal(t, err, "核查错误：属性 Data1 的值 90 不符合条件 [#current + #root.Data2 > 100] ", result, false)
//}
