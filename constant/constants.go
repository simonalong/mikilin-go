package constant

/* 匹配 */
const (
	// Value 值列表
	Value = "value"
	// IsBlank 字符为空匹配
	IsBlank = "isBlank"
	// Range 范围匹配
	Range = "range"
	// Model 固定的几个模式匹配
	Model = "model"
	// Condition 条件表达式
	Condition = "condition"
	// Regex 正则表达式
	Regex     = "regex"
	Customize = "customize"
)

/* 匹配后处理 */
const (
	ErrMsg   = "errMsg"
	ChangeTo = "changeTo"
	Accept   = "accept"
	Disable  = "disable"
)

/* tag关键字 */
const (
	EQUAL = "="
	MATCH = "match"
	CHECK = "check"
)

/* range匹配关键字 */
const (
	LeftEqual    = "["
	LeftUnEqual  = "("
	RightUnEqual = ")"
	RightEqual   = "]"

	Now    = "now"
	Past   = "past"
	Future = "future"
)

/* model类别 */
const (
	IdCard     = "id_card"
	Phone      = "phone"
	FixedPhone = "fixed_phone"
	MAIL       = "mail"
	IpAddress  = "ip"
)
