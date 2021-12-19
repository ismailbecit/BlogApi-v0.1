package controller

import (
	"blogapi/api/config"
	"blogapi/api/helper"
	"blogapi/api/modal"
	"blogapi/request"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func UserInsert(c echo.Context) error {
	var user modal.User
	var rq request.UserInsert

	if err := c.Bind(&rq); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	db := config.Conn()
	result := db.Where("email = ? ", rq.Email).Find(&user)
	if result.RowsAffected != 0 {
		return c.JSON(http.StatusBadRequest, "Kullanıcı Zaten Kayıtlı")
	}

	hashpass, _ := helper.HashPassword(rq.Password)

	// kayıt işlemi
	db.Create(&modal.User{
		Name:     rq.Name,
		Surname:  rq.Surname,
		Email:    rq.Email,
		Password: hashpass,
		Age:      rq.Age,
	})
	return c.JSON(http.StatusOK, "Tebrikler Başarıyla Kayıt Oldunuz")

}

func UserLogin(c echo.Context) error {
	var user modal.User
	var rq request.UserLogin

	if err := c.Bind(&rq); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	db := config.Conn()
	db.Where("email = ?", rq.Email).Find(&user)
	checkpass := helper.CheckPasswordHash(rq.Password, user.Password)
	if !checkpass {
		return c.JSON(http.StatusBadRequest, "Kullanıcı Adı Veya Şifre Hatalı")
	}
	claims := helper.Claims(user)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, _ := token.SignedString([]byte("mykey"))

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
		"info":  user,
	})

}

func UserInfo(c echo.Context) error {
	user := helper.AuthInfo(c)
	return c.JSON(http.StatusOK, user.ID)
}
