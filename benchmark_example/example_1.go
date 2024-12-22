package benchmark_example

func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fib(n-2) + fib(n-1)
}

func fib2(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	first := 0
	second := 1
	for i := 2; i <= n; i++ {
		first, second = second, first+second
	}
	return second
}
