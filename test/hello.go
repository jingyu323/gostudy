package main

import (
	"./clent"
	"fmt"
	"io"
	"strings"
) //系统包用来输出的

//常量定义区

const NAME = "name "

/*

 */
func main() {
	// 打印函数调用语句。用于打印输出信息。
	fmt.Println("hello world")

	var a int = 65
	fmt.Println(a)

	clentt := clent.NewClientTest("host2")
	clentt.SetClentName("this is hoost") //设置clentName的值

	fmt.Println("ClentName is " + clentt.ClentName()) // 获取ClentName的值
	fmt.Println("Host is " + clentt.Host)

	// test for  cycle

	sum := 0
	for i := 0; i < 10; i++ {
		if i > 3 {
			break //可以直接跳出循环
		}
		sum += i
	}

	fmt.Println(sum)

	fmt.Println(unhex('A'))
	fmt.Println(shouldEscape('?'))
	fmt.Println(shouldEscape('='))

	var str string = "test"

	var data []byte = []byte(str)

	reader := strings.NewReader("test")
	len, err := ReadFull(reader, data)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("buffer size is:", len)

}

func unhex(c byte) byte {
	switch {
	case '0' <= c && c <= '9':
		return c - '0'
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10
	}
	return 0
}

//seitch 可以逗号来分割相同的条件
func shouldEscape(c byte) bool {
	switch c {
	case ' ', '?', '&', '=', '#', '+', '%':
		return true
	}
	return false
}

func ReadFull(r io.Reader, buf []byte) (n int, err error) {
	for len(buf) > 0 && err == nil {
		var nr int
		nr, err = r.Read(buf)
		n += nr
		buf = buf[nr:]
	}
	return
}
