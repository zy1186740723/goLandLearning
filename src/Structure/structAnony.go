package main

import "fmt"

type Person struct {
	Name string
	Age int
}
//匿名结构
type person2 struct {
	Name string
	Age int
	contact struct{
		Phone string
		City string
	}
}
//匿名字段
type person3 struct {
	string
	int
}
func main() {
	//匿名结构
	a:= &struct {
		Name string
		Age int
	}{
		Name:"c",
		Age:12,
	}
	fmt.Println(a)

	b:=person2{Name:"zhang",Age:19}
	b.contact.Phone="1234456"
	b.contact.City="beijing"
	fmt.Println(b)
	c:=person3{"zhang",19}//必须严格顺序进行赋值
	fmt.Println(c)

	f:=Person{
		Name:"liu",
		Age: 12,
	}
	//类型可以传递
	var d Person
	d=f
	fmt.Println(d)


}
