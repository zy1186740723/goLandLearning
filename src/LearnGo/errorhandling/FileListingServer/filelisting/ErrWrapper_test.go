package filelisting

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

/**
 * @Author: zhangyan
 * @Date: 2019/5/7 14:30
 * @Version 1.0
 */

func errPanic(writer http.ResponseWriter,
	request *http.Request) error {
	panic(123)
}

type testingUserError string

func (e testingUserError) Error() string {
	return e.Message()
}

func (e testingUserError) Message() string {
	return string(e)
}

func errUserError(writer http.ResponseWriter,
	request *http.Request) error {
	return testingUserError("user error")
}

func errNotFound(writer http.ResponseWriter,
	request *http.Request) error {
	return os.ErrNotExist
}

func errNotPermission(writer http.ResponseWriter,
	request *http.Request) error {
	return os.ErrPermission

}

var tests = []struct {
	h       appHandler
	code    int
	message string
}{
	{errPanic, 500, "Internal Server Error"},
	{errUserError, 400, "user error"},
	{errNotFound, 404, "Not Found"},
	{errNotPermission, 403, "Forbidden"},
}

func TestErrWrapper(t *testing.T) {

	//httptest.NewRecorder()
	for _, tt := range tests {
		//测试的目标函数
		f := ErrWrapper(tt.h)
		response := httptest.NewRecorder()
		request := httptest.NewRequest(
			http.MethodGet,
			"http://www.baidu.com",
			nil)
		f(response, request)

		b, _ := ioutil.ReadAll(response.Body)
		//需要把换行给去掉
		body := strings.Trim(string(b), "\n")
		if response.Code != tt.code || body != tt.message {
			t.Errorf("expected(%d %s);"+"got"+
				"(%d %s)",
				tt.code, tt.message, response.Code, body)
		}
	}
}

func TestErrWrapperInServer(t *testing.T) {
	for _, tt := range tests {
		f := ErrWrapper(tt.h)
		server := httptest.NewServer(http.HandlerFunc(f))
		resp, _ := http.Get(server.URL)

		b, _ := ioutil.ReadAll(resp.Body)
		//需要把换行给去掉
		body := strings.Trim(string(b), "\n")
		if resp.StatusCode != tt.code || body != tt.message {
			t.Errorf("expected(%d %s);"+"got"+
				"(%d %s)",
				tt.code, tt.message, resp.StatusCode, body)
		}
	}
}
