package config

import (
	"blogapi/api/modal"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4/middleware"
)

type UserJwtCustom struct {
	User modal.User `json:"user"`
	jwt.StandardClaims
}

var JwtConfig = middleware.JWTConfig{
	Claims:     &UserJwtCustom{},
	SigningKey: []byte("mykey"),
}
