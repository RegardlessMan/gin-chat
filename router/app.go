/**
 * @Author QG
 * @Date  2025/1/19 21:50
 * @description
**/

package router

import (
	"ginChat/docs"
	"ginChat/service"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	// 集成swagger
	docs.SwaggerInfo.BasePath = ""
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	gin.SetMode(gin.DebugMode)
	r.GET("/index", service.GetIndex)
	// User-related api
	r.GET("/user/list", service.GetUserList)
	r.POST("/user/createUser", service.CreateUser)
	r.DELETE("/user/deleteUser", service.DeleteUser)

	return r
}
