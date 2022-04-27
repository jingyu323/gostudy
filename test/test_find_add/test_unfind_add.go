package main

import "fmt"

func getStr() string {
	return "iswbm"
}
func main() {
	fmt.Println(&getStr())
	// cannot take the address of getStr()
}
