package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// os.Open
func read1(filePath string) string {

	f, err := os.Open(filePath)
	defer f.Close()

	if err != nil {
		panic(err)
	}

	chunks := make([]byte, 1024, 1024)
	buf := make([]byte, 1024)
	for {
		n, err := f.Read(buf)
		if err != nil || 0 == n {
			break
		}
		chunks = append(chunks, buf[:n]...)
	}

	return string(chunks)

}

//bufio
func read2(filePath string) string {
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	chunks := make([]byte, 1024, 1024)
	f := bufio.NewReader(file)
	buf := make([]byte, 1024)
	for {
		n, err := f.Read(buf)
		if err != nil || n == 0 {
			break
		}

		chunks = append(chunks, buf[:n]...)
	}

	return string(chunks)
}

//ioutil
func read3(filePath string) string {
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	f, _ := ioutil.ReadAll(file)

	text := string(f)

	return text

} //ioutil
func read4(filePath string) string {

	result := ""
	if contents, err := ioutil.ReadFile(filePath); err == nil {
		//直接使用\n根本无法替换成功，正确的做法是对\n再进行转义：\\n
		result = strings.Replace(string(contents), "\\n", "", -1)
		fmt.Print("Use ioutil.ReadFile to read a file:", result)
	}

	return result

}

func main() {
	content := read1("./resource/testfile.txt")
	content2 := read3("./resource/testfile.txt")
	content3 := read4("./resource/testfile.txt")

	fmt.Printf(content)
	fmt.Println()
	fmt.Println("************")

	fmt.Printf(content2)
	fmt.Println()
	fmt.Println("************")
	fmt.Printf("result=" + content3)

}
