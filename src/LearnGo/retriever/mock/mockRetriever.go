package mock

import "fmt"

/**
 * @Author: zhangyan
 * @Date: 2019/5/5 21:11
 * @Version 1.0
 */
//值接收者
type Retriever struct {
	Contents string
}

func (r *Retriever) String() string {
	return fmt.Sprintf(
		"Retriever:{Contens=%s}", r.Contents)
}

func (r *Retriever) Post(url string,
	form map[string]string) string {
	r.Contents = form["contens"]
	return "ok"
}

func (r *Retriever) Get(url string) string {
	return r.Contents
}
