package kviewer

import (
	"encoding/json"
	"fmt"

	"github.com/IBM/sarama"
	"github.com/spf13/viper"
)

func Consume(topic string, maxMsgs int) error {
	bootstrapServers := viper.GetString("bootstrap.servers")

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	brokers := []string{bootstrapServers}
	client, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		return fmt.Errorf("failed to create consumer: %s", err)
	}
	defer client.Close()

	partitions, err := client.Partitions(topic)
	if err != nil {
		return fmt.Errorf("failed to get partitions: %s", err)
	}

	msgCount := 0
	unlimited := (maxMsgs == 0) // Se maxMsgs é 0, então consuma indefinidamente

	for _, partition := range partitions {
		consumer, err := client.ConsumePartition(topic, partition, sarama.OffsetOldest)
		if err != nil {
			return fmt.Errorf("failed to start consumer for partition %d: %s", partition, err)
		}

		go func(consumer sarama.PartitionConsumer) {
			for msg := range consumer.Messages() {
				headersMap := make(map[string]string)
				for _, header := range msg.Headers {
					headersMap[string(header.Key)] = string(header.Value)
				}

				messageMap := map[string]interface{}{
					"header": headersMap,
					"body":   string(msg.Value),
				}

				jsonOutput, err := json.Marshal(messageMap)
				if err != nil {
					fmt.Printf("Error encoding JSON: %s\n", err)
					continue
				}

				fmt.Printf("Received message from partition %d at offset %d: %s\n", msg.Partition, msg.Offset, string(jsonOutput))

				if !unlimited {
					msgCount++
					if msgCount >= maxMsgs {
						consumer.Close()
						break
					}
				}
			}
		}(consumer)
	}

	// Mantenha a função em execução para continuar consumindo.
	select {}
}
