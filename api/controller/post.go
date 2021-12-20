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

	if err := c.Bind(&rq); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	authinfo := helper.AuthInfo(c)
	db := config.Conn()
	db.Create(&modal.Post{
		Title:      rq.Title,
		Content:    rq.Content,
		Categoryfk: rq.Categoryfk,
		Userfk:     authinfo.ID,
	})
	return c.JSON(http.StatusOK, "Post Kaydedildi")
}

func PostDel(c echo.Context) error {
	var post modal.Post
	var rq request.PostDel
	if err := c.Bind(&rq); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	db := config.Conn()
	result := db.Find(&post, rq.ID)
	if result.RowsAffected == 0 {
		return c.JSON(http.StatusOK, "Kayıtlı id bulunamadı")
	}
	db.Where("id = ? ", rq.ID).Find(&post)

	return c.JSON(http.StatusOK, "Post Silindi")
}
