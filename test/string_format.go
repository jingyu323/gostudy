package main

import (
	"bytes"
	"fmt"
	"strings"
)

const sourceStr string = "hello world"

type person struct {
	name string
	age  int
}

func main() {

	fmt.Println(strings.Contains(sourceStr, "ll"))

	//统计子符串的次数
	fmt.Println(strings.Count("hello hello", "ll"))
	//判断前缀字符串
	fmt.Println(strings.HasPrefix("hello", "he"))
	//判断后缀字符串
	fmt.Println(strings.HasSuffix("hello", "lo"))
	//找查子符串第一次出现的位置，不存在返回-1
	fmt.Println(strings.Index("hello", "el"))
	//拼接字符串
	fmt.Println(strings.Join([]string{"a", "b", "c"}, ","))
	//重复字符串
	fmt.Println(strings.Repeat("hello", 3))
	//替换前n个字符串
	fmt.Println(strings.Replace("hello", "l", "x", 1))
	//替换所有字符串
	fmt.Println(strings.Replace("hello", "l", "x", -1))

	fmt.Printf("rrrr")
	fmt.Printf("rrrr")
	fmt.Printf("rrrr")
	fmt.Println("line")
	fmt.Println("line")
	fmt.Println("line")
	fmt.Println("line")

	//格式化输出
	var a = person{name: "test", age: 22}

	//打印出结构体
	fmt.Printf("%v\n", a)
	//格式化输出Go语法表示方式
	fmt.Printf("%#v\n", a)
	//输出类型
	fmt.Printf("%T\n", a)
	//格式化布尔型变量
	fmt.Printf("%t\n", true)
	//输出整型
	fmt.Printf("%d\n", 123)
	//输出二进制
	fmt.Printf("%b\n", 10)
	//输出对应字符
	fmt.Printf("%c\n", 65)
	//输出十六进制表示
	fmt.Printf("%x\n", 256)
	//输出浮点数
	fmt.Printf("%f\n", 12.45)
	//科学计数法
	fmt.Printf("%e\n", 1230000000.0)
	fmt.Printf("%E\n", 1230000000.0)
	//输出字符串
	fmt.Printf("%s\n", "hello")
	//输出指针的值
	fmt.Printf("%p\n", &a)
	//控制输出宽度，并用0补齐，输出默认右对齐的
	fmt.Printf("%020d\n", 345)

	//字符串截取 下表是从0开始
	fmt.Println(sourceStr[0:8])

	//字符串拼接

	//1.直接拼劲儿
	str := sourceStr + "000"
	fmt.Println(str)

	//2.Sprintf 字符串输出
	str = fmt.Sprintf("%s%s%s%s%s", str, "123123", "-3333", "-4444", "-5555")
	fmt.Println(str)

	//3.[]string利用strings.Join进行拼接
	s := "1231"
	s = strings.Join([]string{s, "ewrwer"}, "")
	fmt.Println(s)

	//4利用[]string和slice的append的性质进行拼接
	s2 := [3]string{"1", "2", "3"} //声明变量
	fmt.Println(s2)
	var s1 []string //声明类型
	s1 = append(s1, "123123", "666", "0000", "77777")
	fmt.Println(strings.Join(s1, ""))

	//5.使用bytes.Buffer 只能使用字符串
	var buf bytes.Buffer
	buf.WriteString("12312")
	buf.WriteString("werwer")

	fmt.Println(buf.String())

}
