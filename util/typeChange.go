package util

import "strconv"

//将可能是int\int32\int64\float64\string类型的interface{}转换成string，可继续完善类型
func InterfaceToString(i interface{}) string {
	switch i.(type) {
	case int:
		return strconv.Itoa(i.(int))
	case int32:
		return strconv.Itoa(int(i.(int32)))
	case int64:
		return strconv.FormatInt(i.(int64), 10)
	case float64:
		return strconv.FormatFloat(i.(float64), 'f', -1, 64)
	case string:
		return i.(string)
	default:
		return "未知类型"
	}
}
