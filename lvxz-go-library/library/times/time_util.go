package times

import (
	"fmt"
	"go-basiclib/library/utils"
	"strconv"
	"strings"
	"time"
)

// @Author: lvxiaozheng
// @Date: 2021/6/21 11:05
// @Description: 时间工具类

//获取7天前日期
func GetLastWeek() string {
	return GetTimeToString(time.Now().AddDate(0, 0, -7), Format1)
}

// GetStringToTime: 字符串转时间
func GetTime(timeString string, format string) time.Time {
	t, _ := time.Parse(format, timeString)
	return t
}

//获取当前时间点前一个小时
func GetOnePreviousHour() time.Time {
	h1, _ := time.ParseDuration("-1h")
	return time.Now().Add(h1)
}

//自动添加''
func FmtCol(v string) string {
	return fmt.Sprintf("'%s'", strings.Trim(v, " "))
}

//自动时间格式补位(添加'')
func FmtTimeCol(v int) string {
	return FmtCol(FmtTimeFormat(v))
}

//自动时间格式补位
func FmtTimeFormat(v int) string {
	val := strconv.Itoa(v)
	if len(val) == 1 {
		//自动补位0
		val = "0" + val
	}
	return val
}

//获取递减demoteHourVal小时
func GetSubtractHourTimeList(nTime time.Time, demoteHourVal int) []string {
	var timeList []string
	for begin := 0; begin < demoteHourVal; begin++ {
		h0 := fmt.Sprintf("-%dh", begin)
		h1, _ := time.ParseDuration(h0)
		h2 := nTime.Add(h1)
		timeList = append(timeList, FmtTimeCol(h2.Hour()))
	}
	return timeList
}

//获取递减demoteMinuteVal分钟(只支持当天递减)
func GetNewSubtractMinuteTimeList(nTime time.Time, demoteMinuteVal int) ([]string, []string) {
	var hourList []string
	var minuteList []string
	for begin := 0; begin < demoteMinuteVal; begin++ {
		m0 := fmt.Sprintf("-%dm", begin)
		m1, _ := time.ParseDuration(m0)
		m2 := nTime.Add(m1)

		//判断当天时间
		if !IsToday2(GetTimeToString(m2, Format1)) {
			break
		}

		if !utils.SliceContain(hourList, FmtTimeCol(m2.Hour())) {
			hourList = append(hourList, FmtTimeCol(m2.Hour()))
		}

		if !utils.SliceContain(minuteList, FmtTimeCol(m2.Minute())) {
			minuteList = append(minuteList, FmtTimeCol(m2.Minute()))
		}
	}
	return hourList, minuteList
}

//获取递减demoteMinuteVal分钟(支持跨天递减)
func GetSubtractMinuteTimeList(nTime time.Time, demoteMinuteVal int) ([]string, []string) {
	var hourList []string
	var minuteList []string
	for begin := 0; begin < demoteMinuteVal; begin++ {
		m0 := fmt.Sprintf("-%dm", begin)
		m1, _ := time.ParseDuration(m0)
		m2 := nTime.Add(m1)

		if !utils.SliceContain(hourList, FmtTimeCol(m2.Hour())) {
			hourList = append(hourList, FmtTimeCol(m2.Hour()))
		}

		if !utils.SliceContain(minuteList, FmtTimeCol(m2.Minute())) {
			minuteList = append(minuteList, FmtTimeCol(m2.Minute()))
		}
	}
	return hourList, minuteList
}

//获取整除5的时间(0,5,10,15,20...)
func Get5MinuteTimeList(hour, minute int) []string {
	return getMinuteTimeList(hour, minute, 5)
}

//获取00,15,30,45分钟间隔时间
func Get15MinuteTimeList(hour, minute int) []string {
	return getMinuteTimeList(hour, minute, 15)
}

//获取指定时间间隔的时间
func getMinuteTimeList(hour, minute, intervalTime int) []string {
	var timeList []string
	for begin := 0; begin <= hour; begin++ {
		nHour := FmtTimeFormat(begin)

		if begin == hour {
			//获取当前小时的满足间隔
			for mBegin := 0; mBegin <= minute; mBegin = mBegin + intervalTime {
				nMinute := FmtTimeFormat(mBegin)
				timeList = append(timeList, fmt.Sprintf("%s:%s", nHour, nMinute))
			}
		} else {
			//获取该小时的所有间隔
			for mBegin := 0; mBegin <= 59; mBegin = mBegin + intervalTime {
				nMinute := FmtTimeFormat(mBegin)
				timeList = append(timeList, fmt.Sprintf("%s:%s", nHour, nMinute))
			}
		}
	}
	return timeList
}

//获取1小时间隔时间
func Get1HourTimeList(hour int) []string {
	var timeList []string
	for begin := 0; begin <= hour; begin++ {
		nHour := FmtTimeFormat(begin)
		timeList = append(timeList, fmt.Sprintf("%s:00", nHour))
	}
	return timeList
}

//获取本月日期时间集合
func GetThisMonthDateTimeList(logDate string) []string {
	nTime := GetStringToTime2(logDate)
	return GetSubtractDateTimeList(nTime, nTime.Day())
}

//获取递减日期集合
func GetSubtractDateTimeList(nTime time.Time, demoteDate int) []string {
	var timeList []string
	for begin := 0; begin < demoteDate; begin++ {
		//递减日期
		timeList = append(timeList, GetTimeToString(nTime.AddDate(0, 0, -begin), Format1))
	}
	return timeList
}
