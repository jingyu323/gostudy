package main

import "fmt"

func main() {
	pipline := make(chan string)
	go func() {
		pipline <- "hello world"
		pipline <- "hello China"
		close(pipline)
	}()
	for data := range pipline {
		fmt.Println(data)
	}
	test_cache_chanel()
	test_cache_chanel2()
}

//
//<-chan 表示这个信道，只能从里发出数据，对于程序来说就是只读
//
//chan<- 表示这个信道，只能从外面接收数据，对于程序来说就是只写
//需要保证接收者程序在发送数据到信道前就进行阻塞状态
func hello1(pipline chan string) {
	fmt.Println("ddd")
	r := <-pipline
	fmt.Println(r)
}

func test_cache_chanel() {
	pipline := make(chan string)
	go hello(pipline)
	fmt.Println("..944499.")
	pipline <- "hello world22"
}

func hello(pipline chan string) {
	//fmt.Println("..999.")
	//res := <-pipline
	fmt.Println(<-pipline)
	//fmt.Println(res)
}

func test_cache_chanel2() {
	pipline := make(chan string, 1)
	//go hello(pipline)
	fmt.Println("..944499.")
	pipline <- "hello world 2b"
	fmt.Println(<-pipline)
}
