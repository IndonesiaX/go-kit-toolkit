package session

import (
	"context"
	stdhttp "net/http"

	"github.com/go-kit/kit/transport/http"
	"github.com/gorilla/sessions"
)

func ToHTTPContext(store sessions.Store, name string) http.RequestFunc {
	return func(ctx context.Context, r *stdhttp.Request) context.Context {
		s, _ := store.Get(r, name)
		session := &session{r: r, s: s}
		return context.WithValue(ctx, "session", session)
	}
}

func FromHTTPContext() http.ServerResponseFunc {
	return func(ctx context.Context, w stdhttp.ResponseWriter) context.Context {
		session := ctx.Value("session").(session)
		if session.written {
			session.s.Save(session.r, w)
			session.written = false
		}
		return context.WithValue(ctx, "session", session)
	}
}
