package basic

import "fmt"

//数组是值类型的
func printArray(arr [5]int) {
	//arr会进行值的拷贝
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}
func printArray2(arr *[5]int) {
	//arr会进行值的拷贝
	//指针的使用非常灵活
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}
func main() {
	//不同长度是不同的类型
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5} //需要初始化
	//让编译器来数数组的长度
	arr3 := [...]int{2, 3, 4, 6, 8}

	fmt.Println(arr1, arr2, arr3)

	//二维数组
	var grid [2][3]int
	fmt.Println(grid)
	//遍历数组
	for i := 0; i < len(arr3); i++ {
		fmt.Print(arr3[i])
	}
	//range遍历,可同时获得下边和元素的值
	/**
	意义明确、美观
	java 只能for each value
	*/
	for _, v := range arr3 {
		fmt.Print(v)
	}
	printArray(arr1)
	printArray(arr3)
	fmt.Println(arr1, arr3)
	fmt.Println("============使用数组指针的方法================")
	printArray2(&arr1)
	printArray2(&arr3)
	fmt.Println(arr1, arr3)
}
