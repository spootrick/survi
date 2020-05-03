package middleware

import (
	"context"
	"fmt"
	"github.com/spootrick/survi/api/auth"
	"github.com/spootrick/survi/api/model"
	"github.com/spootrick/survi/api/response"
	"github.com/spootrick/survi/api/util/customtype"
	"log"
	"net/http"
)

func SetMiddlewareLogger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s%s %s", r.Method, r.Host, r.RequestURI, r.Proto)
		next(w, r)
	}
}

func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

func SetMiddleWareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := auth.ExtractJWT(w, r)
		if token == nil {
			return
		}
		if token.Valid {
			ctx := context.WithValue(r.Context(), customtype.UserKey("user"), token.Claims.(*model.JWTClaims).User)
			next(w, r.WithContext(ctx))
		} else {
			response.ERROR(w, http.StatusUnauthorized, fmt.Errorf("invalid token"))
		}
	}
}
