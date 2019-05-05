package main

import (
	"LearnGo/retriever/mock"
	"LearnGo/retriever/real"
	"fmt"
	"time"
)

/**
 * @Author: zhangyan
 * @Date: 2019/5/5 21:08
 * @Version 1.0
 */
//接口变量里有什么 实现着的类型与实现着的值
// （或者是实现者的指针）
type Retriever interface {
	Get(url string) string
}
type Poster interface {
	Post(url string,
		form map[string]string) string
}

func download(r Retriever) string {
	return r.Get("http://www.baidu.com")
}

func post(poster Poster) {
	poster.Post("http://www.baidu.com",
		map[string]string{
			"name":   "ccmouse",
			"course": "go",
		})
}

//接口的组合
type RetrieverPoster interface {
	Retriever
	Poster
}

const url = "http://www.baidu.com"

func Session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contens": "fake",
	})
	return s.Get(url)

}

func inspect(r Retriever) {
	fmt.Println("Inspecting ", r)
	fmt.Printf(">%T %v", r, r)
	switch v := r.(type) {
	//使用指针接收者的方式
	case *mock.Retriever:
		fmt.Println("Cotents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
}
func main() {
	var r Retriever
	r = &mock.Retriever{"this is fake website"}
	fmt.Printf("%T %v\n", r, r)
	inspect(r)
	//接口中有两个 一个是类型一个是值
	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	fmt.Printf("%T %v\n", r, r)
	inspect(r)
	//type assertion
	realRetriever := r.(*real.Retriever)
	fmt.Println(realRetriever.TimeOut)
	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("not a mockRetriever.")
	}
	fmt.Println("Try a session")
	s := &mock.Retriever{"fake too"}
	fmt.Println(Session(s))

}

//fmt.Println(download(r))
