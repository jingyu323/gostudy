package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"sync"
	"syscall"
)

func main() {

}

func test_os_read() {
	content, err := os.ReadFile("./test.txt")
	panic(err)
	fmt.Println(string(content))
}

//其实在 Go 1.16 开始，ioutil.ReadFile 就等价于 os.ReadFile，二者是完全一致的
func test_ioutil_read_file() {
	content, err := ioutil.ReadFile("./test.txt")
	panic(err)
	fmt.Println(string(content))

}

func test_reate_refer_reade() {
	file, err := os.Open("a.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	fmt.Println(string(content))
}

func test_os_openfile() {
	file, err := os.OpenFile("a.txt", os.O_RDONLY, 0)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	fmt.Println(string(content))
}

//bufio.ReadBytes 和 bufio.ReadString 去读取单行数据。
func read_one_line() {
	file, err := os.OpenFile("./test.txt", os.O_RDONLY, 0)

	if err != nil {
		panic(err)
	}

	r := bufio.NewReader(file)

	for {

		lineBytes, err := r.ReadBytes('\n')
		line := strings.TrimSpace(string(lineBytes))

		if err != nil && err != io.EOF {
			panic(err)
		}

		if err == io.EOF {
			break
		}
		fmt.Println(line)

	}

}

func test_read_one_line2() {
	file, err := os.OpenFile("./test.txt", os.O_RDONLY, 0)

	if err != nil {
		panic(err)
	}

	// 创建 Reader
	r := bufio.NewReader(file)
	for {
		line, err := r.ReadString('\n')
		line = strings.TrimSpace(line)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if err == io.EOF {
			break
		}
		fmt.Println(line)
	}

}

func test_read_read_byte() {
	// 创建句柄
	fi, err := os.Open("a.txt")
	if err != nil {
		panic(err)
	}

	// 创建 Reader
	r := bufio.NewReader(fi)
	// 每次读取 1024 个字节
	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}

		if n == 0 {
			break
		}
		fmt.Println(string(buf[:n]))
	}

}

func test_syscall() {

	fd, err := syscall.Open("christmas_apple.py", syscall.O_RDONLY, 0)
	if err != nil {
		fmt.Println("Failed on open: ", err)
	}
	defer syscall.Close(fd)

	var wg sync.WaitGroup
	wg.Add(2)
	dataChan := make(chan []byte)
	go func() {
		wg.Done()
		for {
			data := make([]byte, 100)
			n, _ := syscall.Read(fd, data)
			if n == 0 {
				break
			}
			dataChan <- data
		}
		close(dataChan)
	}()

	go func() {
		defer wg.Done()
		for {
			select {
			case data, ok := <-dataChan:
				if !ok {
					return
				}

				fmt.Printf(string(data))
			default:

			}
		}
	}()
	wg.Wait()

}
