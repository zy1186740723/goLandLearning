package main

import "fmt"

type TZ1 int

func main() {
	var a TZ1
	a=0
	a.Increase()
	fmt.Println(a)


}

func (a *TZ1) Increase()  {
	for i:=0;i<100 ;i++  {
		*a+=1
	}
}
