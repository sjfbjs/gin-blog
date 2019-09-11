package controller

import (
	"gin-blog/config"
	"gin-blog/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func loginAPI(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	user := service.User.Login(username, password)
	if nil == user {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"msg":     "用户名或密码错误",
		})
	} else {
		session := sessions.Default(c)
		session.Set(config.LOGIN_SESSION, &user)
		session.Save()
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"msg":     "登录成功",
		})
	}
}
