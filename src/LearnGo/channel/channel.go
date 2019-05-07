package main

import (
	"fmt"
	"time"
)

/**
 * @Author: zhangyan
 * @Date: 2019/5/7 16:12
 * @Version 1.0
 */
//外面的人发数据
func creatework2(id int) chan<- int {
	c := make(chan int)
	go func() {
		for {
			//里面的人收数据
			fmt.Printf("worker %d received %c\n",
				id, <-c)
		}
	}()
	return c

}
func worker(id int, c chan int) {

	for n := range c {
		//n,ok:=<-c
		//if !ok{
		//	break
		//}
		fmt.Printf("worker %d received %c\n",
			id, n)
	}

}
func channelDemo() {
	//var c chan int//c==nil
	//只能给channel发数据，不能从channel收数据
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		//channels[i]=make(chan int)
		//go worker(i,channels[i])
		channels[i] = creatework2(i)
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
	time.Sleep(time.Millisecond)
}

//具备缓冲区的channel
func bufferedChannel() {
	c := make(chan int, 3)
	go worker(0, c)
	//3，缓冲区就是3 3个以后就会报错
	c <- 'a' //发送数据之后一定要有人接受，不然会报错
	c <- 'b'
	c <- 'c'
	c <- 'd'
	time.Sleep(time.Millisecond)
}

func channelClose() {
	c := make(chan int, 3)
	go worker(0, c)
	//3，缓冲区就是3 3个以后就会报错
	c <- 'a' //发送数据之后一定要有人接受，不然会报错
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c) //不断收集一毫秒时间的空串
	time.Sleep(time.Millisecond)

}

func main() {
	//1、channel是一等公民，可以传递
	//channelDemo()
	//2、演示bufferedChannel
	//bufferedChannel()
	//3、关闭channel，用range来收取
	channelClose()

}
