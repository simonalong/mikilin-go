package main

import (
	mikilin "github.com/simonalong/mikilin-go"
	"testing"
)

type DisableEntity1 struct {
	Name string `match:"value=zhou" disable:"true"`
	Age  int
}

func TestDisable1(t *testing.T) {
	var value DisableEntity1
	var result bool
	var err string

	//测试 正常情况
	value = DisableEntity1{Name: "zhou"}
	result, err = mikilin.Check(value, "name")
	TrueErr(t, result, err)

	// 测试 正常情况
	value = DisableEntity1{Name: "chenzhen"}
	result, err = mikilin.Check(value, "name")
	TrueErr(t, result, err)
}
