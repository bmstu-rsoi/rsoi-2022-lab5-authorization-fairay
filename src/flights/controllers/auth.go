package controllers

import (
	"flights/controllers/responses"
	"strings"

	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type Token struct {
	Login string
	jwt.StandardClaims
}


func RetrieveToken(w http.ResponseWriter, r *http.Request) (*Token) {
	reqToken := r.Header.Get("Authorization")
	if len(reqToken) == 0 {
		responses.TokenIsMissing(w)
		return nil
	}
	splitToken := strings.Split(reqToken, "Bearer ")
	tokenStr := splitToken[1]
	tk := &Token{}

	token, err := jwt.ParseWithClaims(tokenStr, tk, func(token *jwt.Token) (interface{}, error) {
		return []byte(nil), nil
	})

	if err != nil || !token.Valid {
		responses.JwtAccessDenied(w)
		return nil
	}

	if time.Now().Unix()-tk.ExpiresAt > 0 {
		responses.TokenExpired(w)
		return nil
	}

	return tk
}

var JwtAuthentication = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if token := RetrieveToken(w, r); token != nil {
			r.Header.Set("X-User-Name", token.Login)
			next.ServeHTTP(w, r)
		}
	})
}
