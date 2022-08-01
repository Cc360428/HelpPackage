package timec

import (
	"github.com/Cc360428/HelpPackage/other"
	"time"
)

var (
	monthDay = map[int]int{
		1:  31,
		2:  28,
		3:  31,
		4:  30,
		5:  31,
		6:  30,
		7:  31,
		8:  31,
		9:  30,
		10: 31,
		11: 30,
		12: 31,
		13: 29, //表示润年二月
	}
)

// IsLeapYear ...
// @Description: 是否润年
// @param year
// @return bool
func IsLeapYear(year int) bool {
	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		return true
	}
	return false
}

// IsMonthLastDay ...
// @Description: 是否每个月最后一天
// @param year
// @param month
// @param day
// @return bool
func IsMonthLastDay(year, month, day int) bool {
	if IsLeapYear(year) && month == 2 {
		month = 13
	}

	if monthDay[month] == day {
		return true
	}

	return false
}

// IsToday 判断时间戳是否是属于今天，秒级时间戳
func IsToday(ts int64) bool {

	y, m, d := time.Unix(ts, 0).Date()
	y1, m1, d1 := time.Now().Date()

	if d != d1 || m != m1 || y != y1 {
		return false
	}

	return true
}

// DateDay 获取当前的日
func DateDay() (day int) {
	now := time.Now()
	day = now.Day()
	return day
}

// GetWeek 获取今天周
func GetWeek() (week int64) {
	t := time.Now()
	i := int(t.Weekday())
	week, _ = other.ToInt64(i)
	return week
}

// TodayZeroTs 获取今天0点时间戳
func TodayZeroTs() int64 {
	y, m, d := time.Now().Date()
	ts := time.Date(y, m, d, 0, 0, 0, 0, time.Local).Unix()
	return ts
}

// NextDayZeroTs 获取次日0点时间戳
func NextDayZeroTs() int64 {
	y, m, d := time.Now().Date()
	ts := time.Date(y, m, d, 0, 0, 0, 0, time.Local).
		AddDate(0, 0, 1).Unix()
	return ts
}

// TodayLastSec 今天剩余多少秒 用于设置0点过期
func TodayLastSec() int64 {
	return NextDayZeroTs() - time.Now().Unix()
}

// DateDayZero 获取当日零时时间戳（秒）
//Time stamp of the morning of the day: unit seconds
func DateDayZero() (r int64) {
	timeStr := time.Now().Format("2006-01-02")
	//使用Parse 默认获取为UTC时区 需要获取本地时区 所以使用ParseInLocation
	t2, _ := time.ParseInLocation("2006-01-02", timeStr, time.Local)
	r = t2.AddDate(0, 0, 0).Unix()
	return r
}

// DateDayMilliseconds Time stamp for the morning of the day: in milliseconds
// DateDayMilliseconds 获取当日凌晨时间戳（毫秒）
func DateDayMilliseconds() (r int64) {
	t := time.Now()
	zeroTm := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).Unix()
	r = zeroTm * 1000
	return r
}

// WeekOneAndWeekSevenTime 当前周一凌晨和周日最后时间戳
func WeekOneAndWeekSevenTime() (one int64, seven int64) {
	d := int64(86400000)
	t := time.Now()
	w := int64(t.Weekday()) - 1
	num := d * w
	m := DateDayMilliseconds()
	one = m - num
	seven = one + (d * 7)
	return one, seven
}

func DateDayFormat() string {
	//timeStr := time.Now().Format("2006-01-02 15:04:05.000")
	//timeStr := time.Now().Format("2006-01-02 15:04:05.000")
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	return timeStr
}

func DayMi() (rr int64) {
	return time.Now().UnixNano() / 1e6
}

// UnixTimeInt64 返回值int64类型
//时间戳（秒）： time.Now().Unix())
//时间戳（纳秒）： time.Now().UnixNano())
//时间戳（毫秒）： time.Now().UnixNano()/1e6)
//时间戳（纳秒转换为秒） time.Now().UnixNano()/1e9)
func UnixTimeInt64(q int64) (p int64) {
	t := time.Now()
	if q == 1 { //秒
		second := t.Unix()
		return second
	} else if q == 2 { //纳秒
		nanosecond := t.UnixNano()
		return nanosecond
	} else if q == 3 { //毫秒
		millisecond := t.UnixNano() / 1e6
		return millisecond
	}
	return 0
}

// UnixTimeString 返回值string类型
//时间戳（秒）： time.Now().Unix())
//时间戳（纳秒）： time.Now().UnixNano())
//时间戳（毫秒）： time.Now().UnixNano()/1e6)
//时间戳（纳秒转换为秒） time.Now().UnixNano()/1e9)
func UnixTimeString(q int64) (p string) {
	t := time.Now()
	if q == 1 { //秒
		second := t.Unix()
		return other.Int64TurnString(second)
	} else if q == 2 { //纳秒
		nanosecond := t.UnixNano()
		return other.Int64TurnString(nanosecond)
	} else if q == 3 { //毫秒
		millisecond := t.UnixNano() / 1e6
		return other.Int64TurnString(millisecond)
	}
	return ""
}

// MonthlyTotal 输入年月 计算一个月的总数
func MonthlyTotal(year int, month int) (days int) {
	if month != 2 {
		if month == 4 || month == 6 || month == 9 || month == 11 {
			days = 30
		} else {
			days = 31
		}
	} else {
		if ((year%4) == 0 && (year%100) != 0) || (year%400) == 0 {
			days = 29
		} else {
			days = 28
		}
	}
	return
}

// StringToInt64 标准时间格式(秒)转换为int64时间戳
func StringToInt64(s string) (t int64) {
	datetime := "2015-01-01 00:00:00" //待转化为时间戳的字符串
	loc, _ := time.LoadLocation(s)
	tmp, _ := time.ParseInLocation(s, datetime, loc)
	timestamp := tmp.Unix()
	return timestamp
}
