package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

//1. M 个接受者，1 个发送者， 发送者通过关闭数据通道说 「不要再发送了」
//这是最简单的情况，只需让发送者在不想再发送数据的时候关闭数据通道：

func main() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)

	// ...
	const MaxRandomNumber = 100000
	const NumReceivers = 100

	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(NumReceivers)

	// ...
	dataCh := make(chan int, 100)

	// 发送者
	go func() {
		for {
			if value := rand.Intn(MaxRandomNumber); value == 0 {
				// 唯一的发送者可以安全地关闭通道。
				close(dataCh)
				return
			} else {
				dataCh <- value
			}
		}
	}()

	// 接收者
	for i := 0; i < NumReceivers; i++ {
		go func() {
			defer wgReceivers.Done()

			// 接收数据直到 dataCh 被关闭或者
			// dataCh 的数据缓存队列是空的。
			for value := range dataCh {
				log.Println(value)
			}
		}()
	}

	wgReceivers.Wait()

}
