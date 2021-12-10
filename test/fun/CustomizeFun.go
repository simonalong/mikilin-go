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

type CustomizeEntity4 struct {
	Name string `match:"customize=judge4Name"`
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
func JudgeString4(customize CustomizeEntity4, name string) (string, bool) {
	if name == "zhou" || name == "宋江" {
		if customize.Age > 12 {
			return "", true
		} else {
			return "用户[" + name + "]" + "没有满足年龄age > 12，" + "当前年龄为：" + fmt.Sprintf("%v", customize.Age), false
		}

	} else {
		return "没有命中可用的值'zhou'和'宋江'", false
	}
}

func JudgeString5(name string, customize CustomizeEntity4) (string, bool) {
	if name == "zhou" || name == "宋江" {
		if customize.Age > 12 {
			return "", true
		} else {
			return "用户[" + name + "]" + "没有满足年龄age > 12，" + "当前年龄为：" + fmt.Sprintf("%v", customize.Age), false
		}

	} else {
		return "没有命中可用的值'zhou'和'宋江'", false
	}
}

func init() {
	mikilin.RegisterCustomize("judge1Name", JudgeString1)
	mikilin.RegisterCustomize("judge2Name", JudgeString2)
	mikilin.RegisterCustomize("judge3Name", JudgeString3)
	mikilin.RegisterCustomize("judge4Name", JudgeString4)
}
