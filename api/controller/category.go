package controller

import (
	"blogapi/api/config"
	"blogapi/api/modal"
	"blogapi/request"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CategoryInsert(c echo.Context) error {
	var rq request.CategoryInsert
	if err := c.Bind(&rq); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	db := config.Conn()
	db.Create(&modal.Category{
		Name: rq.Name,
	})
	return c.JSON(http.StatusOK, "Kategori Başarıyla Oluşturuldu")

}
