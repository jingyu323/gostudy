package main

import "fmt"

type Animal interface {
	shout() string
}

type Dog struct{}

func (self Dog) shout() string {
	return fmt.Sprintf("wang wang")
}

type Cat struct{}

func (self Cat) shout() string {
	return fmt.Sprintf("miao miao")
}
func main() {

	var animal Animal = Dog{}

	switch t := animal.(type) {
	case Cat:
		fmt.Printf("Dog", t)
	default:
		fmt.Printf("default")

	}

}
