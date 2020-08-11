package main

import (
	"fmt"
	"testing"
)

func Benchmark(b *testing.B) {
	for i := 0; i < 2000; i++ {
		fmt.Println("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	}
}
