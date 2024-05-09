package middleware

import (
	"context"
	"eniqilo_store/config"
	"eniqilo_store/helper"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/julienschmidt/httprouter"
)

func VerifyToken(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		tokenString := strings.Split(r.Header.Get("Authorization"), " ")[1]
		if tokenString == "" {
			helper.Unauthorized(w)
		} else {
			//get token value
			claims := &config.JWTClaim{}
			//parsing token
			key := []byte(os.Getenv("JWT_SECRET"))
			token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
				return key, nil
			})
			if err != nil {
				helper.Unauthorized(w)
			} else if !token.Valid {
				helper.Unauthorized(w)
			} else {
				ctx := context.WithValue(r.Context(), "staff_id", claims.StaffId)
				next(w, r.WithContext(ctx), p)
			}
		}
	}
}
