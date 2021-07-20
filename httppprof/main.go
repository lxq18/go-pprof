package main

import (
	"go-pprof/c"
	"go-pprof/f"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	// 性能分析
	go func() {
		log.Println(http.ListenAndServe(":8080", nil))
	}()

	// 实际业务代码
	for {
		f.Add("test")
	}

	c.CalGoroutine("goroutineTest")
}
