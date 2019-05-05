package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
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
	PrintFileContents(file)
}

//省略结束条件
func forever() {
	for {
		fmt.Println("abc")
	}
}
func PrintFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
func main() {

	fmt.Println(
		convertToBin(5), //101
		convertToBin(13),
		convertToBin(100),
		convertToBin(0)) // 1101
	printFile("abc.txt")
	s := `abc"d"
		pkkkk`
	PrintFileContents(strings.NewReader(s))

	//forever() //死循环
}
