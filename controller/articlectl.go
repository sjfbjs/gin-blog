package controller

import (
	"gin-blog/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func editArticle(c *gin.Context) {
	id := c.Query("id")
	idNum, err := strconv.Atoi(id)
	if nil != err {
		log.Println("id Atoi err: ", err.Error())
		c.HTML(http.StatusOK, "edit.html", nil)
	} else {
		article := service.Article.GetBySlugOrId(id)
		c.HTML(http.StatusOK, "edit.html", gin.H{
			"id":      idNum,
			"title":   article.Title,
			"alias":   article.Alias,
			"content": article.Content,
		})
	}
}

func addArticle(c *gin.Context) {

}

func deleteArticle(c *gin.Context) {

}

func articleManage(c *gin.Context) {

}
