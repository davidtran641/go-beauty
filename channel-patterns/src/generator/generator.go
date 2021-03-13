package generator

import "time"

// NewOddNumbers return a channel of odd number less than maximum value
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
