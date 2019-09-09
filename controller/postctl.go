package controller

import (
	"gin-blog/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func post(c *gin.Context) {
	slug := c.Param("slug")
	title := service.Config.GetTitle()
	keyword := service.Config.GetKeyword()
	article := service.Article.GetBySlugOrId(slug)
	c.HTML(http.StatusOK, "post.html", gin.H{
		"title":   title,
		"keyword": keyword,
		"article": article,
	})
}
