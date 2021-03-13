package quitchannel

// NewFibonacci returns a channel for fibonacy numbers until it's recevied quit message from quit channel
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
