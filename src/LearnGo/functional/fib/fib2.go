package fib

/**
 * @Author: zhangyan
 * @Date: 2019/5/6 16:15
 * @Version 1.0
 */
func Fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
