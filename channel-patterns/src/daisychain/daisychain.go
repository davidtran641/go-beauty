package daisychain

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
