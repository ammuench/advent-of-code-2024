package main

import (
	"os"
	"testing"
)

func BenchmarkPart1(b *testing.B) {
	os.Stdout, _ = os.OpenFile(os.DevNull, os.O_WRONLY, 0)
	// for i := 0; i < b.N; i++ {
		part1()
	// }
}

func BenchmarkPart2(b *testing.B) {
	os.Stdout, _ = os.OpenFile(os.DevNull, os.O_WRONLY, 0)
	// for i := 0; i < b.N; i++ {
		part2()
	// }
}

