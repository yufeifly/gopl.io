package main

import "testing"

func BenchmarkPrintArgs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PrintArgs()
	}
}
