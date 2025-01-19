/**
 * @Author QG
 * @Date  2025/1/19 18:06
 * @description
**/

package test

import (
	"ginChat/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestUser(t *testing.T) {
	dsn := "root:dyj3101631524@tcp(localhost:3306)/gin-chat?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	db.AutoMigrate(&models.UserBasic{})

	// Create
	db.Create(&models.UserBasic{Name: "张三"})

	// Read
	var user models.UserBasic
	db.First(&user, 1) // 根据整型主键查找

	// Update - 将 product 的 price 更新为 200
	db.Model(&user).Update("PassWord", 200)
	// Update - 更新多个字段

	// Delete - 删除 product
	db.Delete(&user, 1)
}
