package injector

import (
	"context"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/kyawmyintthein/rzcfg"
	"github.com/kyawmyintthein/rzlog"
	"github.com/kyawmyintthein/twirp-poc/app/config"
)

func ProvideDynamdbCfgClient(globalCfg *config.GlobalCfg, logger rzlog.Logger, sess *session.Session) (rzcfg.Client, error) {
	ctx := context.Background()

	cfgClient, err := rzcfg.NewConfigDBClient(&globalCfg.ClientCfg, sess, rzcfg.WithLogger(logger))
	if err != nil {
		return cfgClient, err
	}

	err = cfgClient.Init(ctx, globalCfg)
	if err != nil {
		rzlog.Error(ctx, err, "failed to load configuration")
		return cfgClient, err
	}

	return cfgClient, nil
}
