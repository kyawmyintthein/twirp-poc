package main

import (
	"context"
	"log"
	"net/http"

	haberdasherpb "github.com/kyawmyintthein/twirp-poc/protopbs/protos/haberdasher"
	"github.com/kyawmyintthein/twirp-poc/rpc/haberdasher"
	"github.com/twitchtv/twirp"
)

func main() {
	server := &haberdasher.Server{} // implements Haberdasher interface
	twirpHandler := haberdasherpb.NewHaberdasherServer(server, twirp.WithServerPathPrefix("rz"), twirp.WithServerInterceptors(NewInterceptorMakeSmallHats()),
		twirp.WithServerHooks(NewLoggingServerHooks()))
	http.ListenAndServe(":8080", twirpHandler)
}

func NewInterceptorMakeSmallHats() twirp.Interceptor {
	return func(next twirp.Method) twirp.Method {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			val, ok := twirp.MethodName(ctx)
			if val == "MakeHat" && ok {
				return next(ctx, &haberdasherpb.Size{Inches: 1})
			}
			return next(ctx, req)
		}
	}
}

func NewLoggingServerHooks() *twirp.ServerHooks {
	return &twirp.ServerHooks{
		RequestRouted: func(ctx context.Context) (context.Context, error) {
			method, _ := twirp.MethodName(ctx)
			log.Println("Method: " + method)
			return ctx, nil
		},
		Error: func(ctx context.Context, twerr twirp.Error) context.Context {
			log.Println("Error: " + string(twerr.Code()))
			return ctx
		},
		ResponseSent: func(ctx context.Context) {
			log.Println("Response Sent (error or success)")
		},
	}
}
