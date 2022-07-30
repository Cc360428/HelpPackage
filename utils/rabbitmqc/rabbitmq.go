/**
 * @Author: Cc
 * @Description: 消息队列
 * @File: rabbitmq
 * @Version: 1.0.0
 * @Date: 2022/7/30 10:22
 * @Software : GoLand
 */

package rabbitmqc

import (
	"github.com/streadway/amqp"
)

func initStart(url string) (*amqp.Connection, error) {
	dial, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}
	return dial, nil
}
