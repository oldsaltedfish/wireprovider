package kafak

import (
	"github.com/IBM/sarama"
	"github.com/google/wire"
	"github.com/oldsaltedfish/logf"
	"github.com/oldsaltedfish/wireprovider/config"
)

var (
	ProviderSet = wire.NewSet(NewAsyncProducer, NewSyncProducer)
)

func NewAsyncProducer(config *config.Config, log logf.Logger) (sarama.AsyncProducer, error) {
	kafkaConf := sarama.NewConfig()
	err := config.UnmarshalKey("kafka.properties", kafkaConf)
	if err != nil {
		log.Errorf("kafka config Unmarshal kafka.properties Key error: %v", err)
		return nil, err
	}
	brokers := []string{}
	err = config.UnmarshalKey("kafka.brokers", &brokers)
	if err != nil {
		log.Errorf("kafka config Unmarshal kafka.brokers error: %v", err)
		return nil, err
	}
	p, err := sarama.NewAsyncProducer(brokers, kafkaConf)
	if err != nil {
		log.Errorf("kafka producer init error: %v", err)
	}
	return p, err
}

func NewSyncProducer(config *config.Config, log logf.Logger) (sarama.SyncProducer, error) {
	kafkaConf := sarama.NewConfig()
	err := config.UnmarshalKey("kafka.properties", kafkaConf)
	if err != nil {
		log.Errorf("kafka config Unmarshal kafka.properties Key error: %v", err)
		return nil, err
	}
	brokers := []string{}
	err = config.UnmarshalKey("kafka.brokers", &brokers)
	if err != nil {
		log.Errorf("kafka config Unmarshal kafka.brokers error: %v", err)
		return nil, err
	}
	p, err := sarama.NewSyncProducer(brokers, kafkaConf)
	if err != nil {
		log.Errorf("kafka producer init error: %v", err)
	}
	return p, err
}
