package controller

import (
	"fmt"
	"gin-blog/config"
	"gin-blog/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func login(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(config.LOGIN_SESSION)
	if user == nil {
		c.HTML(http.StatusOK, "login.html", nil)
	} else {
		c.HTML(http.StatusOK, "dashboard.html", nil)
	}

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
		fmt.Println("user: ", user)
		session.Set(config.LOGIN_SESSION, &user)
		err := session.Save()
		if nil != err {
			log.Println("session save err: ", err)
		}
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"msg":     "登录成功",
		})
	}
}

func logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.String(http.StatusOK, "退出登录成功!")
}
