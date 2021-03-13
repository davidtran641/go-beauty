package timeoutselect

import (
	"fmt"
	"time"
)

// TimeoutSelect return messages from input channel with timeout
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
