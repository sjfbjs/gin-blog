package controller

import (
	"gin-blog/model/param"
	"gin-blog/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func editArticle(c *gin.Context) {
	id := c.Query("id")
	if "" == id {
		c.HTML(http.StatusOK, "edit.html", nil)
		return
	}
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
	form := param.ArticleParam{}
	if err := c.ShouldBind(&form); nil != err {
		log.Println("articleParam binding err: ", err.Error())
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"msg":     "请输入标题和正文",
		})
		return
	}
	saved := service.Article.SaveArticle(&form)
	if saved {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"msg":     "文章发布成功",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"msg":     "文章发布失败",
		})
	}
}

// gin DELETE 的 PostForm() 方法不能解析 content-type: application/x-www-form-urlencoded
func deleteArticle(c *gin.Context) {
	id := c.PostForm("id")

	idNum, err := strconv.Atoi(id)
	if nil != err {
		log.Println("id Atoi err: ", err.Error())
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"msg":     "删除失败",
		})
		return
	}
	if service.Article.DeleteArticle(idNum) {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"msg":     "删除成功",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"msg":     "删除失败",
		})
	}
}

func articleManage(c *gin.Context) {
	articles := service.Article.GetAll()
	c.HTML(http.StatusOK, "manage.html", gin.H{
		"articles": articles,
	})
}
