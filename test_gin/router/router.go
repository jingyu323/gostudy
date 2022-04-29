package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var Router *gin.Engine

func GinInit() {
	// 禁用控制台颜色
	//gin.DisableConsoleColor()

	//gin.New()返回一个*Engine 指针
	//而gin.Default()不但返回一个*Engine 指针，而且还进行了debugPrintWARNINGDefault()和engine.Use(Logger(), Recovery())其他的一些中间件操作
	Router = gin.Default()
	//Router = gin.New()
}

func SetupRouter(projectPath string) {

	//使用日志
	//Router.Use(gin.Logger())
	//使用Panic处理方案
	//Router.Use(gin.Recovery())

	Router.Use(InitErrorHandler)
	Router.Use(InitAccessLogMiddleware)

	// 未知调用方式
	Router.NoMethod(InitNoMethodJson)
	// 未知路由处理
	Router.NoRoute(InitNoRouteJson)

	// Ping
	Router.GET(projectPath+"/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"ping": "pong2222",
		})
	})

	Router.POST(projectPath+"/pp", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"ping": "post",
		})
	})

	Router.GET("/someXML", func(c *gin.Context) {
		// 会输出头格式为 text/xml; charset=UTF-8 的 xml 字符串
		c.XML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	Router.GET(projectPath+"/someYAML", func(c *gin.Context) {
		// 会输出头格式为 text/yaml; charset=UTF-8 的 yaml 字符串
		c.YAML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

}
