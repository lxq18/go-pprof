package main

import (
	"flag"
	"fmt"
	"log"

	// 	"log"
	"os"
	"runtime/pprof"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()
	// 如果命令行设置了 cpuprofile
	if *cpuprofile != "" {
		// 根据命令行指定文件名创建 profile 文件
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		// 开启 CPU profiling
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	for i := 0; i < 45; i++ {
		nums := fibonacci(i)
		fmt.Println(nums)
	}
}

//递归实现的斐波纳契数列
func fibonacci(num int) int {
	if num == 0 {
		return 0
	}
	if num == 1 {
		return 1
	}
	return fibonacci(num-1) + fibonacci(num-2)
}
