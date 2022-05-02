package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	rsp "test_gin/common/response"
	"test_gin/dao"
	"test_gin/service"
)

func GetUserList(c *gin.Context) {
	todoList := service.Query()

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": todoList,
	})

}
func DeleteOrder(c *gin.Context) {

	orderId, ok := c.GetQuery("orderId")
	if !ok {
		rsp.Error(c, "无效的id")
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": orderId,
	})

	if err := service.Delete(orderId); err != nil {
		rsp.Error(c, err.Error())
	} else {
		rsp.Success(c, "删除成功", orderId)
	}

}

func CreateOrder(c *gin.Context) {

	param := c.Param("param")
	fmt.Println(param)

	query := c.Query("query")
	fmt.Println(query)

	defaultQuery := c.DefaultQuery("defaultQuery", "no")
	fmt.Println(defaultQuery)

	getQuery, res := c.GetQuery("getQuery")
	fmt.Println(getQuery, res)

	queryArray := c.QueryArray("queryArray")
	fmt.Println(queryArray)

	getQueryArray, res := c.GetQueryArray("getQueryArray")
	fmt.Println(getQueryArray, res)

	queryMap := c.QueryMap("queryMap")
	fmt.Println(queryMap)

	getQueryMap, res := c.GetQueryMap("getQueryMap")
	fmt.Println(getQueryMap)

	//定义一个User变量
	var user dao.OrderTest
	fmt.Println("c.Params :", c.Params)
	//将调用后端的request请求中的body数据根据json格式解析到User结构变量中
	c.BindJSON(&user)
	fmt.Println("use is :", user)
	//将被转换的user变量传给service层的CreateUser方法，进行User的新建
	err := service.Insert(user)
	//判断是否异常，无异常则返回包含200和更新数据的信息
	if err != nil {
		rsp.Error(c, err.Error())
	} else {
		rsp.Success(c, "新增成功", user)
	}
	todoList := service.Query()

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": todoList,
	})

}

func UpdateOrder(c *gin.Context) {

	param := c.Param("param")
	fmt.Println(param)

	query := c.Query("query")
	fmt.Println(query)

	defaultQuery := c.DefaultQuery("defaultQuery", "no")
	fmt.Println(defaultQuery)

	getQuery, res := c.GetQuery("getQuery")
	fmt.Println(getQuery, res)

	queryArray := c.QueryArray("queryArray")
	fmt.Println(queryArray)

	getQueryArray, res := c.GetQueryArray("getQueryArray")
	fmt.Println(getQueryArray, res)

	queryMap := c.QueryMap("queryMap")
	fmt.Println(queryMap)

	getQueryMap, res := c.GetQueryMap("getQueryMap")
	fmt.Println(getQueryMap)

	//定义一个User变量
	var user dao.OrderTest
	fmt.Println("c.Params :", c.Params)
	//将调用后端的request请求中的body数据根据json格式解析到User结构变量中
	c.BindJSON(&user)
	fmt.Println("use is :", user)
	//将被转换的user变量传给service层的CreateUser方法，进行User的新建
	err := service.Insert(user)
	//判断是否异常，无异常则返回包含200和更新数据的信息
	if err != nil {
		rsp.Error(c, err.Error())
	} else {
		rsp.Success(c, "新增成功", user)
	}
	todoList := service.Query()

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": todoList,
	})

}
