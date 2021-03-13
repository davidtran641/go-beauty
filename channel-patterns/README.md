# Channel patterns in go
Channel is a type in go to help goroutines to synchornize without explicit lock or condition variables. In this article, we will go through some *design pattern* we can do with channel

## Basic about channel
To declare channel
```
ch := make(chan int)
```

To send data to channel
```
ch <- v
```

To receive data from channel
```
v := <- ch
```

## Channel patterns
### Channel generator
Channels are first-class values, just like other types in go. A generator is a function that retuns a channel
```
func NewOddNumbers(max int) <-chan int {
	c := make(chan int)
	go func() {
		for i := 1; i < max; i += 2 {
			c <- i
			time.Sleep(1 * time.Millisecond)
		}
		close(c)
	}()
	return c
}
```
[source](./src/generator)

### Fanin (multiplexing)
Two or more go routines and send to the same go channels, the result will be merged automatically. As a result, we have a fan-in solution.

This patterns can be applied to dispatch the work to multipler workers concurrently, then merge the final result. Such as, search for both texts and images, then merge into a final search result.

```
func FanIn(a, b <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-a
		}
	}()
	
  go func() {
		for {
			c <- <-b
		}
	}()
	
  return c
}
```
[source](./src/fanin)

### Select
When there are multiple channels, we can select which channel to receive.
This feature can be used for fan-in to merge data from multiple channels
```
func FanInSelect(a, b <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-a:
				c <- s
			case s := <-b:
				c <- s
			}
		}
	}()

	return c
}
```
[source](./src/faninselect)

### Timeout using select
When receiving a message from channe, the current goroutine will be blocked until a new message will come.

What happens if the sender takes too long to process data? Go library provides a timeout channel to help us deal with this situation.
```
func TimeoutSelect(input chan int) <-chan int {
	c := make(chan int)

	timeout := time.After(3 * time.Millisecond)
	go func() {
		for {
			select {
			case s := <-input:
				c <- s
			case <-timeout:
				fmt.Println("timeout")
				close(c)
				return
			}
		}
	}()
	return c
}
```
[source](./src/timeoutselect)
### Quit channel
A quit channel can be used to tell the subsequence goroutine to stop the work and clean up resource.

It is useful when the job is cancelled and we want to tell the child goroutines to stop their works

```
// NewFibonacci returns a channel for fibonacy numbers 
// until it's recevied quit message from quit channel
func NewFibonacci(quit chan struct{}) <-chan int {
	c := make(chan int)
	go func() {
		curr := 1
		pre := 0

		for {
			select {
			case c <- curr:
				pre, curr = curr, curr+pre
			case <-quit:
				return
			}
		}

	}()

	return c
}
```

### Daisy-chain of channels
Channels can be chained together and create a very sequence, eg: c1 => c2 => c3,...

In this bellow example, it will create a chain of 10,000 channels:
```
func chain(a, b chan int) {
	a <- 1 + <-b
}

// NewChain return a chain of channel
func NewChain() <-chan int {
	n := 10000
	result := make(chan int)

	a := result
	b := result

	for i := 0; i < n; i++ {
		b = make(chan int)
		go chain(a, b)
		a = b
	}

	go func(ch chan int) {
		ch <- 1
	}(b)

	return result
}
```
[source](./src/daisychain)

### Channel on a channel
A channel can be send on a channel as well. This can be used for back-ward communication via channel
```
type Message struct {
  result string
  wait chan bool
}

func NewMessager() chan Message {
  return make(chan Message)
}
```


