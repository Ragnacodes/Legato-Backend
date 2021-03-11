package authenticate

import (
	"legato_server/domain"
	"legato_server/models"
	"context"
	"net/http"
)

var userCtxKey = &contextKey{"user"}
var requestCtxKey = &contextKey{"request"}
var responseWriterCtxKey = &contextKey{"response_writer"}

type contextKey struct {
	name string
}

func GqlgenAuthMiddleware(u domain.UserUseCase) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), responseWriterCtxKey, &w)
			ctx = context.WithValue(ctx, requestCtxKey, &r)

			// Grab Cookie
			c, err := r.Cookie("Authorization")
			if err != nil {
				r = r.WithContext(ctx)
				next.ServeHTTP(w, r)
				return
			}
			tokenStr := c.Value

			// Allow unauthenticated users in
			if tokenStr == "" {
				r = r.WithContext(ctx)
				next.ServeHTTP(w, r)
				return
			}

			//validate jwt token
			claim, err := CheckToken(tokenStr)
			if err != nil {
				http.Error(w, err.Error(), http.StatusForbidden)
				r = r.WithContext(ctx)
				next.ServeHTTP(w, r)
				return
			}

			// create user and check if user exists in db
			user, dbErr := u.GetUserByUsername(claim.Username)
			if dbErr != nil {
				r = r.WithContext(ctx)
				next.ServeHTTP(w, r)
				return
			}

			// put it in context
			ctx = context.WithValue(ctx, userCtxKey, &user)

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func UserForContext(ctx context.Context) *models.User {
	raw, _ := ctx.Value(userCtxKey).(*models.User)
	return raw
}

func RequestForContext(ctx context.Context) *http.ResponseWriter {
	raw, _ := ctx.Value(requestCtxKey).(*http.ResponseWriter)
	return raw
}

func ResponseWriterForContext(ctx context.Context) *http.ResponseWriter {
	raw, _ := ctx.Value(responseWriterCtxKey).(*http.ResponseWriter)
	return raw
}
