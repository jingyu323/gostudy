package main

import (
	"fmt"
	"strconv"
)

type User struct {
	name string
	age  int
}

func New(name string, age int) *User {
	return &User{name: name, age: age}
}

func (u *User) String() string {
	return "name: " + u.name + " age: " + strconv.Itoa(u.age)
}

func main() {
	u1 := User{name: "test", age: 33}
	fmt.Println(u1.name)
	fmt.Println("u1:", u1)
	u1.name = "test2"
	fmt.Println("u1:", u1)
	fmt.Println()

	u2 := new(User)
	u2.name = "hello"
	u2.age = 12
	fmt.Println("u2: ", u2)
	u2.name = "test2"
	fmt.Println("u2: ", u2)

	fmt.Println()
	u3 := &User{name: "test", age: 32}
	fmt.Println("u3: ", u3)
	u3.name = "test2"
	fmt.Println("u3: ", u3)

	fmt.Println()
	u4 := New("test", 34)
	fmt.Println("u4: ", u4)
	u4.name = "test4"
	fmt.Println("u4: ", u4)

}

/**
四种创建方式：
1.字典形式
2.添加构造
3.&T{}
*/
