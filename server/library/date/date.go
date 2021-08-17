package date

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
	"github.com/jinzhu/now"
	"hess/library/validator"
	"strconv"
	"time"
)

// GetMonthDays 查询指定年份指定月份有多少天
// @params year int 指定年份
// @params month int 指定月份
func GetMonthDays(date string) int {
	var (
		year    int
		month   int
		dateArr g.SliceStr
	)
	if !validator.IsDate(date) && !validator.IsMonth(date) {
		return 0
	} else {
		dateArr = gstr.Split(date, "-")
		year = gconv.Int(dateArr[0])
		month, _ = strconv.Atoi(dateArr[1])
	}
	// 有31天的月份
	day31 := map[int]bool{
		1:  true,
		3:  true,
		5:  true,
		7:  true,
		8:  true,
		10: true,
		12: true,
	}
	aval, aok := day31[month]
	if aok && aval {
		return 31
	}
	// 有30天的月份
	day30 := map[int]bool{
		4:  true,
		6:  true,
		9:  true,
		11: true,
	}
	bval, bok := day30[month]
	if bval == true && bok {
		return 30
	}
	// 计算是平年还是闰年
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		// 得出2月的天数
		return 29
	}
	// 得出2月的天数
	return 28
}
func GetMonthStartDate(date ...string) time.Time {
	var value string
	if len(date) == 0 {
		value = gtime.Now().Format("Y-m")
	} else {
		value = date[0]
	}
	t, err := now.Parse(value)
	if err != nil {
		return time.Now()
	}
	return now.With(t).BeginningOfMonth()
}
func GetMonthEndDate(date ...string) time.Time {
	var value string
	if len(date) == 0 {
		value = gtime.Now().Format("Y-m-d")
	}
	t, err := now.Parse(value)
	if err != nil {
		return time.Now()
	}
	return now.With(t).EndOfMonth()
}

// GetDiffDays 获取两个日期相关几天
func GetDiffDays(start, end string) int {
	startTime := gtime.NewFromStr(start)
	endTime := gtime.NewFromStr(end)
	diff := endTime.Sub(startTime)
	days := diff / 86400
	return int(days)
}

// GetDiffYear 获取两个日期相差几年
// 如，date.GetDiffYear("2019-01-01", "2021-8-01")
func GetDiffYear(start, end string) int {
	startTime := gtime.NewFromStr(start)
	endTime := gtime.NewFromStr(end)
	diff := endTime.Sub(startTime)
	return int(diff.Seconds() / 86400 / 365)
}

// GetDiffMonth 获取两个日期相差几个月
// 如，date.GetDiffMonth("2021-01-01", "2021-8-01")
func GetDiffMonth(start, end string) int {
	startTime := gtime.NewFromStr(start)
	endTime := gtime.NewFromStr(end)
	months := 0
	month := startTime.Month()
	for startTime.Before(endTime) {
		startTime = startTime.Add(time.Hour * 24)
		nextMonth := startTime.Month()
		if nextMonth != month {
			months++
		}
		month = nextMonth
	}

	return months
}
