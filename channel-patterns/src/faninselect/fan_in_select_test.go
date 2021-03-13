package faninselect

import (
	"sort"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestFanInSelect(t *testing.T) {
	want := []string{"image-1", "image-2", "text-1", "text-2"}
	got := []string{}

	a := make(chan string, 2)
	b := make(chan string, 2)
	a <- "image-1"
	a <- "image-2"
	b <- "text-1"
	b <- "text-2"

	c := FanInSelect(a, b)
	for i := 0; i < 4; i++ {
		got = append(got, <-c)
	}

	sort.Strings(got)
	assert.Equal(t, want, got)
}
