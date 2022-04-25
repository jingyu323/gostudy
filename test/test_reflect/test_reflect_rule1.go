package main

import (
	"fmt"
	"reflect"
)

//
//eflection goes from interface value to reflection object.
//
//为了实现从接口变量到反射对象的转换，需要提到 reflect 包里很重要的两个方法：
//
//reflect.TypeOf(i) ：获得接口值的类型
//
//reflect.ValueOf(i)：获得接口值的值

func main() {
	test_from_interface_reflect()
	test_refect_to_interface()
	test_repflect_can_write()
}

/**
等等，上面我们定义的 age 不是 int 类型的吗？第一法则里怎么会说是接口类型的呢？

关于这点，其实在上一节（关于接口的三个 『潜规则』）已经提到过了，由于 TypeOf 和 ValueOf 两个函数接收的是 interface{} 空接口类型，而 Go 语言函数都是值传递，因此Go语言会将我们的类型隐式地转换成接口类型。
*/
func test_from_interface_reflect() {
	var age interface{} = 25

	fmt.Printf("原始接口变量的类型为 %T，值为 %v \n", age, age)

	t := reflect.TypeOf(age)
	v := reflect.ValueOf(age)

	// 从接口变量到反射对象
	fmt.Printf("从接口变量到反射对象：Type对象的类型为 %T \n", t)
	fmt.Printf("从接口变量到反射对象：Value对象的类型为 %T \n", v)

}
func test_refect_to_interface() {
	var age interface{} = 25

	fmt.Printf("原始接口变量的类型为 %T，值为 %v \n", age, age)

	t := reflect.TypeOf(age)
	v := reflect.ValueOf(age)

	// 从接口变量到反射对象
	fmt.Printf("从接口变量到反射对象：Type对象的类型为 %T \n", t)
	fmt.Printf("从接口变量到反射对象：Value对象的类型为 %T \n", v)

	// 从反射对象到接口变量
	i := v.Interface()
	fmt.Printf("从反射对象到接口变量：新对象的类型为 %T 值为 %v \n", i, i)
}
func test_repflect_can_write() {
	var name string = "Go编程时光"
	fmt.Println("真实世界里 name 的原始值为：", name)

	// 如果不传地址会报错
	v1 := reflect.ValueOf(&name)
	fmt.Println("v1 可写性为:", v1.CanSet())
	v2 := v1.Elem()
	fmt.Println("v2 可写性为:", v2.CanSet())

	v2.SetString("Reflect Test....")
	fmt.Println("通过反射对象进行更新后，真实世界里 name 变为：", name)

}
