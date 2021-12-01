package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMemorySlice(t *testing.T) {
	size := 0
	start := initSinglyLinkedList(size)
	assert.Equal(t, FindMiddleWithSlice(start), -1, "middle should be -1 for nil lists")

	size = 1
	start = initSinglyLinkedList(size)
	assert.Equal(t, FindMiddleWithSlice(start), (size)/2, "middle should be half of size (rounded up)")

	size = 4
	start = initSinglyLinkedList(size)
	assert.Equal(t, FindMiddleWithSlice(start), (size)/2, "middle should be half of size (rounded up)")

	size = 5
	start = initSinglyLinkedList(size)
	assert.Equal(t, FindMiddleWithSlice(start), (size)/2, "middle should be half of size (rounded up)")
}

func BenchmarkMemorySlice(b *testing.B) {
	b.ReportAllocs()
	start := initSinglyLinkedList(65535)
	for i := 0; i < b.N; i++ {
		FindMiddleWithSlice(start)
	}
}

func TestMemoryPointer(t *testing.T) {
	size := 0
	start := initSinglyLinkedList(size)
	assert.Equal(t, FindMiddleWithPointer(start), -1, "middle should be -1 for nil lists")

	size = 1
	start = initSinglyLinkedList(size)
	assert.Equal(t, FindMiddleWithPointer(start), (size)/2, "middle should be half of size (rounded up)")

	size = 4
	start = initSinglyLinkedList(size)
	assert.Equal(t, FindMiddleWithPointer(start), (size)/2, "middle should be half of size (rounded up)")

	size = 5
	start = initSinglyLinkedList(size)
	assert.Equal(t, FindMiddleWithPointer(start), (size)/2, "middle should be half of size (rounded up)")
}

func BenchmarkMemoryPointer(b *testing.B) {
	b.ReportAllocs()
	start := initSinglyLinkedList(65535)
	for i := 0; i < b.N; i++ {
		FindMiddleWithPointer(start)
	}
}
