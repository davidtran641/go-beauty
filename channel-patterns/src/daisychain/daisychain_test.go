package daisychain

import (
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestChain(t *testing.T) {
	want := 10001

	c := NewChain()
	got := <-c

	assert.Equal(t, want, got)

}
