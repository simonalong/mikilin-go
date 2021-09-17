package fun

import (
	"fmt"
	"github.com/simonalong/mikilin-go"
)

type CustomizeEntity1 struct {
	Name string `match:"customize=judge1Name"`
}

type CustomizeEntity2 struct {
	Name string `match:"customize=judge2Name"`
}

type CustomizeEntity3 struct {
	Name string `match:"customize=judge3Name"`
	Age  int
}

func JudgeString1(name string) bool {
	if name == "zhou" || name == "宋江" {
		return true
	}

	return false
}

func JudgeString2(name string) (bool, string) {
	if name == "zhou" || name == "宋江" {
		return true, ""
	}

	return false, "没有命中可用的值'zhou'和'宋江'"
}

func JudgeString3(customize CustomizeEntity3, name string) (bool, string) {
	if name == "zhou" || name == "宋江" {
		if customize.Age > 12 {
			return true, ""
		} else {
			return false, "用户[" + name + "]" + "没有满足年龄age > 12，" + "当前年龄为：" + fmt.Sprintf("%v", customize.Age)
		}

	} else {
		return false, "没有命中可用的值'zhou'和'宋江'"
	}
}

func init() {
	mikilin.RegisterFun("judge1Name", JudgeString1)
	mikilin.RegisterFun("judge2Name", JudgeString2)
	mikilin.RegisterFun("judge3Name", JudgeString3)
}
