package benchMark

import (
	"fmt"
	"testing"
)

func BenchmarkParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Hello()
		}
	})
}

func BenchmarkOne(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		Hello()
	}
}

func Hello() {
	fmt.Printf("hello world\n")
}