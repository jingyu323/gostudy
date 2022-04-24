package main

import "fmt"

// 定义一个名为Profile 的结构体
type Profile struct {
	name   string
	age    int
	gender string
	mother *Profile // 指针
	father *Profile // 指针
}

// 定义一个与 Profile 的绑定的方法
func (person Profile) FmtProfile() {
	fmt.Printf("名字：%s\n", person.name)
	fmt.Printf("年龄：%d\n", person.age)
	fmt.Printf("性别：%s\n", person.gender)
}

// 重点在于这个星号: *
func (person *Profile) increase_age() {
	person.age += 1
}

type Phone interface {
	call()
}
type Nokia struct {
	name string
}

// 接收者为 Nokia
func (phone Nokia) call() {
	fmt.Println("我是 Nokia，是一台电话")
}

func main() {

	// 实例化
	myself := Profile{name: "小明", age: 24, gender: "male"}
	// 调用函数
	myself.FmtProfile()

	myself.increase_age()
	fmt.Printf("当前年龄：%d", myself.age)

	p1 := &Profile{name: "iswbm"}
	fmt.Println(p1.name) // output: iswbm

	fmt.Println((*p1).name)

	xm := new(Profile)
	// 等价于: var xm *Profile = new(Profile)
	fmt.Println(xm)
	// output: &{ 0 }

	xm.name = "iswbm222"

	fmt.Println(xm.name) // output: iswbm

	iPhone := Nokia{
		name: "iPhone",
	}
	fmt.Println(iPhone.name)
}
