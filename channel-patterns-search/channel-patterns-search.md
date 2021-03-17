# Channel patterns applying

Let's try to apply channel patterns into a common fanout problem.

We're building a search service. This service will call mutiple services, and then merge the results at the end.

Search in this example is just a function, receive a query and return a list of result
```
type Result struct {
	Title string
}

// Search search by query and return results
type Search func(query string) []Result
```

We will build function `search.All` which call 3 other Search functions, and merge the results.
```
// All returns all search results
func All(query string, timeout time.Duration, web, image, video Search) []Result {
```

Use go-routine and channel to run all Search concurrently

```
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
```

Notice at the select statement, when time out happens, we just return the results and skip waiting other searchs.
```
select {
case <-timeoutChan:
  return results
}
```

## Replicas
To speed up searching, we can also call the search API to multiple replicas, and just select the first result. So the latency = min(replicas' latency)

```
func first(replicas ...Search) Search {
	return func(query string) []Result {
		c := make(chan []Result, len(replicas))

		searchReplica := func(i int) {
			c <- replicas[i](query)
		}
		
    for i := range replicas {
			go searchReplica(i)
		}

		results := <-c

		return results
	}

}
```

Source: [search.go](./src/search/search.go)

More example about how to use this can be found at [search_test.go](./src/search/search_test.go)

## What next?
- How to handle error, cancelation
