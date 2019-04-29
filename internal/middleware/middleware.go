package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"solarium-data-collector/internal/config_reader"
	"strings"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		secret := config_reader.Config().Token
		authValue := strings.SplitN(r.Header.Get("Authorization"), " ", 2)

		if len(authValue) != 2 {

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			body := map[string]string{"error": "Authorization Error"}
			_ = json.NewEncoder(w).Encode(body)
			return

		}

		token, err := jwt.Parse(authValue[1], func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secret), nil

		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

			deviceInfo := map[string]string{"device_id": claims["device_id"].(string),
				"user_id": claims["user_id"].(string)}
			ctx := context.WithValue(r.Context(), "deviceInfo", deviceInfo)
			next.ServeHTTP(w, r.WithContext(ctx))

		} else {

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusForbidden)
			body := map[string]string{"error": err.Error()}
			_ = json.NewEncoder(w).Encode(body)
			return
		}
	})
}
