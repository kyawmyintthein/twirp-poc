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
	ucv1 "github.com/kyawmyintthein/twirp-poc/app/domain/usecase/v1"
	"github.com/kyawmyintthein/twirp-poc/app/infrastructure"
	userpbs "github.com/kyawmyintthein/twirp-poc/protopbs/protos/user/v1"
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

		userRPC userpbs.UserServce
		userUC  ucv1.UserUC
	}
)

func (app *twripApp) Init() error {
	twirpHandler := userpbs.NewUserServceServer(app.userRPC,
		twirp.WithServerPathPrefix("/rz"),
		twirp.WithServerInterceptors(infrastructure.NewInterceptor()),
		twirp.WithServerHooks(rzlog.TwirpServerLoggingHook()))
	app.httpServer.Handler = twirpHandler
	return nil
}

func (app *twripApp) Start() {
	go func() {
		err := app.httpServer.ListenAndServeTLS("server.crt", "server.key")
		if err != nil {
			rzlog.Errorf(context.Background(), err, "Failed to start HTTP server on port :%d ", app.cfg.App.ServerPort)
			os.Exit(-1)
		}
	}()
	rzlog.Infof(context.Background(), "HTTP Server started on port :%d ", app.cfg.App.ServerPort)
}
