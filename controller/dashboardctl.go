package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func dashboard(c *gin.Context) {
	c.HTML(http.StatusOK, "dashboard.html", nil)
}