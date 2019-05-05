package basic

import "fmt"

/**
 * @Author: zhangyan
 * @Date: 2019/5/4 18:11
 * @Version 1.0
 */
func main() {
	//hash map 是乱序的
	m := map[string]string{
		"name":    "zhangyan",
		"course":  "golang",
		"site":    "mysite",
		"quality": "good",
	}
	m2 := make(map[string]int)
	var m3 map[string]int
	fmt.Println(m, m2, m3)
	fmt.Println("===========Traversing map========")
	for k, v := range m {
		fmt.Println(k, v)
	}
	fmt.Println("=======getting values==========")
	courseName, ok := m["course"]
	fmt.Println(courseName)
	//key出错的话 就会空
	courseName1, ok := m["corse"]
	fmt.Println(courseName1, ok)
	if courseName, ok := m["course"]; ok {
		fmt.Println(courseName)
	} else {
		fmt.Println("error")
	}
	fmt.Println("deleting values")
	name, ok := m["name"]
	fmt.Println(name, ok)
	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok)
	/**
	1、map使用哈希表，必须可以比较相等
	2、除了slice map function 都可以作为key
	3、struct不包含上述字段也可作为key
	*/

}
