package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

//先处理错误避免嵌套
type Gopher struct {
	Name     string
	AgeYears int
}

func main() {
	gofer := &Gopher{
		"223", 9,
	}
	file, err := os.OpenFile("./gc2.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("打开文件失败,错误为:%v\n", err)
		return
	}
	defer file.Close() //关闭文件

	gofer.WriteTo(file) // ??? io.writer 竟然可以用file

}

func (g *Gopher) WriteTo(w io.Writer) (size int64, err error) {
	err = binary.Write(w, binary.LittleEndian, int32(len(g.Name)))
	if err != nil {
		return
	}
	size += 4
	n, err := w.Write([]byte(g.Name))
	size += int64(n)
	if err != nil {
		return
	}
	err = binary.Write(w, binary.LittleEndian, int64(g.AgeYears))
	if err == nil {
		size += 4
	}
	return
}
