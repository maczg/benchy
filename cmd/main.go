package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

var port string

func init() {
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})

	port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
}

func main() {
	log.Infof("Starting server on port %s", port)
	http.HandleFunc("/health", healthCheckHandler)
	http.HandleFunc("/cpu", cpuIntensiveHandler)
	http.HandleFunc("/memory", memoryIntensiveHandler)
	addr := fmt.Sprintf(":%s", port)
	log.Fatalln(http.ListenAndServe(addr, nil))
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func cpuIntensiveHandler(w http.ResponseWriter, r *http.Request) {
	n, _ := strconv.Atoi(r.URL.Query().Get("n"))
	result := fibonacci(n)
	_, _ = fmt.Fprintf(w, "Fibonacci sequence of %d: %d\n", n, result)
}

func memoryIntensiveHandler(w http.ResponseWriter, r *http.Request) {
	size, _ := strconv.Atoi(r.URL.Query().Get("size"))
	data := generateData(size)
	_, _ = fmt.Fprintf(w, "Generated %d random numbers.\n", len(data))
}

// fibonacci computes the nth number in the Fibonacci sequence.
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// generateData generates a slice with size elements of random data.
func generateData(size int) []int {
	rand.Seed(time.Now().UnixNano())
	data := make([]int, size)
	for i := range data {
		data[i] = rand.Intn(size)
	}
	return data
}
