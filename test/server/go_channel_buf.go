package main

import "fmt"

func main() {
	ch := make(chan int)
	ch <- 1
	go func() {
		<-ch
		fmt.Println("1")
	}()
	fmt.Println("2")
}
