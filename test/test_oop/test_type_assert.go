package main

import "fmt"

//这个表达式可以断言一个接口对象（i）里不是 nil，并且接口对象（i）存储的值的类型是 T，如果断言成功，就会返回值给 t，如果断言失败，就会触发 panic。
func main() {
	var i interface{} = 10
	t1 := i.(int)
	fmt.Println(t1)

	fmt.Println("=====分隔线=====")

	//t2 := i.(string)
	//fmt.Println(t2) // 会引发错误 panic: interface conversion: interface {} is int, not string	goroutine 1 [running]:

	var ii interface{} = 11
	t1, ok := ii.(int)
	fmt.Printf("%d-%t\n", t1, ok)

	fmt.Println("=====分隔线1=====")
}
