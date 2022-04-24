package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("fff")

	//file_read()
	file_read_all_file()
	file_read_wirte()

}

/**
os.Open 用于只读
*/
func file_read() {
	file, err := os.Open("./test.txt") //只是用来读的时候，用os.Open。相对路径，针对于同目录下。
	if err != nil {
		file, err = os.Create("./test.txt")
		fmt.Printf("打开文件失败,err:%v\n", err)

	}
	var buffer bytes.Buffer
	buffer.WriteString("abcd5")
	file.WriteString("ttttt12")
	file.Write([]byte("sdfd7"))
	file.Write(buffer.Bytes()) //写入字节
	defer file.Close()         //关闭文件,为了避免文件泄露和忘记写关闭文件

	var qiepian = make([]byte, 128) // 大于128之后的字节读不出来，需要循环读取
	n, err := file.Read(qiepian)    //file.Read里面是一个切片.如果读取成功,那么值声明变量n并同时赋值给n。
	//判断
	if err != nil { //if err != nil 意思是如果err不为空，那么就说明err有错误值，打印读取错误。如果为空，说明err没有错误值，打印读取成功
		fmt.Printf("读取文件错误,错误值为:%v\n", err)
		return
	}

	fmt.Printf("读取文件成功,文件字节大小为:%d\n", n)

	fmt.Println(string(qiepian))
}

func file_read_all_file() {
	file, err := os.Open("./test.txt") //只是用来读的时候，用os.Open。相对路径，针对于同目录下。
	if err != nil {
		file, err = os.Create("./test.txt")
		fmt.Printf("打开文件失败,err:%v\n", err)

	}
	var buffer bytes.Buffer
	buffer.WriteString("abcd5")
	file.WriteString("ttttt12")
	file.Write([]byte("sdfd7"))
	file.Write(buffer.Bytes()) //写入字节
	defer file.Close()         //关闭文件,为了避免文件泄露和忘记写关闭文件

	for {
		var qiepian = make([]byte, 128) // 大于128之后的字节读不出来，需要循环读取
		n, err := file.Read(qiepian)    //file.Read里面是一个切片.如果读取成功,那么值声明变量n并同时赋值给n。
		//判断
		if err != nil { //if err != nil 意思是如果err不为空，那么就说明err有错误值，打印读取错误。如果为空，说明err没有错误值，打印读取成功
			fmt.Printf("读取文件错误,错误值为:%v\n", err)
			return
		}

		if err == io.EOF { //EOF表示 end of file，就是文件的末尾，这个判断意思就是，读到文件末尾的时候，将当前读了多少字节打印出来并退出
			//把当前读了多少字节的数据打印出来，然后退出
			fmt.Println(string(qiepian))
			return
		}

		fmt.Printf("读取文件成功,文件字节大小为:%d\n", n)

		fmt.Println(string(qiepian))
		jobDowhile()
		readBybuffio()

		fmt.Println("ioutil start")
		readByioutil()
	}

}

func file_read_wirte() {

	//os.O_CREATE:创建
	//os.O_WRONLY:只写
	//os.O_APPEND:追加
	//os.O_RDONLY:只读
	//os.O_RDWR:读写
	//os.O_TRUNC:清空

	//0644:文件的权限
	//如果没有test.txt这个文件那么就创建，并且对这个文件只进行写和追加内容。
	file, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {

		fmt.Printf("打开文件失败,err:%v\n", err)
	}
	file.WriteString("\n")
	file.WriteString("write file content")
	defer file.Close()

}

//golang里面没有while关键字，可以用for+break实现。
func readBybuffio() {
	//打开文件
	file, err := os.Open("./test.txt") //只是用来读的时候，用os.Open。相对路径，针对于同目录下。
	if err != nil {
		fmt.Printf("打开文件失败,err:%v\n", err)
		return
	}
	defer file.Close() //关闭文件,为了避免文件泄露和忘记写关闭文件

	//使用buffio读取文件内容
	reader := bufio.NewReader(file) //创建新的读的对象
	for {
		line, err := reader.ReadString('\n') //注意是字符，换行符。
		if err == io.EOF {
			fmt.Println("文件读完了")
			break
		}
		if err != nil { //错误处理
			fmt.Printf("读取文件失败,错误为:%v", err)
			return
		}
		fmt.Print(line)
	}

}

//模拟do……while实现输出10次hello,world（先做后判断）
func jobDowhile() {
	var i = 0
	for {
		fmt.Println("hello,world", i)
		i++
		if i >= 10 {
			break
		}
	}
}

//ioutuil ioutil读取整个文件：
func readByioutil() {
	file, err := ioutil.ReadFile("./test.txt")
	if err != nil {
		fmt.Println("file open fail")
		return
	}
	fmt.Println(string(file)) //打印文件内容
}
