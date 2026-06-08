// Package monitor_message
// @Description: 钉钉群预警
package monitor_message

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
)

const baseUrl = "https://oapi.dingtalk.com/robot/send?access_token="

type DingDingClient struct {
	url string
	mu  sync.Mutex
}

var defaultClient *DingDingClient

func NewDingDing(token string) *DingDingClient {
	return &DingDingClient{
		url: fmt.Sprintf("%s%s", baseUrl, token),
	}
}

func InitDefaultDingDing(token string) {
	defaultClient = NewDingDing(token)
}

type NailRobot struct {
	Msgtype string `json:"msgtype"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
	At struct {
		AtMobiles []string `json:"atMobiles"`
		IsAtAll   bool     `json:"isAtAll"`
	} `json:"at"`
}

// Send Send("Error", " test", true)
func Send(messageType, message string, isAtAll bool) error {
	if defaultClient == nil {
		return fmt.Errorf("dingding client not initialized, call InitDefaultDingDing first")
	}
	return defaultClient.Send(messageType, message, isAtAll)
}

// Send 使用客户端发送消息
func (c *DingDingClient) Send(messageType, message string, isAtAll bool) error {
	var messageAll NailRobot
	messageAll.Msgtype = "text"
	messageAll.Text.Content = fmt.Sprintf("%s %s", messageType, message)
	if isAtAll {
		messageAll.At.IsAtAll = true
	}
	//else {
	//messageAll.At.AtMobiles = []string{"18270681615"}
	//}
	marshal, err := json.Marshal(messageAll)
	if err != nil {
		fmt.Println("钉钉预警错误", err.Error())
		return err
	}
	resp, err := http.Post(c.url,
		"application/json",
		strings.NewReader(string(marshal)),
	)
	if err != nil {
		fmt.Println("Cc360428 钉钉预警错误", err.Error())
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println(string(body))

	return nil
}
