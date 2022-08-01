// 阿里云 短信发送
package aliyun

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

const (
	domain = "dysmsapi.aliyuncs.com"
)

func SendMessage(userInput *UserParams) (bool, string, error) {
	params := make(Params)
	params.Set("AccessKeyId", userInput.AccessKeyId)
	params.Set("Timestamp", time.Now().UTC().Format("2006-01-02T15:04:05Z")) // 格式为：yyyy-MM-dd’T’HH:mm:ss’Z’；时区为：GMT
	params.Set("SignatureMethod", "HMAC-SHA1")                               // 建议固定值：HMAC-SHA1
	params.Set("SignatureVersion", "1.0")                                    // 建议固定值：1.0
	params.Set("SignatureNonce", GetRandomString(12))                        // 用于请求的防重放攻击，每次请求唯一
	params.Set("Format", "JSON")
	params.Set("Action", "SendSms")                      // API的命名，固定值，如发送短信API的值为：SendSms
	params.Set("Version", "2017-05-25")                  // API的版本，固定值，如短信API的值为：2017-05-25
	params.Set("RegionId", "cn-hangzhou")                // API支持的RegionID，如短信API的值为：cn-hangzhou
	params.Set("PhoneNumbers", userInput.PhoneNumbers)   // 短信接收号码,支持以逗号分隔的形式进行批量调用，批量上限为1000个手机号码,批量调用相对于单条调用及时性稍有延迟,验证码类型的短信推荐使用单条调用的方式
	params.Set("SignName", userInput.SignName)           // 短信签名
	params.Set("TemplateParam", userInput.TemplateParam) // 短信模板ID
	// 短信模板变量替换JSON串,友情提示:如果JSON中需要带换行符,
	// 请参照标准的JSON协议对换行符的要求,比如短信内容中包含\r\n的情况在JSON中需要表示成\r\n,
	// 否则会导致JSON在服务端解析失败
	params.Set("TemplateCode", userInput.TemplateCode)
	// 构造待签名的请求串
	value := params.SortToJoin()
	// 生成签名 签名采用HmacSHA1算法 + Base64
	sign := Sign(userInput.AppSecret, "GET&"+SpecialUrlEncode("/")+"&"+SpecialUrlEncode(value))
	// 增加签名结果到请求参数中，发送请求。签名也要做特殊URL编码
	getParams := "/?Signature=" + SpecialUrlEncode(sign) + "&" + value
	// fmt.Println("http://" + domain + getParams)
	resp, err := http.Get("http://" + domain + getParams)
	if err != nil {
		// handle error
		return false, "请求错误", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		return false, "接收错误", err
	}
	var res SendSmsResponse
	// json 转换
	err = json.Unmarshal(body, &res)
	if err != nil {
		return false, "JSON转对象错误", err
	}
	return res.Code == "OK", res.Message, nil

}

//生成随机字符串
func GetRandomString(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func (t *Params) SortToJoin() string {
	var keyList []string
	for k := range *t {
		keyList = append(keyList, k)
	}
	sort.Strings(keyList)
	sortQueryStringTmp := ""
	for _, v := range keyList {
		sortQueryStringTmp += "&" + SpecialUrlEncode(v) + "=" + SpecialUrlEncode(t.Get(v))
	}
	// 字符串转切片截取字符串开头多余的&
	// rune切片类型返回的长度为物理长度 len([]rune(string)) 获取的是肉眼可见的长度
	result := []rune(sortQueryStringTmp)
	return string(result[1:])
}

func Sign(appSecret, paramsEncodeValue string) string {
	mac := hmac.New(sha1.New, []byte(appSecret+"&"))
	mac.Write([]byte(paramsEncodeValue))
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}

func SpecialUrlEncode(str string) string {
	return strings.Replace(strings.Replace(strings.Replace(url.QueryEscape(str), "+", "%20", -1), "*", "%2A", -1), "%7E", "~", -1)
}
