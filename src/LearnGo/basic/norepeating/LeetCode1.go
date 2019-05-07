package main

import "fmt"

/**norepeating
**寻找最长不含有重复字符的子串
 * @Author: zhangyan
 * @Date: 2019/5/4 18:24
 * @Version 1.0
*/
var lastOccurred = make([]int, 0xffff) //65535
func lengthofNoRepeatingSubStr(s string) int {
	//lastOccurred := make(map[rune]int)

	for i := range lastOccurred {
		lastOccurred[i] = -1
	}
	//lastOccurred[0x65]=1//e
	//lastOccurred[0x8BFE]=6 //课

	start := 0
	maxLength := 0
	for i, ch := range []rune(s) {
		lastI := lastOccurred[ch]
		if lastI != -1 && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength

}
func main() {
	fmt.Println(lengthofNoRepeatingSubStr("abcd"))
	fmt.Println(lengthofNoRepeatingSubStr("张言是网我"))
}
