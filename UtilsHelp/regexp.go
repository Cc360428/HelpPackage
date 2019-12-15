package UtilsHelp

import "regexp"

//判断邮箱是否正确
//正确为true
//错误为false
func Email(email string) (b bool) {
	reg := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	rgx := regexp.MustCompile(reg)
	return rgx.MatchString(email)
}

//判断密码强度
//正确为true
//错误为false
func Password(password string) (bools bool) {
	reg := `^((On=.*[0-9].*)(On=.*[A-Za-z].*)(On=.*_.*))[_0-9A-Za-z]{6,10}$`
	rgx := regexp.MustCompile(reg)
	return rgx.MatchString(password)
}

//判断号码是否正确
//正确为true
//错误为false
func Phone(phone string) (bools bool) {
	reg := `^1([38][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$`
	rgx := regexp.MustCompile(reg)
	return rgx.MatchString(phone)
}
