package controller

import (
	"fmt"
	"gin-blog/config"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"path/filepath"
	"time"
)

func MapRoutes() *gin.Engine {
	router := gin.Default()
	store := cookie.NewStore([]byte("login"))
	router.Use(sessions.Sessions("my_session", store))

	router.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})
	templates, _ := filepath.Glob("template/*.html")
	adminTemplates, _ := filepath.Glob("template/admin/*.html")
	templates = append(templates, adminTemplates...)
	router.LoadHTMLFiles(templates...)
	router.Static("/static", "template/static")

	router.GET("/", index)
	router.GET("/post/:slug", post)
	router.GET("/login", login)
	router.POST("/login", loginAPI)

	admin := router.Group("/admin", func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(config.LOGIN_SESSION)
		if nil == user {
			c.String(http.StatusOK, "请先登录！")
			c.Abort()
		}
	})
	{
		admin.GET("/", dashboard)
	}

	return router
}

func formatAsDate(timestamp uint) string {
	t := time.Unix(int64(timestamp), 0)
	year, month, day := t.Date()
	return fmt.Sprintf("%d/%02d/%02d", year, month, day)
}
