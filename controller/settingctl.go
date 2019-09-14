package controller

import (
	"fmt"
	"gin-blog/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func setting(c *gin.Context) {
	title := service.Config.GetTitle()
	keyword := service.Config.GetKeyword()
	c.HTML(http.StatusOK, "setting.html", gin.H{
		"title":   title,
		"keyword": keyword,
	})
}

func updateSetting(c *gin.Context) {
	title := c.PostForm("title")
	keywords := c.PostForm("keywords")
	fmt.Println(title, keywords)
	if service.Config.UpdateTitle(title) && service.Config.UpdateKeyword(keywords) {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"msg":     "设置更新成功",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"msg":     "设置更新失败",
		})
	}
}
