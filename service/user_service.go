/**
 * @Author QG
 * @Date  2025/1/19 21:59
 * @description
**/

package service

import (
	"fmt"
	"ginChat/global"
	"ginChat/models"
	"ginChat/utils"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"strconv"
	"time"
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
//func GetIndex(c *gin.Context) {
//	c.JSON(200, gin.H{
//		"message": "pong",
//	})
//}

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
	salt := fmt.Sprintf("%06d", rand.Int31())
	u := models.FindUserByName(user.Name)
	if u.Name != "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "用户名已存在",
		})
		return
	}

	if password != repassword {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "两次密码不一致",
		})
		return
	}
	makePassword := utils.MakePassword(password, salt)
	user.PassWord = makePassword
	user.Salt = salt
	user.LoginTime = time.Now()
	user.LoginOutTime = time.Now()
	user.HeartbeatTime = time.Now()
	models.CreateUser(user)
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "create user successful！",
	})
}

// UpdateUser
// @Summary 修改用户
// @Tags 用户模块
// @param id formData string false "id"
// @param name formData string false "name"
// @param password formData string false "password"
// @param phone formData string false "phone"
// @param email formData string false "email"
// @Success 200 {string} json{"code","message"}
// @Router /user/updateUser [post]
func UpdateUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.PostForm("id"))
	user.ID = uint(id)
	user.Name = c.PostForm("name")
	user.PassWord = c.PostForm("password")
	user.Phone = c.PostForm("phone")
	user.Avatar = c.PostForm("icon")
	user.Email = c.PostForm("email")
	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "字段不符合预期",
		})
	} else {
		models.UpdateUser(user)
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "update user successful！",
		})
	}
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

// FindUserByNameAndPwd
// @Summary 登录
// @Tags 用户模块
// @param name query string false "用户名"
// @param password query string false "密码"
// @Success 200 {string} json{"code","message"}
// @Router /user/findUserByNameAndPwd [post]
func FindUserByNameAndPwd(c *gin.Context) {
	name := c.Request.FormValue("name")
	password := c.Request.FormValue("password")
	u := models.FindUserByName(name)
	if u.Name == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "用户名或密码错误",
		})
	}
	flag := utils.ValidPassword(password, u.Salt, u.PassWord)
	if !flag {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "用户名或密码错误",
		})
	}
	password = utils.MakePassword(password, u.Salt)
	user := models.FindUserByNameAndPwd(name, password)
	//token加密
	str := fmt.Sprintf("%d", time.Now().Unix())
	temp := utils.MD5Encode(str)
	global.DB.Model(&user).Where("id = ?", user.ID).Update("identity", temp)
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "login successful！",
		"data":    user,
	})
}
