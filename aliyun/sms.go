package aliyun

import (
	"github.com/golang/glog"
)

// SendSms 发送短信，需要传入 AccessKeyId 和 AppSecret
func SendSms(tplCode, mobile, code, accessKeyId, appSecret, signName string) bool {
	var userParams UserParams
	userParams.PhoneNumbers = mobile
	userParams.TemplateCode = tplCode
	userParams.AccessKeyId = accessKeyId
	userParams.AppSecret = appSecret
	userParams.SignName = signName
	userParams.TemplateParam = "{\"code\": \"" + code + "\"}"
	ok, msg, err := SendMessage(&userParams)
	if !ok {
		// 根据业务进行错误处理
		glog.Errorln(msg, err)
	}
	return ok
}

// SendSmsRegCode 注册模板
func SendSmsRegCode(mobile, code, accessKeyId, appSecret, signName string) bool {
	return SendSms("SMS_143718474", mobile, code, accessKeyId, appSecret, signName)
}

// SendSmsResetCode 其他模板
func SendSmsResetCode(mobile, code, accessKeyId, appSecret, signName string) bool {
	return SendSms("SMS_155370082", mobile, code, accessKeyId, appSecret, signName)
}

// SendSmsLoginCode 其他模板
func SendSmsLoginCode(mobile, code, accessKeyId, appSecret, signName string) bool {
	return SendSms("SMS_155370082", mobile, code, accessKeyId, appSecret, signName)
}
