package aliyun

import (
	"strconv"
)

type UserParams struct {
	AccessKeyId     string
	AppSecret       string
	PhoneNumbers    string // 接收手机号
	SignName        string // 短信签名
	TemplateCode    string // 短信模板ID
	TemplateParam   string // 短信模板变量替换JSON串,友情提示:如果JSON中需要带换行符,请参照标准的JSON协议对换行符的要求,比如短信内容中包含\r\n的情况在JSON中需要表示成\r\n,否则会导致JSON在服务端解析失败
	SmsUpExtendCode string // 上行短信扩展码,无特殊需要此字段的用户请忽略此字段
	OutId           string // 外部流水扩展字段
}

type Params map[string]string

func (t Params) Get(key string) string {
	if value, ok := t[key]; ok {
		return value
	}
	return ""
}

func (t Params) Set(key, value string) {
	t[key] = value
}

func (t Params) SetInterface(key string, value interface{}) {
	if value == nil {
		return
	}

	switch value := value.(type) {
	case int8, int16, int32, int64:
		v, _ := value.(int64)
		t[key] = strconv.FormatInt(v, 10)

	case uint8, uint16, uint32, uint64:
		v, _ := value.(uint64)
		t[key] = strconv.FormatUint(v, 10)

	case float32, float64:
		v, _ := value.(float64)
		t[key] = strconv.FormatFloat(v, 'f', 0, 64)

	case bool:
		t[key] = strconv.FormatBool(value)

	case string:
		t[key] = value
	}

}

type SendSmsResponse struct {
	Message   string
	RequestId string
	BizId     string
	Code      string
}
