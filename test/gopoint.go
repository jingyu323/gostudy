package main

import "fmt"

type User struct {
	name string
	age  int
}

func main() {

	andes := User{
		name: "andesname",
		age:  18,
	}

	fmt.Printf("点操作符号：" + andes.name)
	fmt.Println()

	andes.name = "11"
	p := &andes
	fmt.Println("指针获取的字段name:" + p.name)
	p.name = "22" //用指针修改对象的值，原来的值也会变，因为指针只是指向了地址，改变了地址的值当然对象的值也就变了

	fmt.Println("指针获取的字段name:" + p.name)
	fmt.Println("andes -->:" + andes.name)

	//GO 不支持指针运算
	//为啥？ 因为GO有支持垃圾回收 带来很多不便
	a := 2
	b := 3

	sum1 := sum(a, b)
	sum2 := sum2(a, b)

	fmt.Println(sum1)
	fmt.Println(sum2)

	var num2323 = [3]int{1, 1, 1}
	fmt.Println(num2323)

	var number3 = [5]int{1, 2, 3, 4, 5}
	//[1:4]的结果为[]int{2,3,4}很明显被切下的部分不包含元素上界索引指向的元素
	var slice1 = number3[1:4]

	fmt.Println("----->>>>")
	fmt.Println(slice1)

	//一个切片值的容量即为它的第一个元素值在其底层数组中的索引值与该数组长度的差值的绝对值  即5-1 =4
	var capacity2 int = cap(slice1)

	fmt.Println(capacity2)

}

func sum(a, b int) *int {
	sum := a + b
	return &sum //允许， sum会分配在heap上
}

func sum2(a, b int) int {
	sum := a + b
	fmt.Println("sun= %s", sum)
	return sum
}
