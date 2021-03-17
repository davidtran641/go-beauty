package search

import (
	"log"
	"time"
)

// Result is a search result
type Result struct {
	Title string
}

// Search search by query and return results
type Search func(query string) []Result

// All returns all search results
func All(query string, timeout time.Duration, web, image, video Search) []Result {

	c := make(chan []Result)

	go func() { c <- web(query) }()
	go func() { c <- image(query) }()
	go func() { c <- video(query) }()

	results := []Result{}
	timeoutChan := time.After(timeout)
	for i := 0; i < 3; i++ {
		select {
		case r := <-c:
			results = append(results, r...)
		case <-timeoutChan:
			return results
		}
	}
	return results
}

// first returns a search function which got the first result
func first(replicas ...Search) Search {
	return func(query string) []Result {
		c := make(chan []Result, len(replicas))

		searchReplica := func(i int) {
			log.Printf("Run %d", i)
			c <- replicas[i](query)
			log.Printf("Done %d", i)
		}
		for i := range replicas {
			go searchReplica(i)
		}

		results := <-c

		return results
	}

}
