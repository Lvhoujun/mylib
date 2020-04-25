package mylib

// @doc 返回斐波拉契函数第N个
func Fibonacci(n int) int64 {
	var f = Fibonacci_imp()
	var value int64
	for i := 0; i <= n; i++ {
		value = f()
	}
	return value
}

// ====================================================================
// Internal functions
// ====================================================================
func Fibonacci_imp() func() int64 {
	var a, b int64
	a = -1
	b = 1
	return func() int64 {
		a, b = b, a+b
		return b
	}
}