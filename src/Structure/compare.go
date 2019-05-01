package main

import "fmt"

type human struct {
	Sex int
}
type student struct {
	human
	Name string
	Age int
}
type teacher struct {
	human
	Name string
	Age int
}
type person4 struct {
	Name string
	Age int
}
func main() {
	f:=person4{
		Name:"liu",
		Age: 12,
	}
	a:=person4{
		Name:"liu",
		Age: 12,
	}
	fmt.Println(f==a)

	c:=teacher{Name:"zhang",Age:10,human:human{Sex:0}}
	d:=student{Name:"zhang",Age:10,human:human{Sex:0}}
	c.human.Sex=1
	fmt.Println(c,d)

}

