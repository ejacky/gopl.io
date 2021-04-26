package main

import "testing"

func BenchmarkEcho1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo1([]string{"zz", "cc", "dd"}, " ")
	}
}

func BenchmarkEcho2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo2([]string{"zz", "cc", "dd"}, " ")
	}
}

func BenchmarkEcho3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo3([]string{"zz", "cc", "dd"}, " ")
	}
}
