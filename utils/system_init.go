/**
 * @Author QG
 * @Date  2025/1/19 21:53
 * @description
**/

package utils

import (
	"fmt"
	"ginChat/global"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// InitConfig Read config file to Viper
func InitConfig() {
	viper.AddConfigPath("config")
	viper.SetConfigName("app")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println("app config is already initialized")
}

func InitMysql() {
	global.DB, _ = gorm.Open(mysql.Open(viper.GetString("mysql.dsn")),
		&gorm.Config{})
	fmt.Println(" MySQL initialized 。。。。")
	fmt.Println("mysql config is already initialized")
}
