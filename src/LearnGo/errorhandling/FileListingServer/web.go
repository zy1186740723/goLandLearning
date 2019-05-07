package main

import (
	"LearnGo/errorhandling/FileListingServer/filelisting"
	"net/http"
)

/**
 * @Author: zhangyan
 * @Date: 2019/5/6 16:36
 * @Version 1.0
 */

func main() {
	http.HandleFunc("/list/",
		filelisting.ErrWrapper(filelisting.HandlerFileList))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
