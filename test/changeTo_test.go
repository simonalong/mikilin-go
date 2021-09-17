package main

import (
	mikilin "github.com/SimonAlong/Mikilin-go"
	"github.com/SimonAlong/Mikilin-go/test/assert"
	"testing"
)

// 测试基本表达式
type ChangeToEntity1 struct {
	Name string `match:"value={zhou, 宋江}" changeTo:"chenzhen"`
}

func TestChangeTo1(t *testing.T) {
	var value ChangeToEntity1
	var result bool

	value = ChangeToEntity1{Name: "zhou"}
	result, _ = mikilin.Check(value)
	assert.Equal(t, result, true, value.Name, "chenzhen")
	//
	//value = ChangeToEntity1{Name: "宋江"}
	//result, _ = mikilin.Check(value)
	//assert.Equal(t, true, result, "chenzhen", value.Name)
	//
	//value = ChangeToEntity1{Name: "陈真"}
	//result, _ = mikilin.Check(value)
	//assert.Equal(t, true, result, "陈真", value.Name)
}
