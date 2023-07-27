package server

import (
	"fmt"
	"github.com/massimo-gollo/benchy/utils"
	"net/http"
	"strconv"
)

func (srv *Server) CpuIntensiveHandler(w http.ResponseWriter, r *http.Request) {
	n, _ := strconv.Atoi(r.URL.Query().Get("n"))
	result := utils.Fibonacci(n)
	_, _ = fmt.Fprintf(w, "Fibonacci sequence of %d: %d\n", n, result)
}

func (srv *Server) MemoryIntensiveHandler(w http.ResponseWriter, r *http.Request) {
	size, _ := strconv.Atoi(r.URL.Query().Get("size"))
	data := utils.GenerateData(size)
	_, _ = fmt.Fprintf(w, "Generated %d random numbers.\n", len(data))
}
