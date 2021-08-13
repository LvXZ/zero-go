package times

import (
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// @Author: lvxiaozheng
// @Date: 2021/2/7 11:31
// @Description: 日期时间处理工具类

const (
	Format1 = "20060102"

	Format2 = "15:04:05"

	Format3 = "2006-01-02 15:04:05"

	Format4 = "2006-01-02"
)

// GetCurrentDay: 获取当天日期
func GetCurrentDay() string {
	return GetTimeToString(time.Now(), Format1)
}

// GetCurrentTime: 获取当前时间
func GetCurrentTime() string {
	return GetTimeToString(time.Now(), Format2)
}

// GetCurrentDayTime: 获取当前日期时间
func GetCurrentDayTime() string {
	return GetTimeToString(time.Now(), Format3)
}

// GetYesterday: 获取昨天日期
func GetYesterday(format string) string {
	return GetTimeToString(time.Now().AddDate(0, 0, -1), format)
}

// GetTomorrow: 获取明天日期
func GetTomorrow(format string) string {
	return GetTimeToString(time.Now().AddDate(0, 0, 1), format)
}

// GetFormatDayTime: 获取自定义日期时间
func GetTimeToString(getTime time.Time, format string) string {
	return getTime.Format(format)
}

// GetStringToTime: 字符串转时间
func GetStringToTime(timeString string, format string) time.Time {
	t, _ := time.Parse(format, timeString)
	return t
}

// IsToday: 判断是否今天 true是 false否
func IsToday(getTime time.Time) bool {
	if IsCurrentYear(getTime) && time.Now().YearDay() == getTime.YearDay() {
		return true
	}
	return false
}

// GetStringToTime: 字符串转时间
func GetStringToTime2(timeString string) time.Time {
	if len(timeString) == 0 {
		return time.Now()
	} else if len(timeString) == 8 {
		//日期8位的时间格式：20210321
		return GetTime(timeString, Format1)
	} else if len(timeString) == 10 {
		//日期10位的时间格式：2021-03-21
		return GetTime(timeString, Format4)
	} else {
		return time.Now()
	}
}

//是否今天
func IsToday2(logDate string) bool {
	if len(logDate) == 0 {
		return false
	} else if len(logDate) == 8 {
		//日期8位的时间格式：20210321
		return strings.EqualFold(time.Now().Format(Format1), logDate)
	} else if len(logDate) == 10 {
		//日期10位的时间格式：2021-03-21
		return strings.EqualFold(time.Now().Format(Format4), logDate)
	} else {
		return false
	}
}

// IsCurrentMonth: 判断是否本月 true是 false否
func IsCurrentMonth(getTime time.Time) bool {
	if IsCurrentYear(getTime) && time.Now().Month() == getTime.Month() {
		return true
	}
	return false
}

// IsCurrentYear: 判断是否今年 true是 false否
func IsCurrentYear(getTime time.Time) bool {
	if time.Now().Year() == getTime.Year() {
		return true
	}
	return false
}

// GetFirstDateOfCurrentWeek: 获取本周周一的日期
func GetFirstDateOfCurrentWeek() time.Time {
	return GetFirstDateOfWeek(time.Now())
}

// GetFirstDateOfWeek: 获取指定日期所属周的周一日期
func GetFirstDateOfWeek(getTime time.Time) time.Time {
	offset := int(time.Monday - getTime.Weekday())
	if offset > 0 {
		offset = -6
	}

	return time.Date(getTime.Year(), getTime.Month(), getTime.Day(), 0, 0, 0, 0, getTime.Location()).AddDate(0, 0, offset)
}

// GetZeroTime: 获取某一天的0点时间
func GetZeroTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}

// GetFirstDateOfMonth: 获取传入的时间所在月份的第一天，即某月第一天的0点。如传入time.Now(), 返回当前月份的第一天0点时间。
func GetFirstDateOfMonth(d time.Time) time.Time {
	d = d.AddDate(0, 0, -d.Day()+1)
	return GetZeroTime(d)
}

// GetLastDateOfMonth: 获取传入的时间所在月份的最后一天，即某月最后一天的0点。如传入time.Now(), 返回当前月份的最后一天0点时间。
func GetLastDateOfMonth(d time.Time) time.Time {
	return GetFirstDateOfMonth(d).AddDate(0, 1, -1)
}

// GetSameWeekdayDate: 获取相同星期的日期 week:1 -下周  -1 -上周
func GetSameWeekdayDate(getTime time.Time, week int) time.Time {
	return getTime.AddDate(0, 0, 7*week)
}

// GetMonthSameWeekdayDate: 获取某月的相同星期日期 month:1 -下月  0 -本月  -1 -上月
func GetMonthSameWeekdayDate(getTime time.Time, month int) []time.Time {
	getWeekday := getTime.Weekday()

	// 获取指定月份的第一天--时间直接AddDate月份是错误的
	firstDateOfMonth := GetFirstDateOfMonth(getTime).AddDate(0, month, 0)

	var firstWeekdayOfMonth time.Time
	for i := 0; i < 7; i++ {
		firstWeekday := firstDateOfMonth.AddDate(0, 0, i)
		if getWeekday == firstWeekday.Weekday() {
			firstWeekdayOfMonth = firstWeekday
			break
		}
	}

	var sameWeekdayList []time.Time
	for i := 0; i < 5; i++ {
		day := firstWeekdayOfMonth.AddDate(0, 0, i*7)
		if firstDateOfMonth.Month() != day.Month() {
			break
		}
		sameWeekdayList = append(sameWeekdayList, day)
	}

	return sameWeekdayList
}

// GetMonthRemainDay: 本月剩余天数
func GetMonthRemainDay() int {
	getTime := time.Now()
	// 获取本月最后一天
	lastDate := GetLastDateOfMonth(getTime)

	return lastDate.YearDay() - getTime.YearDay()
}

// GetCompareNow: 计算指定时间到现在相差的天数
func GetCompareNow(getTime time.Time) int {
	return GetCompareDays(getTime, time.Now())
}

// GetCompareDays: 计算指定两天时间相差的天数
func GetCompareDays(getTime1 time.Time, getTime2 time.Time) int {
	if getTime1.After(getTime2) {
		return int(getTime1.Sub(getTime2).Hours()) / 24
	} else {
		return int(getTime2.Sub(getTime1).Hours()) / 24
	}
}

// GetNanosecondsBetweenDays: 计算指定两天时间的时间差 单位纳秒(除以10的9次方单位为秒)
func GetNanosecondsBetweenDays(getTime1 time.Time, getTime2 time.Time) int64 {
	if getTime1.After(getTime2) {
		return getTime1.Sub(getTime2).Nanoseconds()
	} else {
		return getTime2.Sub(getTime1).Nanoseconds()
	}
}

// IsHoliday: 判断日期是否节假日 true是 false否
func IsHoliday(date string) bool {

	queryUrl := "http://api.tianapi.com/txapi/jiejiari/index?key=c5830184c462e99e12caa40118aab14a&date=" + date

	req, err := http.NewRequest(http.MethodGet, queryUrl, nil)
	if err != nil || req == nil {
		return false
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil || res == nil || res.StatusCode != 200 {
		return false
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil || body == nil {
		return false
	}

	responseData := string(body)
	fmt.Println(responseData)

	newsList := gjson.Get(responseData, "newslist")
	fmt.Println(newsList)

	isNotWork := newsList.Array()[0].Get("isnotwork")
	fmt.Println(isNotWork)
	if strings.Compare(isNotWork.String(), "1") == 0 {
		return true
	}
	return false
}
