package controller

import (
	"blogapi/api/config"
	"blogapi/api/helper"
	"blogapi/api/modal"
	"blogapi/request"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func UserInsert(c echo.Context) error {
	var user modal.User
	var rq request.UserInsert

	if helper.Validator(&c, &rq) != nil {
		return nil
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
	if helper.Validator(&c, &rq) != nil {
		return nil
	}
	db := config.Conn()
	db.Where("email = ?", rq.Email).Find(&user)
	checkpass := helper.CheckPasswordHash(rq.Password, user.Password)
	if !checkpass {
		return c.JSON(http.StatusBadRequest, "Kullanıcı Adı Veya Şifre Hatalı")
	}
	claims := &config.UserJwtCustom{
		ID: user.Model.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, _ := token.SignedString([]byte("mykey"))

	return c.JSON(http.StatusOK, helper.Response(map[string]interface{}{
		"token": t,
		"Email": user}, "Giriş Başarılı Bir Şekilde Tamamlandı"))

}

func UserInfo(c echo.Context) error {
	var user modal.User
	userid := helper.AuthID(c)
	db := config.Conn()
	db.First(&user, userid)
	return c.JSON(http.StatusOK, userid)
}

func UserList(c echo.Context) error {
	var user []modal.User
	db := config.Conn()
	db.Find(&user)
	return c.JSON(http.StatusOK, helper.Response(user, "Kayıtlı Kullanıcılar"))
}
