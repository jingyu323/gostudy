package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1)
	ch <- 1
	go func() {
		<-ch
		fmt.Println("1")
	}()
	time.Sleep(1 * time.Second)
	fmt.Println("2")

}
