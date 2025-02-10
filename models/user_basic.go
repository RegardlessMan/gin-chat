/**
 * @Author QG
 * @Date  2025/1/19 17:58
 * @description
**/

package models

import (
	"fmt"
	"ginChat/global"
	"gorm.io/gorm"
	"time"
)

type UserBasic struct {
	gorm.Model
	Name          string
	PassWord      string
	Phone         string `valid:"matches(^1[3-9]{1}\\d{9}$)"`
	Email         string `valid:"email"`
	Avatar        string //头像
	Identity      string
	ClientIp      string
	ClientPort    string
	Salt          string
	LoginTime     time.Time
	HeartbeatTime time.Time
	LoginOutTime  time.Time `gorm:"column:login_out_time" json:"login_out_time"`
	IsLogout      bool
	DeviceInfo    string
}

func GetUserList() []UserBasic {
	var userList []UserBasic
	global.DB.Debug().Find(&userList)
	for _, v := range userList {
		fmt.Println(v)
	}
	return userList
}

func CreateUser(user UserBasic) *gorm.DB {
	return global.DB.Create(&user)
}

func FindUserByName(name string) UserBasic {
	user := UserBasic{}
	global.DB.Debug().Where("name = ?", name).First(&user)
	return user
}

func UpdateUser(user UserBasic) *gorm.DB {
	return global.DB.Model(&user).Updates(UserBasic{Name: user.Name, PassWord: user.PassWord, Phone: user.Phone, Email: user.Email, Avatar: user.Avatar})
}

func DeleteUser(user UserBasic) *gorm.DB {
	return global.DB.Delete(&user)
}

func FindUserByNameAndPwd(name string, password string) UserBasic {
	user := UserBasic{}
	global.DB.Where("name = ? and pass_word=?", name, password).First(&user)
	return user
}
