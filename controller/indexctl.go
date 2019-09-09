package controller

import (
	"gin-blog/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func index(c *gin.Context) {
	title := service.Config.GetTitle()
	keyword := service.Config.GetKeyword()
	articles := service.Article.GetAll()
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":    title,
		"keyword":  keyword,
		"articles": articles,
	})
}