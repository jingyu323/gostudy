package main

import (
	"context"
	"fmt"
	"time"
)

func HandelRequest1(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("HandelRequest Done.")
			return
		default:
			fmt.Println("HandelRequest running, parameter: ", ctx.Value("parameter"))
			time.Sleep(2 * time.Second)
		}
	}
}

func main() {
	ctx := context.WithValue(context.Background(), "parameter", "1")
	go HandelRequest1(ctx)

	time.Sleep(10 * time.Second)
}

/**
Context仅仅是一个接口定义，根据实现的不同，可以衍生出不同的context类型；
cancelCtx实现了Context接口，通过WithCancel()创建cancelCtx实例；
timerCtx实现了Context接口，通过WithDeadline()和WithTimeout()创建timerCtx实例；
valueCtx实现了Context接口，通过WithValue()创建valueCtx实例；
三种context实例可互为父节点，从而可以组合成不同的应用形式；
*/
