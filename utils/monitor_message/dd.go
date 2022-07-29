// Package monitor_message
// @Description: 钉钉群预警
package monitor_message

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var (
	baseUrl = "https://oapi.dingtalk.com/robot/send?access_token="
	ddToken = "c2a5a11f2d15761fbc6a4222ad77a9778d547361afdc4cf8918e4f4c706870ec"
	url     = fmt.Sprintf("%s%s", baseUrl, ddToken)
)

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
	var messageAll NailRobot
	messageAll.Msgtype = "text"
	messageAll.Text.Content = fmt.Sprintf("%s %s", messageType, message)
	if isAtAll {
		messageAll.At.IsAtAll = true
	} else {
		messageAll.At.AtMobiles = []string{"18270681615"}
	}
	marshal, err := json.Marshal(messageAll)
	if err != nil {
		fmt.Println("钉钉预警错误", err.Error())
		return err
	}
	resp, err := http.Post(url,
		"application/json",
		strings.NewReader(string(marshal)),
	)
	if err != nil {
		fmt.Println("Cc360428 钉钉预警错误", err.Error())
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println(string(body))

	return nil
}
