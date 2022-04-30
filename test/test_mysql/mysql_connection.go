package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

type MysqlConnection struct {
}

var age int = 19

type MysqlConfig struct {
	App struct {
		Name string `yaml:"name"`
	}
	Mysql struct {
		Host     string `yaml:"host"`
		Port     int32  `yaml:"port"`
		DbName   string `yaml:"dbname"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	}
	Cache struct {
		Enable bool     `yaml:"enable"`
		List   []string `yaml:"list,flow"`
	}
}

func ReadConfigFromYaml(config string) {
	myConf := new(MysqlConfig)
	yamlFile, err := ioutil.ReadFile(config)
	log.Println("yamlFile:", string(yamlFile))
	if err != nil {
		log.Printf("yamlFile.Get err %v", err)
	}
	err = yaml.Unmarshal(yamlFile, myConf)
	if err != nil {
		log.Fatalf("Unmarshal: %v when to struct", err)
	}
	log.Println("conf", myConf)
	log.Printf("mysqlHost:%s, port:%d, cacheList:%s", myConf.Mysql.Host, myConf.Mysql.Port, myConf.Cache.List)

	MyConfigInfo = myConf
	//configMap := make(map[string]interface{})
	//err = yaml.Unmarshal(yamlFile, configMap)

	if err != nil {
		log.Fatalf("Unmarshal: %v when to map", err)
	}

	fmt.Printf("---遍历结束---\n")
	Config3 := myConf
	fmt.Println(Config3.Mysql.DbName)

}

var MyConfigInfo *MysqlConfig
var Config2 *MysqlConfig

func DatabaseDialString() string {
	var Config3 *MysqlConfig = MyConfigInfo
	Config2 = MyConfigInfo
	fmt.Println("ssssc%d", age)
	return fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=%s",
		MyConfigInfo.Mysql.User,
		Config2.Mysql.Password,
		"tcp",
		MyConfigInfo.Mysql.Host,
		Config3.Mysql.Port,
		MyConfigInfo.Mysql.DbName,
		"utf8mb4",
	)
}

type Order struct {
	Orderid   string
	Projectid string
	Status    string
}

var Db *sqlx.DB

func init() {

	fmt.Println("config init")

	ReadConfigFromYaml("./test_mysql/mysql.yaml")
	fmt.Println("user:" + MyConfigInfo.Mysql.User)
	fmt.Println("cnn  init")
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
		err := rows.Scan(&u.Orderid, &u.Projectid, &u.Status) // 可以指定小写属性

		if err != nil {
			fmt.Println("query failed,err:%v\n", err)
			return nil
		}
		orders = append(orders, u)
		fmt.Printf("id:%d name:%s age:%d\n", u.Orderid, u.Projectid, u.Status)
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

	test_update()
	getOrders()
}

func getOrders() {
	var orders []Order
	err := Db.Select(&orders, "select  * from  order_test") // 不能指定小写属性，只能使用大写属性，估计是内部反射实现
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}

	for key, value := range orders {
		fmt.Println(key)
		fmt.Println(value.Orderid)
	}
}

func test_update() {
	res, err := Db.Exec("update   order_test set projectid='919' where orderid=?", 3)

	if err != nil {
		fmt.Println("exec update failed ,", err)
	}

	fmt.Println("exec success res is:", res)

	row, err := res.RowsAffected()
	if err != nil {
		fmt.Println("rows failed, ", err)
	}
	fmt.Println("update succ:", row)

}
