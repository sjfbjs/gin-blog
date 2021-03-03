package controller

import (
	"fmt"
	"gin-blog/model"
	"gin-blog/service"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func uploadFile(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "msg": "文件上传失败"})
		return
	}
	//文件内容，后面再进行处理
	content, err := ioutil.ReadAll(file)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "msg": "文件读取失败"})
		return
	}

	fmt.Println(header.Filename)
	fmt.Println(string(content))
	upload := model.Upload{}
	upload.Filename = header.Filename
	saved := service.Upload.SaveFile(&upload)
	if saved {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"msg":     "文件上传成功",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"msg":     "文件上传失败",
		})
		return
	}
}

// gin DELETE 的 PostForm() 方法不能解析 content-type: application/x-www-form-urlencoded
func deleteFile(c *gin.Context) {
	fname := c.PostForm("filename")
	if service.Upload.DeleteByFilename(fname) {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"msg":     "文件删除成功",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"msg":     "文件删除失败",
		})
	}
}

func FileManage(c *gin.Context) {
	Uploads := service.Upload.GetAll()
	c.HTML(http.StatusOK, "uploadfile.html", gin.H{
		"Uploads": Uploads,
	})
}
