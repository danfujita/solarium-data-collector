package middlware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"solarium-golang/internal/configs"
	"strings"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		secret := configs.Config().Token

		authValue := strings.SplitN(r.Header.Get("Authorization"), " ", 2)

		if len(authValue) != 2 {
			http.Error(w, "Authorization Error", 401)
			return
		}
		token, err := jwt.Parse(authValue[1], func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(secret), nil
		})
		if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			next.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte(err.Error()))
			return
		}

	})
}
