package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Addr string `json:"addr,omitempty"`
}
type Person2 struct {
	Name   string `label:"Name is: "`
	Age    int    `label:"Age is: "`
	Gender string `label:"Gender is: " default:"unknown"`
}

func Print(obj interface{}) error {
	// 取 Value
	v := reflect.ValueOf(obj)

	// 解析字段
	for i := 0; i < v.NumField(); i++ {

		// 取tag
		field := v.Type().Field(i)
		tag := field.Tag

		// 解析label 和 default
		label := tag.Get("label")
		defaultValue := tag.Get("default")

		value := fmt.Sprintf("%v", v.Field(i))
		if value == "" {
			// 如果没有指定值，则用默认值替代
			value = defaultValue
		}

		fmt.Println(label + value)
	}

	return nil
}
func main() {
	fmt.Println("s")
	p1 := Person{
		Name: "Jack",
		Age:  22,
	}

	data1, err := json.Marshal(p1)
	if err != nil {
		panic(err)
	}
	// p1 没有 Addr，就不会打印了
	fmt.Printf("%s\n", data1)
	p2 := Person{
		Name: "Jack",
		Age:  22,
		Addr: "China",
	}

	data2, err := json.Marshal(p2)
	if err != nil {
		panic(err)
	}

	// p2 则会打印所有p
	fmt.Printf("%s\n", data2)
	// f反射
	p := reflect.TypeOf(Person{})
	name, _ := p.FieldByName("name")

	fmt.Printf("%q\n", name.Tag) //输出 ""

	Print(p2)
}
