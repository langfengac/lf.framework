package lf

import (
	"errors"
	"time"
)

const (
	TimeF               = "2006-01-02 15:04:05"
	TimeCnF             = "2006年1月2日 15:04:05"
	TimeyyyyMMddHHmmssF = "20060102150405"
)

var CstZone *time.Location //东八区

func init() {
	CstZone = time.FixedZone("CST", 8*3600) //东八区
}

func TimeToString(t time.Time) string {
	return t.Format(TimeF)
}
func TimeToStr14(t time.Time) string {
	return t.Format(TimeyyyyMMddHHmmssF)
}

//yyyyMMddHHmmss转Time
func Str14ToTime(timeStr string) (time.Time, error) {
	//time转数组
	arr := make([]string, len(timeStr))
	for i, v := range []rune(timeStr) {
		arr[i] = string(v)
	}
	if len(arr) != 14 {
		return time.Unix(0, 0), errors.New("长度不为14")
	}
	//yyyyMMddHHmmss
	temp := arr[0] + arr[1] + arr[2] + arr[3] + "-" + arr[4] + arr[5] + "-" + arr[6] + arr[7] + " " + arr[8] + arr[9] + ":" + arr[10] + arr[11] + ":" + arr[12] + arr[13]
	t, err := time.ParseInLocation(TimeF, temp, time.Local)
	if err != nil {
		return time.Unix(0, 0), err
	}
	return t, nil
}
