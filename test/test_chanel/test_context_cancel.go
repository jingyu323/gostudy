package main

import (
	"context"
	"fmt"
	"time"
)

func HandelRequest(ctx context.Context) {
	go WriteRedis(ctx)
	go WriteDatabase(ctx)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("HandelRequest Done.")
			return
		default:
			fmt.Println("HandelRequest running")
			time.Sleep(2 * time.Second)
		}
	}
}

func WriteRedis(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("WriteRedis Done.")
			return
		default:
			fmt.Println("WriteRedis running")
			time.Sleep(2 * time.Second)
		}
	}
}

func WriteDatabase(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("WriteDatabase Done.")
			return
		default:
			fmt.Println("WriteDatabase running")
			time.Sleep(2 * time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go HandelRequest(ctx)

	time.Sleep(5 * time.Second)
	fmt.Println("It's time to stop all sub goroutines!")
	cancel()

	//Just for test whether sub goroutines exit or not
	time.Sleep(5 * time.Second)
}

/**
这里将自身添加到父节点的过程有必要简单说明一下：

如果父节点也支持cancel，也就是说其父节点肯定有children成员，那么把新context添加到children里即可；
如果父节点不支持cancel，就继续向上查询，直到找到一个支持cancel的节点，把新context添加到children里；
如果所有的父节点均不支持cancel，则启动一个协程等待父节点结束，然后再把当前context结束。
*/
