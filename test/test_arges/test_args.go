package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

var interval time.Duration

func init() {
	fmt.Println("init")
	flag.DurationVar(&interval, "interval", 1*time.Second, "循环间隔")
}
func main() {
	if len(os.Args) > 0 {
		for index, arg := range os.Args {
			fmt.Printf("args[%d]=%v\n", index, arg)
		}
	}

	var name string
	flag.StringVar(&name, "name", "jack", "your name")

	flag.Parse() // 解析参数
	fmt.Println(name)

	fmt.Println(interval)

}
