package main

import (
	"go-nornia/models"
	"go-nornia/routers"
	_ "go-nornia/utils/conf"
	"go-nornia/utils/db/mysql"
	"go-nornia/utils/db/redis"
	"go-nornia/utils/log"
)

func init() {
	log.InitConfig()  // config
	mysql.InitMysql() // mysql
	redis.InitRedis() // redis
}

func main() {
	models.InitModels()

	r := routers.InitRouter()
	_ = r.Run(":5071")

}
