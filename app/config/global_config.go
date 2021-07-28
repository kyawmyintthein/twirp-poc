package config

import (
	"github.com/kyawmyintthein/rzcfg"
	"github.com/kyawmyintthein/rzlog"
	"github.com/kyawmyintthein/twirp-poc/app/infrastructure"
)

type (
	GlobalCfg struct {
		App             AppCfg                  `envconfig:"APP"`
		Log             rzlog.LogCfg            `envconfig:"LOG"`
		HeathCheckCache HealthCheckCacheCfg     `envconfig:"HEALTH_CHECK_CACHE" json:"health_check_cache" mapstructure:"health_check_cache" required:"true"`
		ClientCfg       rzcfg.DynamodbDBCfg     `envconfig:"CLIENT_CFG" json:"client_cfg" required:"true"`
		Redis           infrastructure.RedisCfg `envconfig:"REDIS" josn:"redis"`
		AWS             AWSCfg                  `envconfig:"AWS"`
	}

	AWSCfg struct {
		Region  string `envconfig:"REGION" required:"true" default:""`
		Profile string `envconfig:"PROFILE" required:"false" default:""`
	}

	HealthCheckCacheCfg struct {
		Enable      bool  `mapstructure:"enable" json:"enable" default:"true"`
		TTLInSecond int64 `mapstructure:"ttl_in_seconds" json:"ttl_in_seconds" default:"180"`
	}

	AppCfg struct {
		ServerPort int `envconfig:"HTTP_SERVER_PORT" required:"true" default:"8080"`
	}
)
