package kviewer

import (
	"fmt"

	"github.com/IBM/sarama"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Erro ao ler o arquivo de configuração: %s", err))
	}
}

func GetKafkaProducer() (sarama.SyncProducer, error) {
	kafkaHost := viper.GetString("kafka.host")
	kafkaPort := viper.GetString("kafka.port")
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	return sarama.NewSyncProducer([]string{kafkaHost + ":" + kafkaPort}, config)
}

func GetKafkaConsumer() (sarama.Consumer, error) {
	kafkaHost := viper.GetString("kafka.host")
	kafkaPort := viper.GetString("kafka.port")
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	return sarama.NewConsumer([]string{kafkaHost + ":" + kafkaPort}, config)
}
