package main

import "testing"

func BenchmarkMemorySlice(b *testing.B) {
	b.ReportAllocs()
	start := initSinglyLinkedList(65535)
	for i := 0; i < b.N; i++ {
		FindMiddleWithSlice(start)
	}
}

func BenchmarkMemoryPointer(b *testing.B) {
	b.ReportAllocs()
	start := initSinglyLinkedList(65535)
	for i := 0; i < b.N; i++ {
		FindMiddleWithPointer(start)
	}
}
