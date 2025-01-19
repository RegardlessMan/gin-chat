/**
 * @Author QG
 * @Date  2025/1/19 17:57
 * @description
**/

package main

import (
	"ginChat/router"
	"ginChat/utils"
)

func main() {
	// Init Config
	utils.InitConfig()
	utils.InitMysql()

	// run http server
	engine := router.Router()
	engine.Run(":8080")
}
