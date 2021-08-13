package times

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

// @Author: lvxiaozheng
// @Date: 2021/2/7 11:39
// @Description:

func TestDateTime(t *testing.T) {
	fmt.Println("日期:" + GetCurrentDay())
	fmt.Println("时间:" + GetCurrentTime())
	fmt.Println("日期时间:" + GetCurrentDayTime())
	fmt.Println("自定义日期时间:" + GetTimeToString(time.Now(), "2006年01月02日 15:04:05"))
	fmt.Println("自定义StringToTime:" + GetStringToTime("2021年01月06日 15:04:05", "2006年01月02日 15:04:05").String())
	fmt.Println("是否今日:" + strconv.FormatBool(IsToday(GetStringToTime("2021年02月07日 15:04:05", "2006年01月02日 15:04:05"))))
	fmt.Println("是否本月:" + strconv.FormatBool(IsCurrentMonth(GetStringToTime("2021年02月07日 15:04:05", "2006年01月02日 15:04:05"))))
	fmt.Println("是否今年:" + strconv.FormatBool(IsCurrentYear(GetStringToTime("2022年02月07日 15:04:05", "2006年01月02日 15:04:05"))))
	fmt.Println("本周周一时间:" + GetFirstDateOfCurrentWeek().String())
	fmt.Println("获取指定日期所属周的周一日期:" + GetFirstDateOfWeek(GetStringToTime("2021年01月02日", "2006年01月02日")).String())
	fmt.Println("传入的时间所在月份的第一天:" + GetFirstDateOfMonth(GetStringToTime("2020年12月28日", "2006年01月02日")).String())
	fmt.Println("传入的时间所在月份的最后一天:" + GetLastDateOfMonth(GetStringToTime("2020年12月28日", "2006年01月02日")).String())
	fmt.Println("上周相同星期日期:" + GetSameWeekdayDate(GetStringToTime("2022年01月03日", "2006年01月02日"), -1).String())
	fmt.Println("------------上月相同星期日期:")
	fmt.Println(GetMonthSameWeekdayDate(GetStringToTime("2021年01月03日", "2006年01月02日"), -1))

	fmt.Println("本月剩余天数:" + strconv.Itoa(GetMonthRemainDay()))
	fmt.Println(GetCompareDays(GetStringToTime("2021年01月04日", "2006年01月02日"), GetStringToTime("2021年02月02日", "2006年01月02日")))
	fmt.Println(GetNanosecondsBetweenDays(GetStringToTime("2021年02月04日 12:33:55", "2006年01月02日 15:04:05"), GetStringToTime("2021年02月02日", "2006年01月02日")))

}

func TestIsHoliday(t *testing.T) {
	dateList := []string{"2021-06-05", "2021-06-06", "2021-06-14"}
	for _, date := range dateList {
		if IsHoliday(date) {
			fmt.Println(date + " 节假日")
		} else {
			fmt.Println(date + " 上班")
		}
	}
}
