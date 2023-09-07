package simple

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/massimo-gollo/benchy/pkg/util"
	"strconv"
	"time"
)

func HealthHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"health": "ok",
	})
}

func CpuTaskHandler(c *gin.Context) {
	start := time.Now()
	n, _ := strconv.Atoi(c.Query("n"))
	result := util.FibonacciOptimized(uint64(n))
	elapsed := time.Since(start)
	c.JSON(200, gin.H{
		"result": fmt.Sprintf("Fibonacci sequence of %d: %d\n. Elapsed: %s", n, result, elapsed),
	})
}

func CpuIntensiveTaskHandler(c *gin.Context) {
	start := time.Now()
	n, _ := strconv.Atoi(c.Query("n"))
	result := util.Fibonacci(n)
	elapsed := time.Since(start)
	c.JSON(200, gin.H{
		"result": fmt.Sprintf("Fibonacci sequence of %d: %d\n. Elapsed: %s", n, result, elapsed),
	})
}

func MemTaskHandler(c *gin.Context) {
	size, _ := strconv.Atoi(c.Query("size"))
	data := util.GenerateData(size)
	c.JSON(200, gin.H{
		"result": fmt.Sprintf("Generated %d random numbers.\n", len(data)),
	})
}
