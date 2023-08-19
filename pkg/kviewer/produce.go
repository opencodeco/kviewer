package kviewer

import (
	"github.com/IBM/sarama"
)

func Produce(topic string, message string) error {
	producer, err := GetKafkaProducer()
	if err != nil {
		return err
	}
	defer producer.Close()

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	_, _, err = producer.SendMessage(msg)
	return err
}
