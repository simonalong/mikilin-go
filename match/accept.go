package matcher

import (
	"github.com/simonalong/mikilin-go/constant"
	"reflect"
	"strconv"
	"strings"
)

func CollectAccept(objectTypeFullName string, fieldKind reflect.Kind, objectFieldName string, tagName string, subCondition string, errMsg string) {
	if constant.Accept != tagName {
		return
	}

	accept, err := strconv.ParseBool(strings.TrimSpace(subCondition))
	if err != nil {
		return
	}
	addMatcher(objectTypeFullName, objectFieldName, nil, errMsg, accept)
}
