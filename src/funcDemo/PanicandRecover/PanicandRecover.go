package main

import "fmt"

func main() {
	A()
	B()
	C()
}

func A()  {
	fmt.Println("Func A")
}

func B()  {
	defer func() {
		if err:=recover();err!=nil{
			fmt.Println("recover in B")
		}
	}()
	panic("panic in B")
}

func C()  {
	fmt.Println("Func C")
}

