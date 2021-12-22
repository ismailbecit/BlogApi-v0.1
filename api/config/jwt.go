package config

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4/middleware"
)

type UserJwtCustom struct {
	ID uint `json:"id"`
	jwt.StandardClaims
}

var JwtConfig = middleware.JWTConfig{
	Claims:     &UserJwtCustom{},
	SigningKey: []byte("mykey"),
}
