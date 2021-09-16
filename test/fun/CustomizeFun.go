package fun

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
