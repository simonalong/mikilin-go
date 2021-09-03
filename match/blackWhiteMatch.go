package matcher

import (
	"encoding/json"
	"fmt"
)

type BlackWhiteMatch struct {
	BlackMsg string
	WhiteMsg string
}

func (blackWhiteMatch *BlackWhiteMatch) SetBlackStruct(format string, a ...interface{}) {
	var values []interface{}
	for _, data := range a {
		myValue, _ := json.Marshal(data)
		values = append(values, string(myValue))
	}
	blackWhiteMatch.BlackMsg = fmt.Sprintf("核查错误："+format, values...)
}

func (blackWhiteMatch *BlackWhiteMatch) SetWhiteStruct(format string, a ...interface{}) {
	var values []interface{}
	for _, data := range a {
		myValue, _ := json.Marshal(data)
		values = append(values, string(myValue))
	}
	blackWhiteMatch.WhiteMsg = fmt.Sprintf("核查错误："+format, values...)
}

func (blackWhiteMatch *BlackWhiteMatch) SetBlackString(format string, a ...interface{}) {
	blackWhiteMatch.BlackMsg = fmt.Sprintf("核查错误："+format, a...)
}

func (blackWhiteMatch *BlackWhiteMatch) SetWhiteString(format string, a ...interface{}) {
	blackWhiteMatch.WhiteMsg = fmt.Sprintf("核查错误："+format, a...)
}

func (blackWhiteMatch *BlackWhiteMatch) GetWhitMsg() string {
	return blackWhiteMatch.WhiteMsg
}

func (blackWhiteMatch *BlackWhiteMatch) GetBlackMsg() string {
	return blackWhiteMatch.BlackMsg
}
