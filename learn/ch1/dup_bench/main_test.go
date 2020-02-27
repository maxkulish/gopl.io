package main

import "testing"

const testFile = "test.txt"

func BenchmarkReadFileByLine(b *testing.B) {
	b.ReportAllocs()

	counts := make(map[string]int)

	for i := 0; i < b.N; i++ {
		ReadFileByLine(testFile, counts)
	}
}

func BenchmarkReadWholeFile(b *testing.B) {
	b.ReportAllocs()

	counts := make(map[string]int)

	for i := 0; i < b.N; i++ {
		ReadWholeFile(testFile, counts)
	}
}
