package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//省略初始条件
func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

//省略了递增的条件
func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

//省略结束条件
func forever() {
	for {
		fmt.Println("abc")
	}
}

func main() {

	fmt.Println(
		convertToBin(5), //101
		convertToBin(13),
		convertToBin(100),
		convertToBin(0)) // 1101
	printFile("abc.txt")
	forever() //死循环
}
