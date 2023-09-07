package util

import (
	"math/rand"
	"time"
)

// Fibonacci computes the nth number in the Fibonacci sequence.
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
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
