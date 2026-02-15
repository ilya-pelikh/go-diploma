package server

import (
	"fmt"
	"net/http"

	"diploma/internal/identity"
	"diploma/internal/pkg/config"
	"diploma/internal/pkg/logger"

	"go.uber.org/zap"
)

func auth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pass := config.ENV.TODO_PASSWORD

		if len(pass) > 0 {
			var jwt string

			cookie, err := r.Cookie("token")
			if err == nil {
				jwt = cookie.Value
			}

			valid, err := identity.VerifyToken(jwt)

			if err != nil {
				logger.Logger.Warn("Validation token: ", zap.Error(err))
			}

			fmt.Println(valid)

			if !valid {
				// возвращаем ошибку авторизации 401
				http.Error(w, "Authentification required", http.StatusUnauthorized)
				return
			}
		}
		next(w, r)
	})
}
