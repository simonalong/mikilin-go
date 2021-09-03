package util

import (
	"reflect"
	"strconv"
)

// 判断是否是核查的类型
func IsCheckedKing(fieldKing reflect.Kind) bool {
	switch fieldKing {
	case reflect.Int:
		return true
	case reflect.Int8:
		return true
	case reflect.Int16:
		return true
	case reflect.Int32:
		return true
	case reflect.Int64:
		return true
	case reflect.Uint:
		return true
	case reflect.Uint8:
		return true
	case reflect.Uint16:
		return true
	case reflect.Uint32:
		return true
	case reflect.Uint64:
		return true
	case reflect.Float32:
		return true
	case reflect.Float64:
		return true
	case reflect.Bool:
		return true
	case reflect.String:
		return true
	default:
		return false
	}
}

func Cast(fieldKind reflect.Kind, valueStr string) (interface{}, error) {
	switch fieldKind {
	case reflect.Int:
		return strconv.Atoi(valueStr)
	case reflect.Int8:
		return strconv.ParseInt(valueStr, 10, 8)
	case reflect.Int16:
		return strconv.ParseInt(valueStr, 10, 16)
	case reflect.Int32:
		return strconv.ParseInt(valueStr, 10, 32)
	case reflect.Int64:
		return strconv.ParseInt(valueStr, 10, 64)
	case reflect.Uint:
		return strconv.ParseUint(valueStr, 10, 0)
	case reflect.Uint8:
		return strconv.ParseUint(valueStr, 10, 8)
	case reflect.Uint16:
		return strconv.ParseUint(valueStr, 10, 16)
	case reflect.Uint32:
		return strconv.ParseUint(valueStr, 10, 32)
	case reflect.Uint64:
		return strconv.ParseUint(valueStr, 10, 64)
	case reflect.Float32:
		return strconv.ParseFloat(valueStr, 32)
	case reflect.Float64:
		return strconv.ParseFloat(valueStr, 64)
	case reflect.Bool:
		return strconv.ParseBool(valueStr)
	}
	return valueStr, nil
}
