package matcher

import (
	"fmt"
)

type BlackWhiteMatch struct {
	BlackMsg string
	WhiteMsg string
}

func (blackWhiteMatch *BlackWhiteMatch) SetBlackMsg(format string, a ...interface{}) {
	blackWhiteMatch.BlackMsg = fmt.Sprintf("核查错误："+format, a...)
}

func (blackWhiteMatch *BlackWhiteMatch) SetWhiteMsg(format string, a ...interface{}) {
	blackWhiteMatch.WhiteMsg = fmt.Sprintf("核查错误："+format, a...)
}

func (blackWhiteMatch *BlackWhiteMatch) GetWhitMsg() string {
	return blackWhiteMatch.WhiteMsg
}

func (blackWhiteMatch *BlackWhiteMatch) GetBlackMsg() string {
	return blackWhiteMatch.BlackMsg
}
