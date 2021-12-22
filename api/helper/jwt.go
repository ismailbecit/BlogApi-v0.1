package helper

import (
	"blogapi/api/config"
	"blogapi/api/modal"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func AuthID(c echo.Context) uint {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*config.UserJwtCustom)
	info := modal.User{
		Model: gorm.Model{
			ID: claims.ID,
		},
	}
	return info.Model.ID

}
