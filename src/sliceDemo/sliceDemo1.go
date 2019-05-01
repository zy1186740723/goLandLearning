package main

import "fmt"
/**
 cap的用法！！
 */
func main() {
	var numbers4 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//规定了cap的长度到索引为8
	slice5 := numbers4[4:6:8]
	//cap(slice6)一直到数组的结尾
	slice6:=numbers4[4:6]
	fmt.Println(slice5)
	fmt.Println(slice6)
	fmt.Println(cap(slice5))
	fmt.Println(cap(slice6))
	fmt.Printf("%x, %x\n",  len(slice5),  cap(slice5))
	//将切片的长度和容量调整到了一个数值
	slice5 = slice5[:cap(slice5)]
	fmt.Println(slice5)
	slice5 = append(slice5, 11, 12, 13)
	fmt.Printf("%v\n",  len(slice5))
	slice7 := []int{0, 0, 0}
	copy(slice5, slice7)
	fmt.Println(slice5)
	fmt.Printf("%v, %v, %v\n",  slice5[2],  slice5[3],  slice5[4])
}
