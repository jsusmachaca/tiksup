package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/jsusmachaca/tiksup/api/response"
	"github.com/jsusmachaca/tiksup/internal/util"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if !strings.HasPrefix(token, "Bearer ") {
			response.WriteJsonError(w, "Token not provided", http.StatusUnauthorized)
			return
		}
		token = token[7:]
		claims, err := util.ValidateToken(token)
		if err != nil {
			response.WriteJsonError(w, "Token is not valid", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), TokenClaims, claims)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
