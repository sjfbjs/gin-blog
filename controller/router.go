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
	store := cookie.NewStore([]byte("d56b699830e77ba53855679cb1d252da"))
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
	router.GET("/logout", logout)

	admin := router.Group("/admin", func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(config.LOGIN_SESSION)
		if nil == user {
			c.String(http.StatusOK, "请先登录！")
			c.Abort()
		}
		c.Next()
	})
	{
		admin.GET("/", dashboard)
		admin.GET("/edit", editArticle)
		admin.POST("/article", addArticle)
		/**
		 * gin DELETE 的 PostForm() 方法不能解析 content-type: application/x-www-form-urlencoded
		 * 先使用 PUT 代替 DELETE
		 * https://github.com/gin-gonic/gin/issues/1755
		 */
		admin.PUT("/article", deleteArticle)
		admin.GET("/article", articleManage)

		admin.GET("/keyedit", editKey)
		admin.POST("/key", addKey)
		/**
		 * gin DELETE 的 PostForm() 方法不能解析 content-type: application/x-www-form-urlencoded
		 * 先使用 PUT 代替 DELETE
		 * https://github.com/gin-gonic/gin/issues/1755
		 */
		admin.PUT("/key", deleteKey)
		admin.GET("/key", keyManage)

		admin.GET("/setting", setting)
		admin.POST("/setting", updateSetting)
	}

	return router
}

func formatAsDate(timestamp int64) string {
	t := time.Unix(timestamp, 0)
	year, month, day := t.Date()
	return fmt.Sprintf("%d/%02d/%02d", year, month, day)
}
