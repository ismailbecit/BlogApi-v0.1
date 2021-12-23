package controller

import (
	"blogapi/api/config"
	"blogapi/api/helper"
	"blogapi/api/modal"
	"blogapi/repository"
	"blogapi/request"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CategoryInsert(c echo.Context) error {
	var rq request.CategoryInsert
	if helper.Validator(&c, &rq) != nil {
		return nil
	}
	category := modal.Category{
		Name: rq.Name,
	}
	err := repository.Get().Category().New(category)
	return c.JSON(http.StatusOK, helper.Response(err, "Kayıt Başarılı"))

}
func CategoryList(c echo.Context) error {
	var category []modal.Category
	db := config.Conn()
	result := db.Find(&category)
	if result.RowsAffected == 0 {
		return c.JSON(http.StatusBadRequest, helper.Response(nil, "Kategori Bulunamadı"))
	}
	return c.JSON(http.StatusOK, category)
}
func CategoryDel(c echo.Context) error {
	var rq request.CategoryDel

	if helper.Validator(&c, &rq) != nil {
		return nil
	}
	category := modal.Category{
		Model: gorm.Model{
			ID: rq.ID,
		},
	}
	fmt.Println(category.ID)

	repository.Get().Category().Del(category)

	return c.JSON(http.StatusOK, helper.Response(nil, "Silme İşlemi Başarılı!"))
}
