package controller

//
//import (
//	"fmt"
//	"gin-blog/model/param"
//	"gin-blog/service"
//	"github.com/gin-gonic/gin"
//	"log"
//	"net/http"
//)
//
//func editKey(c *gin.Context) {
//	ip := c.Query("ip")
//	if "" == ip {
//		c.HTML(http.StatusOK, "editkey.html", nil)
//		return
//	}
//
//	Key := service.Key.GetByIp(ip)
//	c.HTML(http.StatusOK, "editkey.html", gin.H{
//		"Ip":  Key.Ip,
//		"Key": Key.Key,
//	})
//
//}
//
//func addKey(c *gin.Context) {
//	form := param.KeyParam{}
//	if err := c.ShouldBind(&form); nil != err {
//		log.Println("KeyParam binding err: ", err.Error())
//		c.JSON(http.StatusOK, gin.H{
//			"success": false,
//			"msg":     "请输入ip和私钥内容",
//		})
//		return
//	}
//	fmt.Println(form)
//	saved := service.Key.SaveKey(&form)
//	if saved {
//		c.JSON(http.StatusOK, gin.H{
//			"success": true,
//			"msg":     "私钥新增成功",
//		})
//	} else {
//		c.JSON(http.StatusOK, gin.H{
//			"success": false,
//			"msg":     "私钥新增失败",
//		})
//	}
//}
//
//// gin DELETE 的 PostForm() 方法不能解析 content-type: application/x-www-form-urlencoded
//func deleteKey(c *gin.Context) {
//	ip := c.PostForm("ip")
//	if service.Key.DeleteKey(ip) {
//		c.JSON(http.StatusOK, gin.H{
//			"success": true,
//			"msg":     "私钥删除成功",
//		})
//	} else {
//		c.JSON(http.StatusOK, gin.H{
//			"success": false,
//			"msg":     "私钥删除失败",
//		})
//	}
//}
//
//func keyManage(c *gin.Context) {
//	Keys := service.Key.GetAll()
//	c.HTML(http.StatusOK, "managekey.html", gin.H{
//		"Keys": Keys,
//	})
//}
