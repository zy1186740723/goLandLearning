package main

import (
	"fmt"
	"runtime"
	"sync"
)
//线程共享的变量
var counter int = 0

func Count(lock *sync.Mutex) {
	lock.Lock()
	counter++
	fmt.Println("counter =", counter)
	lock.Unlock()
}


func main() {

	lock := &sync.Mutex{}

	for i:=0; i<10; i++ {
		go Count(lock)
	}

	for {
		lock.Lock()

		c := counter

		lock.Unlock()

		runtime.Gosched()// 出让时间片

		if c >= 10 {
			break
		}
	}
}
