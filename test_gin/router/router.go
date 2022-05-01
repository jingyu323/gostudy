package router

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"test_gin/dao"
	"test_gin/mysql_connection"
)

var Router *gin.Engine
var orderTest *dao.OrderTest
func GinInit() {
	orderTest =new(dao.OrderTest)
	// 禁用控制台颜色
	//gin.DisableConsoleColor()

	//gin.New()返回一个*Engine 指针
	//而gin.Default()不但返回一个*Engine 指针，而且还进行了debugPrintWARNINGDefault()和engine.Use(Logger(), Recovery())其他的一些中间件操作
	Router = gin.Default()
	//Router = gin.New()

	//Router.LoadHTMLGlob("./views/**/*")
	//Router.LoadHTMLGlob("views/**/*")
	Router.LoadHTMLGlob("templates/**/*")
	//Router.LoadHTMLGlob("templates/*")

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
	Router.GET(projectPath+"/datafromdb", func(c *gin.Context) {
		// 会输出头格式为 text/yaml; charset=UTF-8 的 yaml 字符串
		orders := mysql_connection.GetAllOrders()
		fmt.Print("orders:is")
		fmt.Println(len(orders))

		for key, v := range orders {
			fmt.Println("key is %s", key)
			fmt.Println(v)
		}

		data, err := json.Marshal(orders)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("ttt      %s\n", data)
		c.JSON(http.StatusOK, gin.H{"message": data, "status": http.StatusOK})
	})

	Router.POST("/form", func(context *gin.Context) {
		types := context.DefaultPostForm("type", "post")
		username := context.PostForm("username")
		password := context.PostForm("password")

		context.String(http.StatusOK, fmt.Sprintf("username:%s , password:%s , types:%s", username, password, types))
	})

	Router.GET(projectPath+"/userlogin", func(c *gin.Context) {

		c.HTML(http.StatusOK, "user_login22.html", gin.H{
			"title": "Users",
		})

	})
	Router.GET(projectPath+"/getLastOrder", func(c *gin.Context) {

		orders := orderTest.Query()
		fmt.Print("orders:is")
		fmt.Println(len(orders))

		for key, v := range orders {
			fmt.Println("key is %s", key)
			fmt.Println(v)
		}

		data, err := json.Marshal(orders)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("ttt      %s\n", data)
		c.JSON(http.StatusOK, gin.H{"message": data, "status": http.StatusOK})
	})

}
