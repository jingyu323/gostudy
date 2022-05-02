package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
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

func CreateOrder(c *gin.Context) {

	//定义一个User变量
	var user dao.OrderTest
	//将调用后端的request请求中的body数据根据json格式解析到User结构变量中
	c.BindJSON(&user)
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
