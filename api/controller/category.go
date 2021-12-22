package controller

import (
	"blogapi/api/config"
	"blogapi/api/helper"
	"blogapi/api/modal"
	"blogapi/repository"
	"blogapi/request"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CategoryInsert(c echo.Context) error {
	var rq request.CategoryInsert
	if helper.Validator(&c, &rq) != nil {
		return nil
	}
	err := repository.Get().Category().New(&rq)
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
	var category modal.Category
	var rq request.CategoryDel
	if helper.Validator(&c, &rq) != nil {
		return nil
	}
	db := config.Conn()
	//  id e ait kategorileri sorgulaama
	result := db.Find(&category, rq.ID)
	if result.RowsAffected == 0 {
		return c.JSON(http.StatusBadRequest, helper.Response(nil, "Kayıtlı İd Bulunamadı"))
	}
	db.Where("id = ? ", rq.ID).Find(&category)

	db.Unscoped().Delete(&category)

	return c.JSON(http.StatusOK, helper.Response(nil, "Silme İşlemi Başarılı!"))
}
