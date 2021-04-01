// string
package utils

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

//判断是否为空
// 输入："" 				输出：false
// 输入：长度等于0 		输出：false
// 输入："1" 			输出：true
func IsNull(str string) (bool bool) {
	if str == "" && len(str) == 0 {
		return false
	}
	return true
}

//RandomString 在数字、大写字母、小写字母范围内生成num位的随机字符串
func RandomString(length int) string {
	// 48 ~ 57 数字
	// 65 ~ 90 A ~ Z
	// 97 ~ 122 a ~ z
	// 一共62个字符，在0~61进行随机，小于10时，在数字范围随机，
	// 小于36在大写范围内随机，其他在小写范围随机
	rand.Seed(time.Now().UnixNano())
	result := make([]string, 0, length)
	for i := 0; i < length; i++ {
		t := rand.Intn(62)
		if t < 10 {
			result = append(result, strconv.Itoa(rand.Intn(10)))
		} else if t < 36 {
			result = append(result, string(rune(rand.Intn(26)+65)))
		} else {
			result = append(result, string(rune(rand.Intn(26)+97)))
		}
	}
	return strings.Join(result, "")
}

//start：正数 - 在字符串的指定位置开始,超出字符串长度强制把start变为字符串长度
//       负数 - 在从字符串结尾的指定位置开始
//       0 - 在字符串中的第一个字符处开始
//length:正数 - 从 start 参数所在的位置返回
//       负数 - 从字符串末端返回
func Substr(str string, start, length int) string {
	if length == 0 {
		return ""
	}
	runeStr := []rune(str)
	lenStr := len(runeStr)
	if start < 0 {
		start = lenStr + start
	}
	if start > lenStr {
		start = lenStr
	}
	end := start + length
	if end > lenStr {
		end = lenStr
	}
	if length < 0 {
		end = lenStr + length
	}
	if start > end {
		start, end = end, start
	}
	return string(runeStr[start:end])
}

//随机取出一定长度的随机数字
func GenValidateCode(width int) string {
	str := "0123456789"
	b := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < width; i++ {
		result = append(result, b[r.Intn(len(b))])
	}
	return string(result)
}

//加密
func Md5V3(m string) string {
	w := md5.New()
	_, _ = io.WriteString(w, m)
	md5str := fmt.Sprintf("%x", w.Sum(nil))
	return md5str
}

//密码和盐加密存入数据库
func Encryption(password string, salt string) string {
	buf := bytes.NewBufferString(password)
	buf.Write([]byte(salt))
	return Md5V3(buf.String())
}

//数组去重
func RemoveDuplicatesAndEmpty(a []string) (ret []string) {
	aLen := len(a)
	for i := 0; i < aLen; i++ {
		if (i > 1 && a[i-1] == a[i]) || len(a[i]) == 0 {
			continue
		}
		ret = append(ret, a[i])
	}
	return
}

func Intercept(parameter string, condition string) string {
	c := strings.Index(parameter, condition)
	pos := strings.Index(parameter[c:], "")
	s := parameter[c+pos:]
	return s
}

//字符串分割
//第一个参数为要分割的字符串，第二个参数为分割的条件
//返回分割后的数组
func SplitUtil(parameter string, condition string) []string {
	str := strings.Split(parameter, condition)
	return str
}
