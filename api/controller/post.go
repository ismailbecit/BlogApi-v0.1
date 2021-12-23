package controller

import (
	"blogapi/api/helper"
	"blogapi/api/modal"
	"blogapi/repository"
	"blogapi/request"
	"net/http"

	"github.com/labstack/echo/v4"
)

func PostList(c echo.Context) error {
	postlist := repository.Get().Post().List()
	return c.JSON(http.StatusOK, postlist)
}
func PostInsert(c echo.Context) error {
	var rq request.PostInsert
	if helper.Validator(&c, &rq) != nil {
		return nil
	}
	result := repository.Get().Post().CategoryFK(rq.Categoryfk)
	if result == 0 {
		return c.JSON(http.StatusBadRequest, helper.Response(nil, "Kategori İd Bulunamadı"))
	}
	authinfo := helper.AuthID(c)

	post := modal.Post{
		Title:      rq.Title,
		Content:    rq.Content,
		Categoryfk: rq.Categoryfk,
		Userfk:     authinfo,
	}
	repository.Get().Post().New(post)

	return c.JSON(http.StatusOK, helper.Response(nil, "Post Kaydedildi"))
}

func PostDel(c echo.Context) error {
	var post modal.Post
	var rq request.PostDel
	if helper.Validator(&c, &rq) != nil {
		return nil
	}
	result := repository.Get().Post().Query(post, uint(rq.ID))
	if result == 0 {
		return c.JSON(http.StatusBadRequest, helper.Response(nil, "Kayıtlı id bulunamadı"))
	}
	repository.Get().Post().Del(uint(rq.ID))
	return c.JSON(http.StatusOK, helper.Response(nil, "Post Silindi"))
}
