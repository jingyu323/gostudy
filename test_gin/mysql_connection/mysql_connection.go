package mysql_connection

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"test_gin/config"
)

type MysqlConnection struct {
}

var Config = config.MyConfigInfo

func DatabaseDialString() string {
	return fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=%s",
		Config.Mysql.User,
		Config.Mysql.Password,
		"tcp",
		Config.Mysql.Host,
		Config.Mysql.Port,
		Config.Mysql.DbName,
		"utf8mb4",
	)
}

type Order struct {
	Orderid   string
	projectid string
	Status    string
}

var Db *sqlx.DB

func init() {
	str := DatabaseDialString()
	fmt.Println(str)
	database, err := sqlx.Open("mysql", str)
	if err != nil {
		fmt.Println("scan failed ,err:%v/n", err)
		return
	}
	Db = database

}
func GetAllOrders() []Order {
	var orders []Order
	rows, err := Db.Query("select  * from  order_test")
	if err != nil {
		fmt.Println("scan failed ,err:%v/n", err)
		return nil
	}

	for rows.Next() {
		var u Order
		err := rows.Scan(&u.Orderid, &u.projectid, &u.Status)

		if err != nil {
			fmt.Println("query failed,err:%v\n", err)
			return nil
		}
		orders = append(orders, u)
		fmt.Printf("id:%d name:%s age:%d\n", u.Orderid, u.projectid, u.Status)
	}
	fmt.Println(len(orders))
	return orders

}
func main() {
	orders := GetAllOrders()
	fmt.Print("orders:is")
	fmt.Println(len(orders))

	for key, v := range orders {
		fmt.Println("key is %s", key)
		fmt.Println(v)
	}

	data, err := json.Marshal(&orders)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("ttt      %s\n", data)
}
