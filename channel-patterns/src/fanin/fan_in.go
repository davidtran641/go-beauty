package fanin

// FanIn return a channel from multiple channels
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
