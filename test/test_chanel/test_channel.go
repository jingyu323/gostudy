package main

import (
	"fmt"
	"time"
)

//定义只写信道类型
type Sender = chan<- int

//定义只读信道类型
type Receiver = <-chan int

func main() {
	var pipline = make(chan int)

	go func() {
		var sender Sender = pipline
		fmt.Println("准备发送数据: 100")
		sender <- 100
	}()

	go func() {
		var receiver Receiver = pipline
		num := <-receiver
		fmt.Printf("接收到的数据是: %d", num)
	}()
	// 主函数sleep，使得上面两个goroutine有机会执行
	time.Sleep(time.Second)

	pipline1 := make(chan int, 10)

	go fibonacci(pipline1)

	for k := range pipline1 {
		fmt.Println(k)
	}

}

func fibonacci(mychan chan int) {
	n := cap(mychan)
	x, y := 1, 1
	for i := 0; i < n; i++ {
		mychan <- x
		x, y = y, x+y
	}
	// 记得 close 信道
	// 不然主函数中遍历完并不会结束，而是会阻塞。会报错
	//fatal error: all goroutines are asleep - deadlock!
	//
	//goroutine 1 [chan receive]:
	close(mychan)
}
