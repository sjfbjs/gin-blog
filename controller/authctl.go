package controller

import (
	"fmt"
	"gin-blog/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func loginAPI(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	fmt.Println(username, password)
	user := service.User.Login(username, password)
	fmt.Println("user: ", user)
	if nil == user {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"msg":     "用户名或密码错误",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"msg":     "登录成功",
		})
	}
}
