package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-nornia/controllers"
	"go-nornia/middleware"
	"time"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	// r.Use(middleware.Cors(),middleware.ZapRecovery(true))  // 线上关闭路由日志
	r.Use(middleware.Cors(), middleware.ZapLogger(), middleware.ZapRecovery(true))

	// 测试路由
	r.GET("/login", func(c *gin.Context) {
		c.String(200, "pong "+fmt.Sprint(time.Now().Unix()))
	})

	app := r.Group("/app")
	app.GET("/login", controllers.NewLoginController().IsNewUser)

	// 使用中间件检验
	//app.Use(middleware.Cors())
	//app.GET("/login", appControllers.NewLoginController().IsNewUser)

	return r
}
