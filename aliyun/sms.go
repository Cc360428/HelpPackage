package aliyun

import (
	"github.com/golang/glog"
)

func SendSms(tplCode, mobile, code string) bool {
	var userParams UserParams
	userParams.PhoneNumbers = mobile
	userParams.TemplateCode = tplCode
	userParams.AccessKeyId = ""
	userParams.AppSecret = ""
	userParams.SignName = "智慧乐园"
	userParams.TemplateParam = "{\"code\": \"" + code + "\"}"
	ok, msg, err := SendMessage(&userParams)
	if !ok {
		// 根据业务进行错误处理
		glog.Errorln(msg, err)
	}
	return ok
}

//注册模板
func SendSmsRegCode(mobile, code string) bool {
	return SendSms("SMS_143718474", mobile, code)
}

//其他模板
func SendSmsResetCode(mobile, code string) bool {
	return SendSms("SMS_155370082", mobile, code)
}

//其他模板
func SendSmsLoginCode(mobile, code string) bool {
	return SendSms("SMS_155370082", mobile, code)
}
