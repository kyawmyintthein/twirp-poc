package infrastructure

import (
	"context"

	"github.com/twitchtv/twirp"
)

func NewInterceptor() twirp.Interceptor {
	return func(next twirp.Method) twirp.Method {
		return func(ctx context.Context, req interface{}) (interface{}, error) {

			return next(ctx, req)
		}
	}
}
