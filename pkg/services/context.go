package services

import (
	"context"
	"lenslocked/pkg/auth"
	"lenslocked/pkg/datamodel"
	"log"
	"net/http"
)

type CtxKey string

const (
	CtxKeyUser CtxKey = "user"
)

func SetUserCtx(ctx context.Context, user *datamodel.User) context.Context {
	return context.WithValue(ctx, CtxKeyUser, user)
}
func GetUserCtx(ctx context.Context) *datamodel.User {
	if user, ok := ctx.Value(CtxKeyUser).(*datamodel.User); ok {
		return user
	}
	return nil
}
func DeleteUserCtx(ctx context.Context) context.Context {
	return context.WithValue(ctx, CtxKeyUser, nil)
}

type UserMiddleware struct {
	Session *SessionService
}

func (u *UserMiddleware) SetUser(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := r.Cookie(auth.CookieTypeSession)
		if err != nil {
			handler.ServeHTTP(w, r)
			return
		}

		user, err := u.Session.GetUserForSession(token.Value)
		if err != nil {
			log.Default().Printf("trouble finding user for session %s %s", token.Value, err.Error())
			handler.ServeHTTP(w, r)
			return
		}

		r = r.WithContext(SetUserCtx(r.Context(), user))
		handler.ServeHTTP(w, r)
	})
}

func (u *UserMiddleware) RequireUser(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := GetUserCtx(r.Context())
		if user == nil {
			log.Default().Println("could not find user in context. redirecting to signin.")
			http.Redirect(w, r, "/signin", http.StatusFound)
			return
		}
		handler.ServeHTTP(w, r)
	})
}
