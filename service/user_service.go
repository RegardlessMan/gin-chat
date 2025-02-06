/**
 * @Author QG
 * @Date  2025/1/19 21:59
 * @description
**/

package service

import (
	"ginChat/models"
	"github.com/gin-gonic/gin"
)

func GetUserList(c *gin.Context) {
	data := models.GetUserList()
	c.JSON(200, gin.H{
		"code":    0, //  0成功   -1失败
		"message": "successful！",
		"data":    data,
	})
}

// GetIndex
// @Tags 首页
// @Success 200 {string} welcome
// @Router /index [get]
func GetIndex(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
