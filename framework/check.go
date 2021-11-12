package lf

import "time"

//校验参数是否为空
func CheckParamsEmpty(params ...string) bool {
	for _, value := range params {
		if value == "" {
			return false
		}
	}
	return true
}

//校验时间
func CheckS14Time(timeStr string, minute float64) bool {
	t, err := Str14ToTime(timeStr)
	if err != nil {
		return false
	}

	f := t.Sub(time.Now()).Minutes()
	if f > minute || f < -minute {
		return false
	}
	return true
}
