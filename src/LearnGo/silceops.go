package main

import "fmt"

/**
 * @Author: zhangyan
 * @Date: 2019/5/4 17:57
 * @Version 1.0
 */
func printSilce(s []int) {
	fmt.Printf("len=%d,cap=%d,value=%v\n", len(s), cap(s), s)

}
func main() {
	fmt.Println("Creating slice....................")
	var s []int //zero value
	for i := 0; i < 100; i++ {
		printSilce(s)
		s = append(s, 2*i+1)
	}
	//cap乘以2进行扩容
	fmt.Println(s)

	s1 := []int{2, 3, 4, 5}
	printSilce(s1)

	s2 := make([]int, 16)
	s3 := make([]int, 10, 32)
	printSilce(s2)
	printSilce(s3)
	fmt.Println("Copying slice.........")
	copy(s2, s1)
	printSilce(s2)
	fmt.Println("Deleting elements from slice")
	s2 = append(s2[:3], s2[4:]...)
	printSilce(s2)
	fmt.Println("Poping from front")
	front := s2[0]
	s2 = s2[1:]

	fmt.Println("Poping fron tail")
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	fmt.Println(front, tail)
	printSilce(s2)

}
