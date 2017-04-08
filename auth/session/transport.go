package session

import (
	"context"
	stdhttp "net/http"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport/http"
	"github.com/gorilla/sessions"
)

const defaultKey = "session"

func ToHTTPContext(store sessions.Store, name string) http.RequestFunc {
	return func(ctx context.Context, r *stdhttp.Request) context.Context {
		s, _ := store.Get(r, name)
		session := &session{r: r, s: s}
		return context.WithValue(ctx, defaultKey, session)
	}
}

func FromHTTPContext(logger log.Logger) http.ServerResponseFunc {
	return func(ctx context.Context, w stdhttp.ResponseWriter) context.Context {
		session, ok := ctx.Value(defaultKey).(*session)
		if !ok {
			logger.Log("session", "after_request", "err", "session empty")
			return ctx
		}

		if session.written {
			if err := session.s.Save(session.r, w); err != nil {
				logger.Log("session", "after_request", "err", err)
				return ctx
			}
			session.written = false
		}

		return context.WithValue(ctx, "session", session)
	}
}
