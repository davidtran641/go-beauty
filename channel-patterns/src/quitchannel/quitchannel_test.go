package quitchannel

import (
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestQuitChannel(t *testing.T) {
	want := []int{1, 1, 2, 3, 5, 8, 13, 21}
	got := []int{}

	quit := make(chan struct{})
	c := NewFibonacci(quit)
	for i := 0; i < 8; i++ {
		got = append(got, <-c)
	}
	quit <- struct{}{}

	assert.Equal(t, got, want)
}
