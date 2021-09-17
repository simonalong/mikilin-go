package matcher

import (
	"github.com/SimonAlong/Mikilin-go/constant"
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
