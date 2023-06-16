package trader

import (
	"fmt"
	"sync"

	"go.uber.org/zap"

	"github.com/Sackbuoy/gupTrade/internal/strategies"
	"github.com/Sackbuoy/gupTrade/pkg/brokers"
	"github.com/Sackbuoy/gupTrade/pkg/brokers/tradier"
)

type Trader struct {
	Strategies []strategies.Strategy
	logger     *zap.SugaredLogger
}

func NewTrader(config TraderConfiguration, logger *zap.SugaredLogger) (*Trader, error) {
	trader := Trader{
		logger: logger,
	}
	var client brokers.BrokerClient

	switch config.Name {
	case "tradier":
		client = tradier.NewClient(config.Broker.Auth, &config.Broker.Url)
	default:
		return nil, fmt.Errorf("unsupported trader %s", config.Name)
	}

	trader.initializeStrategies(config.Strategies, client, logger)

	return &trader, nil
}

func (t *Trader) initializeStrategies(configs []strategies.StrategyConfiguration, client brokers.BrokerClient, logger *zap.SugaredLogger) {
	traderStrategies := make([]strategies.Strategy, 0)

	for _, config := range configs {
		strat, err := strategies.NewStrategy(config, client, logger)
		if err != nil {
			logger.Errorf("could not initialize strategy %s\n", err.Error())
		}
		traderStrategies = append(traderStrategies, strat)
	}

	t.Strategies = traderStrategies
}

func (t *Trader) RunAll(wg *sync.WaitGroup) error {
	for _, s := range t.Strategies {
		wg.Add(1)
		go s.Run(wg)
	}

	return nil
}
