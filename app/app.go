package app

import (
	"context"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/go-redis/redis"
	"github.com/kyawmyintthein/rzcfg"
	"github.com/kyawmyintthein/rzlog"
	"github.com/kyawmyintthein/twirp-poc/app/config"
	"github.com/kyawmyintthein/twirp-poc/app/domain/repository"
	"github.com/kyawmyintthein/twirp-poc/app/infrastructure"
	objectpbs "github.com/kyawmyintthein/twirp-poc/protopbs/protos/objects"
	"github.com/twitchtv/twirp"
)

type (
	App interface {
		Init() error
	}

	twripApp struct {
		cfg    *config.GlobalCfg
		cfgCli rzcfg.Client
		logger rzlog.Logger

		httpServer  *http.Server
		redisClient redis.Cmdable
		cache       infrastructure.Cache

		session  *session.Session
		dynamodb repository.Dynamodb

		objectRPC objectpbs.ObjectService
	}
)

func (app *twripApp) Init() error {

	twirpHandler := objectpbs.NewObjectServiceServer(app.objectRPC,
		twirp.WithServerPathPrefix("/rz"),
		twirp.WithServerInterceptors(infrastructure.NewInterceptor()),
		twirp.WithServerHooks(infrastructure.NewLoggingServerHooks()))
	app.httpServer.Handler = twirpHandler
	return nil
}

func (app *twripApp) Start() {
	go func() {
		err := app.httpServer.ListenAndServe()
		if err != nil {
			rzlog.Errorf(context.Background(), err, "Failed to start HTTP server on port :%d ", app.cfg.App.ServerPort)
			os.Exit(-1)
		}
	}()
	rzlog.Infof(context.Background(), "HTTP Server started on port :%d ", app.cfg.App.ServerPort)
}
