package constant

/* 匹配 */
const (
	// EQUAL 值列表
	EQUAL = "="
	// VALUE 值列表
	VALUE = "value"
	// IsBlank 字符为空匹配
	IsBlank = "isBlank"
	// RANGE 范围匹配
	RANGE = "range"
	// MODEL 模式类型匹配
	MODEL     = "model"
	CONDITION = "condition"
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
