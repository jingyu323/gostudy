package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type OrderTest struct {
	Orderid   string `gorm:"column:orderid;type:varchar(255);size(64);not null"`
	Projectid string `gorm:"column:projectid;not null"`
	Status    string `gorm:"column:status;not null"`
}

var db *gorm.DB

func main() {
	db, err := gorm.Open("mysql",
		"root:root@tcp(localhost:3306)/tets?charset=utf8mb4&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("connection succedssed")
	}
	// Migrate the schema
	//db.AutoMigrate(&OrderTest{})
	defer db.Close()

	fmt.Println("ss")

	var order_test []OrderTest

	if err := db.Find(&order_test).Error; err != nil {

		fmt.Println(err)
	}

	//fmt.Println(len(order_test))
	fmt.Println(order_test)

}
func (t *OrderTest) TableName() string {
	return "order_test"
}
func GetProjects(c *gin.Context) {
	var people []OrderTest
	if err := db.Find(&people).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, people)
	}
}
