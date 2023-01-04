package controllers

import (
	"gateway/controllers/responses"
	"strings"

	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

type Token struct {
	Login string
	jwt.StandardClaims
}

func NewToken(login string) (*Token, time.Time) {
	expTime := time.Now().Add(300 * time.Minute)
	return &Token{
			Login: login,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expTime.Unix(),
			},
		},
		expTime
}

func (tk *Token) ToString() string {
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenStr, err := token.SignedString([]byte(nil))
	if err != nil {
		tokenStr = ""
	}

	return tokenStr
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

type authCtrl struct {
}

func InitAuth(r *mux.Router) {
	ctrl := &authCtrl{}
	r.HandleFunc("/oauth/token", ctrl.token).Methods("POST")
}

type TokenRequest struct {
	Scope        string `json:"scope"`
	GrantType    string `json:"grant_type"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Clientid     string `json:"clientId"`
	Clientsecret string `json:"clientSecret"`
}
type TokenResponse struct {
	AccessToken string `json:"access_token"`
}

func (ctrl *authCtrl) token(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		responses.BadRequest(w, "wrong form")
	}
	token, _ := NewToken(r.PostFormValue("username"))
	tokenStr := token.ToString()
	responses.JsonSuccess(w, &TokenResponse{tokenStr})
}
