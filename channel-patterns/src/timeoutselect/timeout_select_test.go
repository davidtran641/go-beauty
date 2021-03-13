package timeoutselect

import (
	"testing"
	"time"

	"gopkg.in/go-playground/assert.v1"
)

func TestTimeoutSelect(t *testing.T) {
	input := make(chan int, 3)
	go func() {
		for i := 0; i < 5; i++ {
			input <- (i + 1)
			time.Sleep(2 * time.Millisecond)
		}

	}()

	c := TimeoutSelect(input)

	want := []int{1, 2}
	got := []int{}
	for v := range c {
		got = append(got, v)
	}
	assert.Equal(t, want, got)
}
