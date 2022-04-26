package main

func SafeSend(ch chan bool, value bool) (closed bool) {
	defer func() {
		if recover() != nil {
			closed = true
		}
	}()

	ch <- value  // 如果ch已关闭，则产生一个恐慌。
	return false // <=> closed = false; return
}
func SafeClose(ch chan bool) (justClosed bool) {
	defer func() {
		if recover() != nil {
			// 一个函数的返回结果可以在defer调用中修改。
			justClosed = false
		}
	}()

	// 假设ch != nil。
	close(ch)   // 如果ch已关闭，则产生一个恐慌。
	return true // <=> justClosed = true; return
}

func main() {
	c := make(chan bool)
	go func() {
		for i := 0; i < 25; i++ {
			SafeSend(c, true)
		}
	}()
	SafeClose(c)
}
