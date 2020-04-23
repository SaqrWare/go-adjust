package main

import (
	"fmt"
	"sync"
)

var ParallelValue int

func main() {
	// Waiting group
	var wg sync.WaitGroup

	// Parse flag
	ParallelValue = ParseParallelFlag()

	// Get urls from cli args
	urls := FixURLs(ListParams())

	// Limit parallel requests
	var maxGoroutines int
	if ParallelValue != 0 {
		maxGoroutines = ParallelValue
	} else {
		maxGoroutines = len(urls)
	}
	limitChan := make(chan struct{}, maxGoroutines)

	for _, url := range urls {
		wg.Add(1)
		go runOne(url, &wg, &limitChan)
	}
	wg.Wait()
}

func runOne(url string, wg *sync.WaitGroup, limitChan *chan struct{}) {
	defer wg.Done()
	*limitChan <- struct{}{}
	body := CallHTTP(url)
	url, hash := HashResponse(url, body)
	fmt.Println(url + " " + hash)
	<-*limitChan
}
