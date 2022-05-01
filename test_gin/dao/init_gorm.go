package dao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"test_gin/config"
	"test_gin/mysql_connection"
)


var db *gorm.DB
var err error
var Config = config.MyConfigInfo

type OrderTest struct {
	Orderid       string `gorm:"column:orderid;type:varchar(255);size(64);not null"`
	Projectid string `gorm:"column:projectid;not null"`
	Status  string


}

func init() {
	fmt.Println("dao init")
	InitGorm()
}
func (t *OrderTest) TableName() string {
	return "order_test"
}
func  InitGorm()  {
	fmt.Println("InitGorm init")
	var err error
	dsn := mysql_connection.DatabaseDialString()
	fmt.Println(dsn)

	db, err := gorm.Open("mysql",
		"root:root@tcp(localhost:3306)/tets?charset=utf8mb4&parseTime=True&loc=Local")

	fmt.Println("db is:",db)
	defer db.Close()

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}
	if db.Error != nil {
		fmt.Printf("database error %v", db.Error)
	}
	//var order_test  []OrderTest
	//
	//if err := db.Find(&order_test).Error; err != nil {
	//
	//	fmt.Println(err)
	//}
	//
	////fmt.Println(len(order_test))
	//fmt.Println( order_test)

	fmt.Println("2323db is:",db)
}

//插入数据
func (order *OrderTest) Insert()  {
	fmt.Println("create start")
	//这里使用了Table()函数，如果你没有指定全局表名禁用复数，或者是表名跟结构体名不一样的时候
	//你可以自己在sql中指定表名。这里是示例，本例中这个函数可以去除。
	db.Table(order.TableName()).Create(order)
}

func(order *OrderTest) Update()  {
	fmt.Println("update start")
	db.Model(&order).Update(order)

	//注意到上面Update中使用了一个Struct，你也可以使用map对象。
	//需要注意的是：使用Struct的时候，只会更新Struct中这些非空的字段。
	//对于string类型字段的""，int类型字段0，bool类型字段的false都被认为是空白值，不会去更新表

	//下面这个更新操作只使用了where条件没有在Model中指定id
	//update user set name='xiaohong' wehre sex=1
	//db.Model(&User{}).Where("sex = ?",1).Update("name","xiaohong")
	//如果你想手动将某个字段set为空值, 可以使用单独选定某些字段的方式来更新：
	//user := User{Id: 1}
	//db.Model(&user).Select("name").Update(map[string]interface{}{"name":"","age":0})
}
func(order *OrderTest) Delete()  {
	fmt.Println("Delete start")

}

func (order *OrderTest) Query() (u []OrderTest) {
	fmt.Println("dddd is :",db)

	if err := db.Find(&u).Error; err != nil {

		fmt.Println(err)
	}



	////Find方法可以带 where 参数
	//db.Find(&u,"id > ? and age > ?",2,12)
	//
	////带where 子句的查询，注意where要在find前面
	//db.Where("id > ?", 2).Find(&u)
	//
	//// where name in ("xiaoming","xiaohong")
	//db.Where("name in (?)",[]string{"xiaoming","xiaohong"}).Find(&u)
	//
	////获取第一条记录，按照主键顺序排序
	//db.First(&u)
	//
	////First方法可以带where 条件
	//db.First(&u,"where sex = ?",1)

	//获取最后一条记录，按照主键顺序排序
	//同样 last方法也可以带where条件
	//db.Last(&u)

	return u
}
