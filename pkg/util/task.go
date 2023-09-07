package util

import (
	"math/rand"
	"time"
)

// Fibonacci computes the nth number in the Fibonacci sequence. O(2^n)
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

// FibonacciOptimized computes the nth number in the Fibonacci sequence. O(n)
func FibonacciOptimized(n uint64) uint64 {
	if n <= 1 {
		return n
	}
	a, b := uint64(0), uint64(1)
	for i := uint64(2); i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

// GenerateData generates a slice with size elements of random data.
func GenerateData(size int) []int {
	rand.Seed(time.Now().UnixNano())
	data := make([]int, size)
	for i := range data {
		data[i] = rand.Intn(size)
	}
	return data
}
