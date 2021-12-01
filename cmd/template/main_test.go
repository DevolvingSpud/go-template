package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMemorySlice(t *testing.T) {
	size := 65535
	expected := (size + 1) / 2
	start := initSinglyLinkedList(size)
	assert.Equal(t, FindMiddleWithSlice(start), expected, "middle should be half of size (rounded up)")
}

func BenchmarkMemorySlice(b *testing.B) {
	b.ReportAllocs()
	start := initSinglyLinkedList(65535)
	for i := 0; i < b.N; i++ {
		FindMiddleWithSlice(start)
	}
}

func TestMemoryPointer(t *testing.T) {
	size := 65535
	expected := (size + 1) / 2
	start := initSinglyLinkedList(size)
	assert.Equal(t, FindMiddleWithPointer(start), expected, "middle should be half of size (rounded up)")
}

func BenchmarkMemoryPointer(b *testing.B) {
	b.ReportAllocs()
	start := initSinglyLinkedList(65535)
	for i := 0; i < b.N; i++ {
		FindMiddleWithPointer(start)
	}
}
