package main

//channel的机制是先进先出，如果你给channel赋值了，那么必须要读取它的值，不然就会造成阻塞，当然这个只对无缓冲的channel有效。
// 对于有缓冲的channel，发送方会一直阻塞直到数据被拷贝到缓冲区；
// 如果缓冲区已满，则发送方只能在接收方取走数据后才能从阻塞状态恢复
import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int)
	go produce(ch)
	go consumer(ch)
	time.Sleep(1 * time.Second)
}

func produce(p chan<- int) {
	for i := 0; i < 10; i++ {
		p <- i
		fmt.Println("send:", i)
	}
}

func consumer(c <-chan int) {
	for i := 0; i < 10; i++ {
		v := <-c
		fmt.Println("receive:", v)
	}
}

//
//我们创建了一个无缓冲的channel，然后给这个channel赋值了，程序就是在赋值完成后陷入了死锁。因为我们的channel是无缓冲的，即同步的，赋值完成后来不及读取channel，
//程序就已经阻塞了。这里介绍一个非常重要的概念：channel的机制是先进先出，如果你给channel赋值了，那么必须要读取它的值，不然就会造成阻塞，当然这个只对无缓冲的channel有效。
//对于有缓冲的channel，发送方会一直阻塞直到数据被拷贝到缓冲区；如果缓冲区已满，则发送方只能在接收方取走数据后才能从阻塞状态恢复
