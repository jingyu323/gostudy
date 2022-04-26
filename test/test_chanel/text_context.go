package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	test1_chan()
	test_context()
	test_gowithtimeout()

}

func test1_chan() {
	stop := make(chan bool)

	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("监控退出，停止了...")
				return
			default:
				fmt.Println("goroutine监控中...")
				time.Sleep(2 * time.Second)
			}
		}
	}()
	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")
	stop <- true
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)
}

func monitor(ctx context.Context, number int) {

	for {
		select {
		// 其实可以写成 case <- ctx.Done()
		// 这里仅是为了让你看到 Done 返回的内容
		case v := <-ctx.Done():
			fmt.Printf("监控器%v，接收到通道值为：%v，监控结束。\n", number, v)
			return
		default:
			fmt.Printf("监控器%v，正在监控中...\n", number)
			time.Sleep(2 * time.Second)
		}

	}

}
func test_context() {
	ctx, cancel := context.WithCancel(context.Background())
	for i := 1; i <= 5; i++ {
		go monitor(ctx, i)
	}

	time.Sleep(1 * time.Second)
	// 关闭所有 goroutine
	cancel()
	// 等待5s，若此时屏幕没有输出 <正在监控中> 就说明所有的goroutine都已经关闭
	time.Sleep(5 * time.Second)
	fmt.Println("主程序退出！！")

}

func test_gowithtimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	ctx02, cancel := context.WithTimeout(ctx, 1*time.Second)

	defer cancel()

	for i := 1; i <= 5; i++ {
		go monitor(ctx02, i)
	}

	time.Sleep(5 * time.Second)
	if ctx02.Err() != nil {
		fmt.Println("监控器取消的原因: ", ctx02.Err())
	}

	fmt.Println("主程序退出！！")

}
