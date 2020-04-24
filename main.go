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
	urls := ListParams()

	// Limit parallel requests
	/* This examples limits the parallel concurrnt requests using a buffered channel
	whn the buffer is full, it blocks go routine, till it releases.
	*/
	/* We can also control number of parallel requests by dividing url array into n of arrays
	, n is the parallel requests and each array run in seperate go routine.
	*/
	var maxGoroutines int
	if ParallelValue != 0 {
		maxGoroutines = ParallelValue
	} else {
		maxGoroutines = len(urls)
	}
	limitChan := make(chan struct{}, maxGoroutines)

	for _, url := range urls {
		wg.Add(1)
		go RunOne(url, &wg, &limitChan)
	}
	wg.Wait()
}

func RunOne(url string, wg *sync.WaitGroup, limitChan *chan struct{}) {
	defer wg.Done()
	*limitChan <- struct{}{}
	hash := FetchAndHash(url)
	fmt.Printf("%v %v\n", url, hash)
	<-*limitChan
}

func FetchAndHash(url string) string {
	body := CallHTTP(url)
	return HashContent(body)
}
