package matcher

import (
	log "github.com/sirupsen/logrus"
	"reflect"
	"regexp"
)

type RangeMatch struct {
	BlackWhiteMatch

	beginAli string
	begin    interface{}
	end      interface{}
	endAli   string
}

type RangeEntity struct {
	beginAli    string
	begin       interface{}
	end         interface{}
	endAli      string
	dateFlag    bool
	dynamicTime bool
}

type Predicate func(subCondition string) bool

func (rangeMatch *RangeMatch) Match(object interface{}, field reflect.StructField, fieldValue interface{}) bool {

	return false
}

func (rangeMatch *RangeMatch) IsEmpty() bool {
	// todo
	return false
}

/*
 * []：范围匹配
 */
var rangeRegex = regexp.MustCompile("^(\\(|\\[)(.*),(\\s)*(.*)(\\)|\\])$")

// digitRegex 全是数字匹配（整数，浮点数，0，负数）
var digitRegex = regexp.MustCompile("^[-+]?(0)|([1-9]+\\d*|0\\.(\\d*)|[1-9]\\d*\\.(\\d*))$")

/*
 * 时间或者数字范围匹配
 */
//private Pattern rangePattern = Pattern.compile("^(\\(|\\[)(.*),(\\s)*(.*)(\\)|\\])$");

func BuildRangeMatcher(objectTypeFullName string, fieldKind reflect.Kind, objectFieldName string, subCondition string) {
	subData := rangeRegex.FindAllStringSubmatch(subCondition, 1)
	if len(subData) > 0 {
		//beginAli := subData[0][1]
		begin := subData[0][2]
		end := subData[0][4]
		//endAli := subData[0][5]

		if (begin == "nil" || begin == "") && (end == "nil" || end == "") {
			log.Errorf("range匹配器格式输入错误，start和end不可都为null或者空字符, input=%v", subCondition)
			return
		} else if begin == "past" || begin == "future" {
			log.Errorf("range匹配器格式输入错误, start不可含有past或者future, input=%v", subCondition)
			return
		} else if end == "past" || end == "future" {
			log.Errorf("range匹配器格式输入错误, end不可含有past或者future, input=%v", subCondition)
			return
		}

		// 如果是数字，则按照数字解析
		if digitRegex.MatchString(begin) || digitRegex.MatchString(end) {

		}

		// 如果是解析动态时间，按照动态时间解析

		// 否则按照时间解析

	} else {
		// 匹配过去和未来的时间

		//if (input.equals(PAST) || input.equals(FUTURE)) {
		//	return parseRangeDate(input);
		//}
	}
	//reg1.SubexpNames()
	//
	//input = input.trim();
	//java.util.regex.Matcher matcher = rangePattern.matcher(input);
	//if matcher.find() {
	//	String beginAli = matcher.group(1);
	//	String begin = matcher.group(2);
	//	String end = matcher.group(4);
	//	String endAli = matcher.group(5);
	//
	//	if ((begin.equals(NULL_STR) || "".equals(begin)) && (end.equals(NULL_STR) || "".equals(end))) {
	//		log.error(MK_LOG_PRE + "range匹配器格式输入错误，start和end不可都为null或者空字符, input={}", input);
	//	} else if (begin.equals(PAST) || begin.equals(FUTURE)) {
	//		log.error(MK_LOG_PRE + "range匹配器格式输入错误, start不可含有past或者future, input={}", input);
	//	} else if (end.equals(PAST) || end.equals(FUTURE)) {
	//		log.error(MK_LOG_PRE + "range匹配器格式输入错误, end不可含有past或者future, input={}", input);
	//	}
	//
	//	// 如果是数字，则按照数字解析
	//	if (digitPattern.matcher(begin).matches() || digitPattern.matcher(end).matches()) {
	//		return RangeEntity.build(beginAli, parseNum(begin), parseNum(end), endAli, false);
	//	} else if (timePlusPattern.matcher(begin).matches() || timePlusPattern.matcher(end).matches()) {
	//		// 解析动态时间
	//		DynamicTimeNum timeNumBegin = parseDynamicTime(begin);
	//		DynamicTimeNum timeNumEnd = parseDynamicTime(end);
	//
	//		if (null != timeNumBegin && null != timeNumEnd && timeNumBegin.compareTo(timeNumEnd) > 0) {
	//			log.error(MK_LOG_PRE + "时间的动态时间不正确，动态起点时间不应该大于动态终点时间");
	//			return null;
	//		}
	//
	//		if (null == timeNumBegin && null == timeNumEnd) {
	//			log.error(MK_LOG_PRE + "动态时间解析失败");
	//			return null;
	//		}
	//		return RangeEntity.build(beginAli, timeNumBegin, timeNumEnd, endAli);
	//	} else {
	//		Date beginDate = parseDate(begin);
	//		Date endDate = parseDate(end);
	//		if (null != beginDate && null != endDate) {
	//			if (beginDate.compareTo(endDate) > 0) {
	//				log.error(MK_LOG_PRE + "时间的范围起始点不正确，起点时间不应该大于终点时间");
	//				return null;
	//			}
	//			return RangeEntity.build(beginAli, LocalDateTimeUtil.dateToLong(beginDate), LocalDateTimeUtil.dateToLong(endDate), endAli, true);
	//		} else if (null == beginDate && null == endDate) {
	//			log.error(MK_LOG_PRE + "range 匹配器格式输入错误，解析数字或者日期失败, input={}", input);
	//		} else {
	//			return RangeEntity.build(beginAli, LocalDateTimeUtil.dateToLong(beginDate), LocalDateTimeUtil.dateToLong(endDate), endAli, true);
	//		}
	//		return null;
	//	}
	//} else {
	//	// 匹配过去和未来的时间
	//	if (input.equals(PAST) || input.equals(FUTURE)) {
	//		return parseRangeDate(input);
	//	}
	//}
	//return null;
}
