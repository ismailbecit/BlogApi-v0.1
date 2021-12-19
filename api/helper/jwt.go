package helper

import (
	"blogapi/api/config"
	"blogapi/api/modal"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Claims(user modal.User) *config.UserJwtCustom {
	claims := &config.UserJwtCustom{
		modal.User{
			Name:     user.Name,
			Surname:  user.Surname,
			Email:    user.Email,
			Password: user.Password,
			Age:      user.Age,
			Admin:    user.Admin,
			Model: gorm.Model{
				ID:        user.ID,
				CreatedAt: user.CreatedAt,
				UpdatedAt: user.UpdatedAt,
				DeletedAt: user.DeletedAt,
			},
		},
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}
	return claims
}

func AuthInfo(c echo.Context) modal.User {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*config.UserJwtCustom)
	info := modal.User{
		Model: gorm.Model{
			ID:        claims.User.ID,
			CreatedAt: claims.User.CreatedAt,
			UpdatedAt: claims.User.UpdatedAt,
			DeletedAt: claims.User.DeletedAt,
		},
		Name:     claims.User.Name,
		Surname:  claims.User.Surname,
		Email:    claims.User.Email,
		Password: claims.User.Password,
		Age:      claims.User.Age,
		Admin:    claims.User.Admin,
	}
	return info

}
