package basic

import "fmt"

/**
 * @Author: zhangyan
 * @Date: 2019/5/4 17:34
 * @Version 1.0
 */

func updateSlice(s []int) {
	s[0] = 100

}
func main() {
	//slice上一个视图的数据结构
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s := arr[2:6]
	fmt.Println("arr[2:6]", s)
	fmt.Println(arr[:6])
	fmt.Println(arr[2:])
	fmt.Println(arr[:]) //全部
	s1 := arr[2:]
	fmt.Println(s1)
	s2 := arr[:]
	fmt.Println(s2)
	//slice底层是没有数据的，是对底层array的一个view
	fmt.Println("updated............")
	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println(arr)
	//进行reslice的操作
	s2 = s2[:5]
	fmt.Println(s2)
	s2 = s2[2:]
	fmt.Println(s2)

	fmt.Println("Extending slice")
	/**
	slice存在cap 可以向后扩展，但是不能向前拓展
	*/
	arr[0], arr[2] = 0, 2
	fmt.Println(arr)
	s1 = arr[2:6]
	s2 = s1[3:6]
	fmt.Println(s1, s2)
	fmt.Println(len(s1), cap(s1))
	fmt.Println(len(s2), cap(s2))

	//当超越cap的时候，会开辟一个新的arr，s4 s5不要再是对arr的view了
	//arr并不会改变，会重新分配更大的低层数组
	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	fmt.Println(s3, s4, s5)
	fmt.Println(arr)

}
