/**
 * @Author: Cc
 * @Description: kafka
 * @File: kafka
 * @Version: 1.0.0
 * @Date: 2022/7/30 16:45
 * @Software : GoLand
 */

package kafkac

import (
	"github.com/Shopify/sarama"
	"log"
)

func InitStartProducer(address []string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	config.Version = sarama.V0_10_2_0
	//client, err := sarama.NewClient(address, config)
	//if err != nil {
	//	return nil, err
	//}

	//fromClient, err := sarama.NewAsyncProducerFromClient(client)
	//successes := fromClient.Successes()

	producer, err := sarama.NewSyncProducer(address, config)
	if err != nil {
		log.Fatalln("InitStartKafka Error", err.Error())
	}

	return producer, nil
}

// InitStartConsumer 传不同的groupID
func InitStartConsumer(address []string, groupID string) (sarama.ConsumerGroup, error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Version = sarama.V0_10_2_0
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	consumer, err := sarama.NewConsumerGroup(address, groupID, config)
	if err != nil {
		log.Fatalln("InitStartKafka Error", err.Error())
	}
	return consumer, nil
}
