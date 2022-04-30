package main

import (
	"fmt"
	"testing"
)

func foo() int {
	v := 1024
	return v
}

func main() {
	m := foo()
	fmt.Println(m) // 逃逸到heap
	b := foo1()
	fmt.Println(b)

	println(testing.AllocsPerRun(1, func() { bar() }))

}

func foo1() []int {
	a := []int{1, 2, 3}
	return a

}

func bar() (*int, *int) {
	m, n := 21, 22
	return &m, &n
}
