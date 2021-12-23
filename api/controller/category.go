package controller

import (
	"blogapi/api/helper"
	"blogapi/api/modal"
	"blogapi/repository"
	"blogapi/request"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
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
	categories := repository.Get().Category().List()
	return c.JSON(http.StatusOK, helper.Response(categories, "Kategori Listesi"))
}
func CategoryDel(c echo.Context) error {
	var rq request.CategoryDel
	if helper.Validator(&c, &rq) != nil {
		return nil
	}

	result := repository.Get().Category().Query(rq.ID)
	fmt.Println(result)
	if result == 0 {
		return c.JSON(http.StatusBadRequest, helper.Response(nil, "Böyle Bir Id Bulunamadı"))
	}

	repository.Get().Category().Del(rq.ID)
	return c.JSON(http.StatusOK, helper.Response(nil, "Silme İşlemi Başarılı!"))
}
