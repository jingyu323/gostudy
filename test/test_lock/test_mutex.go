package main

import (
	"fmt"
	"sync"
)

func add(count *int, wg *sync.WaitGroup, lock *sync.Mutex) {
	for i := 0; i < 1000; i++ {
		lock.Lock()
		*count = *count + 1
		lock.Unlock()
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	count := 0
	// 修改原来的值需要给地址
	wg.Add(3)
	//lock :=  new (sync.Mutex)
	lock := &sync.Mutex{}
	go add(&count, &wg, lock)
	go add(&count, &wg, lock)
	go add(&count, &wg, lock)

	wg.Wait()
	fmt.Println("count 的值为：", count)

}
