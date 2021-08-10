package infrastructure

import (
	"context"

	"github.com/kyawmyintthein/rzmiddleware"
	"github.com/twitchtv/twirp"
)

func NewInterceptor() twirp.Interceptor {
	return func(next twirp.Method) twirp.Method {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			//rzmiddleware.RequestID(ctx)
			twirp.SetHTTPResponseHeader(ctx, "req_id", rzmiddleware.GetReqID(ctx))
			return next(ctx, req)
		}
	}
}
