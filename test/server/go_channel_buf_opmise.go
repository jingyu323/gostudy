package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		v := <-ch
		fmt.Println(v)

		//<-ch
		//fmt.Println("1")
	}()

	ch <- 1 //取出channel的子值
	fmt.Println("2")

	//这里就不用让主线程休眠了，因为channel在主线程中被赋值后，主线程就会阻塞，直到channel的值在子线程中被取出

}
