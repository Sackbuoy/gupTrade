package main

import (
	"sync"

	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/Sackbuoy/gupTrade/internal"
)

func main() {
	var config internal.Configuration

	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugaredLogger := logger.Sugar()

	// unmarshal config with viper
	err := internal.BuildConfiguration(viper.New(), &config)
	if err != nil {
		sugaredLogger.Fatalf("Failed to build configuration %s\n", err.Error())
	}

	var wg sync.WaitGroup

	// set up traders from config
	traders, err := internal.SetupTraders(config.Traders, sugaredLogger)
	if err != nil {
		sugaredLogger.Fatalf("Failed to set up Traders: %s\n", err.Error())
	}

	for _, t := range traders {
		t.RunAll(&wg)
	}

	wg.Wait()
}
