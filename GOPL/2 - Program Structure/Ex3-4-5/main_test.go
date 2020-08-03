package main

import "testing"

var x uint64 = 965842371

func BenchmarkCountEx3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countEx23(x)
	}
}

func BenchmarkCountEx4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countEx24(x)
	}
}

func BenchmarkCountEx5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countEx25(x)
	}
}
