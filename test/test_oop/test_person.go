package main

import "fmt"

// 类声明
type Person3 struct {
	Name string
	Age  int
	Doc  []string // slice切片
}

// 类方法声明-传递值对象
func (p *Person3) P1() {
	p.Name = "Tinywan"
	p.Age = 24
}

// 类方法声明-传递指针对象
func (p *Person3) P2() {
	p.Name = "Tinyaiai"
	p.Age = 22
}

func main() {
	p1 := &Person3{}
	p1.Name = "ShaoBo Wan"
	p1.Age = 20

	p2 := &Person3{Name: "HaHa"}
	fmt.Println(p2.Name + "this is p2:")
	p3 := new(Person3)
	p3.Name = "New Name"

	p4 := Person3{}
	p4.Name = "Name 001"
	p4.Age = 26

	p5 := Person3{Name: "Name 002", Age: 28}

	fmt.Println(p5.Name + "this is p5:")

	// 使用中如果包含数组(动态数组 slice切片)，结构体的实例化需要加上类型如上如果intro的类型是[]string
	p6 := &Person3{
		"zhangsan",
		25,
		[]string{"lisi", "wangwu"}}

	fmt.Println(p6.Name + "this is p6:")
}
