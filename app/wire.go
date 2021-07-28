//+build wireinject

package app

import (
	"github.com/google/wire"
	"github.com/kyawmyintthein/rzlog"
	"github.com/kyawmyintthein/twirp-poc/app/devliery/rpcv1"
	ucv1 "github.com/kyawmyintthein/twirp-poc/app/domain/usecase/v1"
	"github.com/kyawmyintthein/twirp-poc/app/infrastructure"
	"github.com/kyawmyintthein/twirp-poc/app/injector"
)

func New(configFilePath string) (*twripApp, error) {
	panic(wire.Build(
		injector.ProvideGlobalConfig,
		injector.ProvideAWSSession,
		injector.ProvideLogConfig,
		injector.ProvideRedisConfig,
		rzlog.Init,
		injector.ProvideDynamodb,
		injector.ProvideDynamdbCfgClient,
		injector.ProvideHTTPServer,
		infrastructure.NewRedisClient,
		infrastructure.NewRedisCache,
		ucv1.ProvideUserUC,
		rpcv1.NewUserServiceRPCImpl,
		wire.Struct(new(twripApp), "*"),
	))
}
