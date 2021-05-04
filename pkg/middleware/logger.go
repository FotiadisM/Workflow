package middleware

import (
	"context"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport/http"
)

func Logger(l log.Logger) endpoint.Middleware {
	return func(in endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {

			method := ctx.Value(http.ContextKeyRequestMethod)
			if method != nil {
				l = log.With(l, "method", method)
			}

			path := ctx.Value(http.ContextKeyRequestPath)
			if path != nil {
				l = log.With(l, "path", path)
			}

			defer func(begin time.Time) {
				_ = l.Log(
					"took", time.Since(begin),
					"err", err,
				)
			}(time.Now())

			return in(ctx, request)
		}
	}
}
