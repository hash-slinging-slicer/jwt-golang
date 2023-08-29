package middlewares

import (
	"jwt/config"
	"jwt/helper"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

func JWTMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				response := map[string]string{"message": "Unauthorized"}
				helper.ResponeJSON(w, http.StatusUnauthorized, response)
				return
			}
		}

		// Get Token User
		tokenString := c.Value

		claims := &config.JWTClaim{}

		// Parsing Token JWT
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return config.JWT_KEY, nil
		})

		// Invalid
		if err != nil {
			response := map[string]string{"message": "Unauthorized"}
			helper.ResponeJSON(w, http.StatusUnauthorized, response)
			return
		}

		if !token.Valid {
			response := map[string]string{"message": "Unauthorized"}
			helper.ResponeJSON(w, http.StatusUnauthorized, response)
			return
		}

		next.ServeHTTP(w, r)

	})
}
