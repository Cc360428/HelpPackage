/**
 * @Author: Cc
 * @Description: SendTelegram message
 * @File: telegram
 * @Version: 1.0.0
 * @Date: 2022/7/27 16:05
 * @Software : GoLand
 */

package utils

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// SendTelegram ...
// @Description: SendTelegram 发送小飞机预警
// @param chatID
// @param token
// @param text context
func SendTelegram(chatID int64, token, text string) {

	bot, err := tgbotapi.NewBotAPI(token)

	if err != nil {
		fmt.Println("NewBotApi Error", err.Error())
		return
	}

	newMessage := tgbotapi.NewMessage(chatID, text)
	_, err = bot.Send(newMessage)
	if err != nil {
		fmt.Println("SendTelegram Error", err.Error())
		return
	}
}
