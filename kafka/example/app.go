package example

import (
	"github.com/IBM/sarama"
	"github.com/google/wire"
	"github.com/oldsaltedfish/logf"
	"github.com/oldsaltedfish/wireprovider/config"
)

var ProviderSet = wire.NewSet(NewApp)

type App struct {
	Logger        logf.Logger
	AsyncProducer sarama.AsyncProducer
	SyncProducer  sarama.SyncProducer
	Config        *config.Config
}

func NewApp(
	Logger logf.Logger,
	AsyncProducer sarama.AsyncProducer,
	SyncProducer sarama.SyncProducer,
	Config *config.Config,
) *App {
	return &App{
		Logger:        Logger,
		AsyncProducer: AsyncProducer,
		SyncProducer:  SyncProducer,
		Config:        Config,
	}
}
