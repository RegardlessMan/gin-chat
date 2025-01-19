/**
 * @Author QG
 * @Date  2025/1/19 21:50
 * @description
**/

package router

import (
	"ginChat/service"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	gin.SetMode(gin.DebugMode)
	r.GET("/index", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// User-related api
	r.GET("/user/list", service.GetUserList)

	return r
}
