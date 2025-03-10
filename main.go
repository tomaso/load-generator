package main

import (
	"errors"
	"flag"
	"fmt"
)

/*
 TODO:
 - Define and document configuration file
 - Read configuration from a file using viper
 - Find a module/lib for simple option/argument parsing
 - Implement cpu load
*/

var ErrHelp = errors.New("flag: help requested")

func main() {

	help := flag.Bool("help", false, "Show this help.")
	cpuload := flag.Int("cpuload", 20, "Basic cpu load per core that will be generated.")
	threads := flag.Int("threads", 1, "Number of threads to run with the specified load.")
	mem := flag.Int("mem", 100, "Memory to allocate (in MiB).")

	flag.Parse()
	if *help {
		flag.PrintDefaults()
		return
	}
	fmt.Printf("cpuload: %d, threads: %d, memory: %d\n", *cpuload, *threads, *mem)
}
