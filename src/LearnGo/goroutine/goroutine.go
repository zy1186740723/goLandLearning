package main

import (
	"fmt"
	"time"
)

/**Coroutine携程是轻量级的线程
1、非抢占式多任务处理，由协程主动交出控制权，所以轻量级
2、编译器 解释器 虚拟机层面的多任务
3、轻量级的线程
4、多个携程可以在一个或者多个线程上运行
 * @Author: zhangyan
 * @Date: 2019/5/7 15:39
 * @Version 1.0
*/
/*子程序是协程的特例
协程main与dowork之间有一个双向的通道，统治权可以相互交换

主要的切换点：
i/o slelct
channel
等待锁
函数调用
runtime.Gosched()

*/

func main() {
	//var a [10]int
	for i := 0; i < 1000; i++ {
		go func(i int) {
			for {
				//无法交出控制权
				//a[i]++
				//手动交出控制权
				//runtime.Gosched()
				fmt.Printf("hellow goruntine %d\n", i)
			}
		}(i) //不传入i的话 ，跳出的时候i会等于10跳出
	}
	time.Sleep(time.Millisecond)
	//fmt.Println(a)
}
