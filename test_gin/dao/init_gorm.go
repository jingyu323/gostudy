package dao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"test_gin/config"
	"test_gin/mysql_connection"
)

type OrderTest struct {
	Orderid   string `gorm:"column:orderid;type:varchar(255);size(64);not null"`
	Projectid string `gorm:"column:projectid;not null"`
	Status    string
}

var Db *gorm.DB
var SqlSession *gorm.DB
var err error
var Config = config.MyConfigInfo

func InitGorm() *gorm.DB {
	fmt.Println("InitGorm init")
	var err error
	dsn := mysql_connection.DatabaseDialString()
	fmt.Println(dsn)

	Db, err = gorm.Open("mysql",
		"root:root@tcp(localhost:3306)/tets?charset=utf8mb4&parseTime=True&loc=Local")

	fmt.Println("db is:", Db)
	//defer db.Close()

	Db.DB().SetMaxIdleConns(10)
	Db.DB().SetMaxOpenConns(100)

	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}
	if Db.Error != nil {
		fmt.Printf("database error %v", Db.Error)
	}
	var order_test []OrderTest

	if err := Db.Find(&order_test).Error; err != nil {

		fmt.Println(err)
	}

	//fmt.Println(len(order_test))
	fmt.Println(order_test)

	fmt.Println("2323db is:", Db)
	SqlSession = Db
	return Db
}

//关闭数据库连接
func Close() {
	fmt.Println("cloaddd....")
	SqlSession.Close()
}
func (OrderTest) TableName() string {
	return "order_test"
}
