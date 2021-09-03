package matcher

import "reflect"

type RangeMatch struct {
	BlackWhiteMatch

	PredicateRange Predicate
}

type Predicate func(subCondition string) bool

func (rangeMatch *RangeMatch) Match(object interface{}, field reflect.StructField, fieldValue interface{}) bool {

	return false
}

func (rangeMatch *RangeMatch) IsEmpty() bool {
	// todo
	return false
}

func BuildRangeMatcher(objectTypeFullName string, fieldKind reflect.Kind, objectFieldName string, subCondition string) {
	// todo
}
