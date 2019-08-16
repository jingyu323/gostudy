package main

import "fmt"

//map[K]T
func main() {

	mm := map[int]string{1: "a", 2: "b", 3: "c"}

	b := mm[2]
	fmt.Println(b)

	mm[2] = b + "2"

	mm[2] = b + "2"
	fmt.Println(mm)

	e, ok := mm[5]

	fmt.Println(e)
	fmt.Println(ok)
	e1, ok1 := mm[2]

	fmt.Println(e1)
	fmt.Println(ok1)

	delete(mm, 4)

	fmt.Println(mm)

	/***
	make(chan int, 5)
	与其他的数据类型不同，我们无法表示一个通道类型的值，因此，我们无法用字面量来为通道类型的变量赋值。
	只能通过调用内建函数make来达到目的。make参数可接受两个参数，
	第一个参数是代表了将被初始化的值的类型的字面量(例: chan int),而第二个参数则是值的长度，例如,若我们想要初始化一个长度为5且元素类型为int的通道值
	*/
	ch1 := make(chan string, 5)

	ch1 <- "value1"
	//可以直接从通道中去除来值
	//fmt.Printf(<- ch1)

	value, ok := <-ch1

	fmt.Println(value)
	fmt.Println(ok)

	ch2 := make(chan string, 34)
	ch2 <- "value1"

}
