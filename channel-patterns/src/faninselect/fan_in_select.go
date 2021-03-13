package faninselect

// FanInSelect return channel from multiple channels using select
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
