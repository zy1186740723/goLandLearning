package main

import (
	"fmt"
	"sync"
)

/**
 * @Author: zhangyan
 * @Date: 2019/5/7 21:31
 * @Version 1.0
 */

//worker中的done做什么由createWorker来配置
//函数返回一个worker
func creatework(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doworker(id, w.in, w)
	return w

}
func doworker(id int, c chan int, w worker) {
	for n := range w.in {
		//n,ok:=<-c
		//if !ok{
		//	break
		//}
		fmt.Printf("worker %d received %c\n",
			id, n)
		w.done()

	}

}

//用worker将两个chan包装 传出来
type worker struct {
	in   chan int
	done func()
}

func channelDemo() {
	var wg sync.WaitGroup

	//var c chan int//c==nil
	//只能给channel发数据，不能从channel收数据
	var workers [10]worker
	for i := 0; i < 10; i++ {
		//channels[i]=make(chan int)
		//go worker(i,channels[i])
		workers[i] = creatework(i, &wg)
	}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		workers[i].in <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		workers[i].in <- 'A' + i
	}
	wg.Wait()
	//wait for all of them,等待多人完成任务
	//for _, worker := range workers {
	//	<-worker.done
	//	<-worker.done
	//}
}
func main() {
	channelDemo()
}
