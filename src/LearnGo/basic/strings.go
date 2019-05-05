package basic

import (
	"fmt"
	"unicode/utf8"
)

/**
 * @Author: zhangyan
 * @Date: 2019/5/5 8:03
 * @Version 1.0
 */
func main() {
	s := "Yes我是张言!" //UTF-8 中文三个字节 英文1个字节
	fmt.Println(len(s))
	fmt.Println("", []byte(s))
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println()
	/*ch is a rune string 进行utf-8解码，
	转unicode，放在了rune4字节的类型中
	*/
	fmt.Println("decode转为unicode")
	for i, ch := range s {
		fmt.Printf("(%d %X)", i, ch)
	}
	fmt.Println()
	/**/
	fmt.Println("Rune count，获取字符的数量", utf8.RuneCountInString(s))
	//获取所有的字节
	bytes := []byte(s)
	//len获取字节的长度
	for len(bytes) > 0 {
		//返回字符和字符的size，英文是1 中文是3
		//rune类型的ch
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	//每个rune是四个字节，进行了decode
	//转成rune的方式，下标与字符相对应
	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c)", i, ch)
	}
	//Fileds Split Join
	//Trim trimRight trimleft

}
