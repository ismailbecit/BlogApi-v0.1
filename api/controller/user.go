package controller

import (
	"blogapi/api/config"
	"blogapi/api/helper"
	"blogapi/api/modal"
	"blogapi/repository"
	"blogapi/request"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func UserInsert(c echo.Context) error {
	var rq request.UserInsert

	if helper.Validator(&c, &rq) != nil {
		return nil
	}
	hashpass, _ := helper.HashPassword(rq.Password)

	// kayıt işlemi
	user := modal.User{
		Name:     rq.Name,
		Surname:  rq.Surname,
		Email:    rq.Email,
		Password: hashpass,
		Age:      rq.Age,
	}

	result := repository.Get().User().EmailQuery(rq.Email)
	if result != 0 {
		return c.JSON(http.StatusBadRequest, "Kullanıcı Zaten Kayıtlı")
	}

	claims := &config.UserJwtCustom{
		ID: user.Model.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, _ := token.SignedString([]byte("mykey"))

	newuser := repository.Get().User().New(user)

	return c.JSON(http.StatusOK, helper.Response(map[string]interface{}{
		"userinfo": newuser, "token": t},
		"Tebrikler Başarıyla Kayıt Oldunuz"))

}

func UserLogin(c echo.Context) error {
	var user modal.User
	var rq request.UserLogin
	if helper.Validator(&c, &rq) != nil {
		return nil
	}
	result := repository.Get().User().EmailQuery(rq.Email)
	if result == 1 {
		checkpass := helper.CheckPasswordHash(rq.Password, user.Password)
		if !checkpass {
			return c.JSON(http.StatusBadRequest, "Kullanıcı Adı Veya Şifre Hatalı")
		}
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
	userid := helper.AuthID(c)
	user := repository.Get().User().UserInfo(userid)
	return c.JSON(http.StatusOK, helper.Response(user, "Kullanıcı Bilgisi"))
}

func UserList(c echo.Context) error {
	var user []modal.User
	users := repository.Get().User().UserList(user)
	return c.JSON(http.StatusOK, helper.Response(users, "Kayıtlı Kullanıcılar"))
}

func UserDel(c echo.Context) error {
	var rq request.UserDel
	if helper.Validator(&c, &rq) != nil {
		return nil
	}
	repository.Get().User().UserDel(rq.Id)
	return c.JSON(http.StatusOK, helper.Response(nil, "Kullanıcı Silindi"))
}
