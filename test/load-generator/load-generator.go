package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	var concurrentRequests int = 10
	var numRequests int64 = 1000
	var totalTime time.Duration

	for i := 0; i < concurrentRequests; i++ {
		go func() {
			requestStartTime := time.Now()
			for j := int64(0); j < numRequests; j++ {
				resp, err := http.Get("http://localhost:8080")
				if err != nil {
					log.Println(err)
				}
				defer resp.Body.Close()

				if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
					fmt.Printf("Request %d was successful\n", j+1)
				} else {
					fmt.Printf("Request %d failed with status code %d\n", j+1, resp.StatusCode)
				}
			}

			totalTime += time.Since(requestStartTime)
		}()
	}

	var wg sync.WaitGroup
	for i := 0; i < concurrentRequests; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
		}()
	}

	wg.Wait()

	fmt.Printf("Total execution time: %s\n", totalTime.String())
	fmt.Printf("Average request time: %fms\n", float64(totalTime.Milliseconds())/float64(numRequests*int64(concurrentRequests)))

	select {}
}
