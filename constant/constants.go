package constant

/* 匹配 */
const (
	// VALUE 值列表
	VALUE = "value"
	// IsBlank 字符为空匹配
	IsBlank = "isBlank"
	// RANGE 范围匹配
	RANGE     = "range"
	MODEL     = "model"
	CONDITION = "condition"
	// REGEX 正则表达式
	REGEX     = "regex"
	CUSTOMIZE = "customize"
)

/* 匹配后处理 */
const (
	ERR_MSG   = "errMsg"
	CHANGE_TO = "changeTo"
	ACCEPT    = "accept"
	DISABLE   = "disable"
)

/* tag关键字 */
const (
	EQUAL = "="
	MATCH = "match"
	CHECK = "check"
)

/* range匹配关键字 */
const (
	LEFT_EQUAL     = "["
	LEFT_UN_EQUAL  = "("
	RIGHT_UN_EQUAL = ")"
	RIGHT_EQUAL    = "]"

	NOW    = "now"
	PAST   = "past"
	FUTURE = "future"
)

/* model类别 */
const (
	ID_CARD     = "id_card"
	PHONE       = "phone"
	FIXED_PHONE = "fixed_phone"
	MAIL        = "mail"
	IP_ADDRESS  = "ip"
)
