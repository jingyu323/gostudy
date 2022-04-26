package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	lock := new(sync.RWMutex)
	lock.Lock()
	for i := 0; i < 4; i++ {
		go func(i int) {
			fmt.Printf("第 %d 个协程准备开始... \n", i)
			lock.RLock()
			fmt.Printf("第 %d 个协程获得读锁, sleep 1s 后，释放锁\n", i)
			time.Sleep(time.Second)
			lock.RUnlock()
		}(i)
	}
	time.Sleep(time.Second * 2)

	fmt.Println("准备释放写锁，读锁不再阻塞")
	// 写锁一释放，读锁就自由了
	lock.Unlock()
	// 由于会等到读锁全部释放，才能获得写锁
	// 因为这里一定会在上面 4 个协程全部完成才能往下走
	lock.Lock()
	fmt.Println("程序退出...")
	lock.Unlock()
}
