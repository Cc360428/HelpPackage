/**
 * @Author: Cc
 * @Description: 格式化输出
 * @File: format
 * @Version: 1.0.0
 * @Date: 2022/7/27 16:22
 * @Software : GoLand
 */

package timec

import "time"

// 当前时间

// TimeFormat ...
// @Description: 格式返回2021-01-06
// @return string
func TimeFormat() string {
	return time.Now().Format("2006-01-02")
}

// TimeFormat1 ...
// @Description: 2021-01-06 11:22:31
// @return string
func TimeFormat1() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// TimeFormat2 ...
// @Description: 格式返回2021/01/06
// @return string
func TimeFormat2() string {
	return time.Now().Format("2006/01/02")
}

// TimeFormat3 ...
// @Description: 2021/01/06 11:22:31
// @return string
func TimeFormat3() string {
	return time.Now().Format("2006/01/02 15:04:05")
}

// TimeFormat4 ...
// @Description: 20210106
// @return string
func TimeFormat4() string {
	return time.Now().Format("20060102")
}

// TimeFormat5 ...
// @Description: 20210106 11:22:31
// @return string
func TimeFormat5() string {
	return time.Now().Format("20060102 15:04:05")
}

// GetTodayZero ...
// @Description: 获取今天凌晨
// @return string
func GetTodayZero() string {
	t := time.Now()
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).Format("20060102")
}

// GetYesterdayZero ...
// @Description: 获取昨天凌晨
// @return string
func GetYesterdayZero() string {
	t := time.Now().AddDate(0, 0, -1)
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).Format("20060102")
}
