package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	name   string
	age    int
	gender string
}

func (p Person) SayBye() {
	fmt.Print("Bye")
}

func (p Person) SayHello() {
	fmt.Println("Hello")
}

type Profile struct {
	name   string
	age    int
	gender string
}

func main() {
	test_kind()
	test_point()
}

func test_kind() {
	m := Profile{}

	t := reflect.TypeOf(m)
	fmt.Println("Type: ", t)
	fmt.Println("Kind: ", t.Kind())
}

func test_point() {
	m := Profile{}

	t := reflect.TypeOf(&m)

	fmt.Println("&m Type: ", t)
	fmt.Println("&m Kind: ", t.Kind())

	fmt.Println("m Type: ", t.Elem())
	fmt.Println("m Kind: ", t.Elem().Kind())

	p := &Person{"写代码的明哥", 27, "male"}

	v := reflect.ValueOf(p)

	v.MethodByName("SayHello").Call(nil)
	v.MethodByName("SayBye").Call(nil)

	fmt.Println("&m Type: ", t)
	fmt.Println("&m Kind: ", t.Kind())

	fmt.Println("m Type: ", t.Elem())
	fmt.Println("m Kind: ", t.Elem().Kind())

	v1 := reflect.ValueOf(&m)

	fmt.Println("&m Type: ", v1.Type())
	fmt.Println("&m Kind: ", v1.Kind())

	fmt.Println("m Type: ", v1.Elem().Type())
	fmt.Println("m Kind: ", v1.Elem().Kind())

	var score float64 = 99.5

	v11 := reflect.ValueOf(score)
	fmt.Printf("转换前， type: %T, value: %v \n", v1, v1)
	v2 := v11.Float()
	fmt.Printf("转换后， type: %T, value: %v \n", v2, v2)

}
