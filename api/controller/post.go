package controller

import (
	"blogapi/api/config"
	"blogapi/api/modal"
	"net/http"

	"github.com/labstack/echo/v4"
)

func PostList(c echo.Context) error {
	var user []modal.User
	db := config.Conn()

	result := db.Find(&user)
	if result.RowsAffected == 0 {
		return c.JSON(http.StatusBadRequest, "Kayıtlı Post Bulunamadı")
	}
	return c.JSON(http.StatusOK, user)

}
