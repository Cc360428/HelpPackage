package utils

import "time"

//获取当前的日
func DateDay() (day int) {
	now := time.Now()
	day = now.Day()
	return day
}

// 获取今天周
func GetWeek() (week int64) {
	t := time.Now()
	i := int(t.Weekday())
	week, _ = ToInt64(i)
	return week
}

//获取当日零时时间戳（秒）
//Time stamp of the morning of the day: unit seconds
func DateDayZero() (r int64) {
	timeStr := time.Now().Format("2006-01-02")
	//使用Parse 默认获取为UTC时区 需要获取本地时区 所以使用ParseInLocation
	t2, _ := time.ParseInLocation("2006-01-02", timeStr, time.Local)
	r = t2.AddDate(0, 0, 0).Unix()
	return r
}

//Time stamp for the morning of the day: in milliseconds
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

//返回值int64类型
//时间戳（秒）： time.Now().Unix())
//时间戳（纳秒）： time.Now().UnixNano())
//时间戳（毫秒）： time.Now().UnixNano()/1e6)
//时间戳（纳秒转换为秒） time.Now().UnixNano()/1e9)
//
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

//返回值string类型
//时间戳（秒）： time.Now().Unix())
//时间戳（纳秒）： time.Now().UnixNano())
//时间戳（毫秒）： time.Now().UnixNano()/1e6)
//时间戳（纳秒转换为秒） time.Now().UnixNano()/1e9)
//
func UnixTimeString(q int64) (p string) {
	t := time.Now()
	if q == 1 { //秒
		second := t.Unix()
		return Int64TurnString(second)
	} else if q == 2 { //纳秒
		nanosecond := t.UnixNano()
		return Int64TurnString(nanosecond)
	} else if q == 3 { //毫秒
		millisecond := t.UnixNano() / 1e6
		return Int64TurnString(millisecond)
	}
	return ""
}

//
//  输入年月 计算一个月的总数
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

// 标准时间格式(秒)转换为int64时间戳
func StringToInt64(s string) (t int64) {
	datetime := "2015-01-01 00:00:00" //待转化为时间戳的字符串
	loc, _ := time.LoadLocation(s)
	tmp, _ := time.ParseInLocation(s, datetime, loc)
	timestamp := tmp.Unix()
	return timestamp
}
