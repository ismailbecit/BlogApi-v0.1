package controller

import (
	"blogapi/api/config"
	"blogapi/api/helper"
	"blogapi/api/modal"
	"blogapi/request"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CategoryInsert(c echo.Context) error {
	var rq request.CategoryInsert
	if helper.Validator(&c, &rq) != nil {
		return nil
	}

	db := config.Conn()
	db.Create(&modal.Category{
		Name: rq.Name,
	})
	return c.JSON(http.StatusOK, "Kategori Başarıyla Oluşturuldu")

}
func CategoryList(c echo.Context) error {
	var category []modal.Category
	db := config.Conn()
	result := db.Find(&category)
	if result.RowsAffected == 0 {
		return c.JSON(http.StatusOK, "Kategori Bulunamadı")
	}
	return c.JSON(http.StatusOK, category)
}
func CategoryDel(c echo.Context) error {
	var category modal.Category
	var rq request.CategoryDel
	if helper.Validator(&c, &rq) != nil {
		return nil
	}
	db := config.Conn()
	//  id e ait kategorileri sorgulaama
	result := db.Find(&category, rq.ID)
	if result.RowsAffected == 0 {
		return c.JSON(http.StatusBadRequest, "Kayıtlı İd Bulunamadı")
	}
	db.Where("id = ? ", rq.ID).Find(&category)

	db.Unscoped().Delete(&category)

	return c.JSON(http.StatusOK, "Silme İşlemi Başarılı!")
}
