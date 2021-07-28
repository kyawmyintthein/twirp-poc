package injector

import (
	"context"

	"github.com/kyawmyintthein/rzcfg"
	"github.com/kyawmyintthein/rzlog"
	"github.com/kyawmyintthein/twirp-poc/app/config"
	"github.com/kyawmyintthein/twirp-poc/app/infrastructure"
)

func ProvideGlobalConfig(configFilePath string) (*config.GlobalCfg, error) {
	var globalCfg config.GlobalCfg
	err := rzcfg.LoadConfig(configFilePath, "TPOC", &globalCfg)
	if err != nil {
		return &globalCfg, err
	}
	rzlog.DebugKV(context.Background(), rzlog.KV{"config": globalCfg}, "configuration is loaded")
	return &globalCfg, nil
}

func ProvideLogConfig(globalCfg *config.GlobalCfg) rzlog.LogCfg {
	return globalCfg.Log
}

func ProvideRedisConfig(globalCfg *config.GlobalCfg) *infrastructure.RedisCfg {
	return &globalCfg.Redis
}
