package real

import (
	"net/http"
	"net/http/httputil"
	"time"
)

/**
 * @Author: zhangyan
 * @Date: 2019/5/5 21:18
 * @Version 1.0
 */
type Retriever struct {
	UserAgent string
	TimeOut   time.Duration
}

//指针接收者
func (r *Retriever) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	result, err := httputil.DumpResponse(resp, true)

	resp.Body.Close()
	if err != nil {
		panic(err)
	}
	return string(result)
}
