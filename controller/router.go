package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"path/filepath"
	"time"
)

func MapRoutes() *gin.Engine {
	router := gin.Default()

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

	return router
}

func formatAsDate(timestamp uint) string {
	t := time.Unix(int64(timestamp), 0)
	year, month, day := t.Date()
	return fmt.Sprintf("%d/%02d/%02d", year, month, day)
}