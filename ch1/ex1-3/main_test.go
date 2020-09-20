package main

import "testing"

func BenchmarkPrint1(b *testing.B) {
	strvec := []string{"i", "am", "yufeixiong"}
	for i := 0; i < b.N; i++ {
		Print1(strvec)
	}
}

func BenchmarkPrint2(b *testing.B) {
	strvec := []string{"i", "am", "yufeixiong"}
	for i := 0; i < b.N; i++ {
		Print2(strvec)
	}
}
