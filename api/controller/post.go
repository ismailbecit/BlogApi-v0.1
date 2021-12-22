package controller

import (
	"blogapi/api/config"
	"blogapi/api/helper"
	"blogapi/api/modal"
	"blogapi/request"
	"net/http"

	"github.com/labstack/echo/v4"
)

func PostList(c echo.Context) error {
	var post []modal.Post
	db := config.Conn()

	result := db.Find(&post)

	db.Preload("Category").Preload("User").Find(&post)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusBadRequest, "Kayıtlı Post Bulunamadı")
	}
	return c.JSON(http.StatusOK, post)

}
func PostInsert(c echo.Context) error {
	var rq request.PostInsert
	var category modal.Category
	if helper.Validator(&c, &rq) != nil {
		return nil
	}
	db := config.Conn()

	result := db.Where("id = ? ", rq.Categoryfk).Find(&category)
	if result.RowsAffected == 0 {
		return c.JSON(http.StatusBadRequest, helper.Response(nil, "Kategori İd Bulunamadı"))
	}
	authinfo := helper.AuthInfo(c)

	db.Create(&modal.Post{
		Title:      rq.Title,
		Content:    rq.Content,
		Categoryfk: rq.Categoryfk,
		Userfk:     authinfo.ID,
	})

	return c.JSON(http.StatusOK, helper.Response(nil, "Post Kaydedildi"))
}

func PostDel(c echo.Context) error {
	var post modal.Post
	var rq request.PostDel
	if helper.Validator(&c, &rq) != nil {
		return nil
	}
	db := config.Conn()
	result := db.Find(&post, rq.ID)
	if result.RowsAffected == 0 {
		return c.JSON(http.StatusBadRequest, helper.Response(nil, "Kayıtlı id bulunamadı"))
	}
	db.Delete(&post, rq.ID)
	return c.JSON(http.StatusOK, helper.Response(nil, "Post Silindi"))
}
