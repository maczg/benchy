package standalone

import (
	"fmt"
	"github.com/massimo-gollo/benchy/pkg/mixin"
	"net/http"
	"strconv"
)

func CpuIntensiveHandler(w http.ResponseWriter, r *http.Request) {
	n, _ := strconv.Atoi(r.URL.Query().Get("n"))
	result := mixin.Fibonacci(n)
	_, _ = fmt.Fprintf(w, "Fibonacci sequence of %d: %d\n", n, result)
}

func MemoryIntensiveHandler(w http.ResponseWriter, r *http.Request) {
	size, _ := strconv.Atoi(r.URL.Query().Get("size"))
	data := mixin.GenerateData(size)
	_, _ = fmt.Fprintf(w, "Generated %d random numbers.\n", len(data))
}
