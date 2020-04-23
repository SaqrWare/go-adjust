package main

import (
	"flag"
	"os"
	"strconv"
)

func ListParams() []string {
	args := os.Args
	result := make([]string, 0)
	for i, item := range args {
		// Skip filename
		// Skip flags
		// skip flag values
		if i == 0 || string(item[0]) == "-" || item == strconv.Itoa(ParallelValue) {
			continue
		}
		result = append(result, item)
	}
	return result
}

func ParseParallelFlag() int {
	parallel := flag.Int("parallel", 0, "To set the maximum number of parallel requests.")
	flag.Parse()
	return *parallel
}
