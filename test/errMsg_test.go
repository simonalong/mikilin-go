package main

import (
	"testing"
)

type ErrMsgEntity1 struct {
	Name string `match:"value=zhou" errMsg:"对应的值不合法"`
	Age  int
}

type ErrMsgEntity2 struct {
	Name string `match:"value=zhou" errMsg:"当前的值不合法，值为current"`
	Age  int
}

type ErrMsgEntity3 struct {
	Name string `match:"value=zhou" errMsg:"当前的值不合法，值为root"`
	Age  int
}

func TestErrMsg(t *testing.T) {
	// todo
}
