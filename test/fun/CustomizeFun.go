package fun

import (
	"fmt"
	"github.com/simonalong/mikilin-go"
	"strconv"
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

type CustomizeEntity5 struct {
	Name string `match:"customize=judge5Name"`
	Age  int
}

type CustomizeEntity6 struct {
	Name1 string `match:"customize=judge6Name_1"`
	Name2 string `match:"customize=judge6Name_2"`
	Name3 string `match:"customize=judge6Name_3"`
	Age   int
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

func JudgeString5(customize CustomizeEntity5) (string, bool) {
	var name = customize.Name
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

func JudgeString6_1(customize CustomizeEntity6, parameterMap map[string]interface{}) (string, bool) {
	nameV := parameterMap["name"]
	name := fmt.Sprintf("%v", nameV)

	ageV := parameterMap["age"]
	age, _ := strconv.Atoi(fmt.Sprintf("%v", ageV))
	if name == "zhou" || name == "宋江" {
		if age > 12 {
			return "", true
		} else {
			return "用户[" + name + "]" + "没有满足年龄age > 12，" + "当前年龄为：" + fmt.Sprintf("%v", age), false
		}

	} else {
		return "没有命中可用的值'zhou'和'宋江'", false
	}
}

func JudgeString6_2(customize CustomizeEntity6, name string, parameterMap map[string]interface{}) (string, bool) {
	ageV := parameterMap["age"]
	age, _ := strconv.Atoi(fmt.Sprintf("%v", ageV))
	if name == "zhou" || name == "宋江" {
		if age > 12 {
			return "", true
		} else {
			return "用户[" + name + "]" + "没有满足年龄age > 12，" + "当前年龄为：" + fmt.Sprintf("%v", age), false
		}
	} else {
		return "没有命中可用的值'zhou'和'宋江'", false
	}
}

func JudgeString6_3(name string, parameterMap map[string]interface{}) (string, bool) {
	ageV := parameterMap["age"]
	age, _ := strconv.Atoi(fmt.Sprintf("%v", ageV))
	if name == "zhou" || name == "宋江" {
		if age > 12 {
			return "", true
		} else {
			return "用户[" + name + "]" + "没有满足年龄age > 12，" + "当前年龄为：" + fmt.Sprintf("%v", age), false
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
	mikilin.RegisterCustomize("judge5Name", JudgeString5)
	mikilin.RegisterCustomize("judge6Name_1", JudgeString6_1)
	mikilin.RegisterCustomize("judge6Name_2", JudgeString6_2)
	mikilin.RegisterCustomize("judge6Name_3", JudgeString6_3)
}
