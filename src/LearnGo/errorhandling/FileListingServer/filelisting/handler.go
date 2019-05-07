package filelisting

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

/**服务器统一的出错处理1
**defer+panic+recover的方式综合处理
 * @Author: zhangyan
 * @Date: 2019/5/6 16:46
 * @Version 1.0
*/
//出错处理的逻辑 ,使用了别名
type appHandler func(writer http.ResponseWriter,
	request *http.Request) error

//函数式编程，将一个函数包装为了另一个函数
func ErrWrapper(handler appHandler) func(writer http.ResponseWriter,
	request *http.Request) {
	return func(writer http.ResponseWriter,
		request *http.Request) {
		//处理遇到的错误,统一出错处理之recover
		//用recover进行保护
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic: %v", r)
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		err := handler(writer, request)
		if err != nil {
			//输出：error handling,path must start with/list/
			log.Printf("error handling：%s", err.Error())

			//类型判断
			//fmt.Println(reflect.TypeOf(err))
			//fmt.Println(reflect.TypeOf(err.(UserError)))

			//user Error
			if userError, ok := err.(UserError); ok {
				http.Error(writer,
					userError.Message(), http.StatusBadRequest)

				//处理完要进行return退出
				return
			}
			//system error
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
				//http.Error(
				//	writer,
				//	http.StatusText(http.StatusNotFound),
				//	http.StatusNotFound)
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError

			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

type UserError interface {
	error            //系统看的 error()
	Message() string // 给用户看的
}

type userError string

//给系统看的
func (e userError) Error() string {
	return e.Message()
}

//给用户看的
func (e userError) Message() string {
	return string(e)
}

const prefix = "/list/"

//业务逻辑
func HandlerFileList(writer http.ResponseWriter,
	request *http.Request) error {
	if strings.Index(
		request.URL.Path, prefix) != 0 {
		//
		// return errors.New("path must start "+"with"+prefix)
		//返回的是userError类型的错误
		return userError("!!path must start " + "with " + prefix)
	}
	path := request.URL.Path[len("/list/"):] //  /list/fib.txt
	file, err := os.Open(path)
	if err != nil {
		//错误进行了包装
		//http.Error(writer,err.Error(),
		//	http.StatusInternalServerError)
		return err
	}
	defer file.Close()
	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	writer.Write(all)
	return nil
}
