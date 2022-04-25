package main

import "fmt"

type Phone interface {
	call()
}

type iPhone struct {
	name string
}

func (phone iPhone) call() {
	fmt.Println("Hello, iPhone.")
}

func (phone iPhone) send_wechat() {
	fmt.Println("Hello, Wechat.")
}

func main() {

	test_interface_rule1()
	a := 10
	test_interface_rule2(a)

	test_interface_rule3()
}

/**
接口是一组固定的方法集，由于静态类型的限制，接口变量有时仅能调用其中特定的一些方法

特定声明类型 会限制只能调用接口中的方法，如果不声明接口则可以调用
*/
func test_interface_rule() {
	var phone Phone
	phone = iPhone{name: "ming's iphone"}
	phone.call()
	//phone.send_wechat()
}
func test_interface_rule1() {

	phone2 := iPhone{name: "ming's 2 iphone"}
	phone2.call()
	phone2.send_wechat()
}

/**
调用函数时的隐式转换
Go 语言中的函数调用都是值传递的，变量会在方法调用前进行类型转换
*/

func test_interface_rule2(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("参数的类型是 int")
	case string:
		fmt.Println("参数的类型是 string")
	}
}

/**
如何进行接口类型的显示转换
*/
func test_interface_rule3() {
	a := 10

	switch interface{}(a).(type) {
	case int:
		fmt.Println("参数的类型是 int")
	case string:
		fmt.Println("参数的类型是 string")
	}
}
