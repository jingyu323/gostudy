package main

import "fmt"

func foo() int {
	v := 1024
	return v
}

func main() {
	m := foo()
	fmt.Println(m)
	b := foo1()
	fmt.Println(b)
}

func foo1() []int {
	a := []int{1, 2, 3}
	return a

}
