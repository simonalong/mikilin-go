package main

import (
	"github.com/simonalong/mikilin-go"
	log "github.com/sirupsen/logrus"
)

type ValueBaseEntityOne struct {
	Name string `match:"value=zhou"`
	Age  int
}

func main() {
	var value ValueBaseEntityOne
	var result bool
	var err string

	value = ValueBaseEntityOne{Name: "zhou"}

	// 核查
	result, err = mikilin.Check(value)
	if !result {
		log.Errorf(err)
	}

	value = ValueBaseEntityOne{Name: "chen"}
	result, err = mikilin.Check(value)
	if !result {
		// 核查错误：属性 Name 的值 chen 不在只可用列表 [zhou] 中
		log.Errorf(err)
	}
}
