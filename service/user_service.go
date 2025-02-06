/**
 * @Author QG
 * @Date  2025/1/19 21:59
 * @description
**/

package service

import (
	"ginChat/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetUserList
// @Tags 用户列表
// @Success 200 {string} welcome
// @Router /user/list [get]
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

// CreateUser
// @Summary 新增用户
// @Tags 用户模块
// @param name query string false "用户名"
// @param password query string false "密码"
// @param repassword query string false "确认密码"
// @Success 200 {string} json{"code","message"}
// @Router /user/createUser [post]
func CreateUser(c *gin.Context) {
	user := models.UserBasic{}
	user.Name = c.Query("name")
	password := c.Query("password")
	repassword := c.Query("repassword")
	if password != repassword {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "两次密码不一致",
		})
		return
	} else {
		user.PassWord = password
	}
	models.CreateUser(user)
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "create user successful！",
	})
}

// DeleteUser
// @Summary 删除用户
// @Tags 用户模块
// @param id query string false "用户id"
// @Success 200 {string} json{"code","message"}
// @Router /user/deleteUser [delete]
func DeleteUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.Query("id"))
	user.ID = uint(id)
	models.DeleteUser(user)
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "delete user successful！",
	})
}
