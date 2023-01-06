module gateway

go 1.17

require github.com/jinzhu/gorm v1.9.2

require github.com/golang-jwt/jwt/v4 v4.4.2 // indirect

require (
	github.com/MicahParks/keyfunc v1.9.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gorilla/mux v1.8.0
	github.com/lib/pq v1.1.1 // indirect
	github.com/rs/cors v1.8.2
)
