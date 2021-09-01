package matcher

import "fmt"

type BlackWhiteMatch struct {
	BlackMsg string
	WhiteMsg string
}

func (blackWhiteMatch *BlackWhiteMatch) SetBlack(format string, a ...interface{}) {
	blackWhiteMatch.BlackMsg = fmt.Sprintf("核查错误："+format, a)
}

func (blackWhiteMatch *BlackWhiteMatch) SetWhite(format string, a ...interface{}) {
	blackWhiteMatch.WhiteMsg = fmt.Sprintf("核查错误："+format, a)
}
