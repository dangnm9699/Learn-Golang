package main

import "testing"

var args = []string{"nguyen", "minh", "dang"}

func BenchmarkEcho1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo1(args)
	}
}

func BenchmarkEcho2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo2(args)
	}
}
