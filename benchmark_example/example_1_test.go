// fib_test.go
package benchmark_example

import (
	"testing"
	"time"
)

func BenchmarkFib(b *testing.B) {
	time.Sleep(time.Second * 3) // 模拟耗时准备任务
	// b.ResetTimer() // 重置定时器

	for n := 0; n < b.N; n++ {
		fib(30) // run fib(30) b.N times
	}
}

func BenchmarkFib2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fib2(30) // run fib2(30) b.N times
	}
}
