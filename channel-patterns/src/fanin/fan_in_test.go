package fanin_test

import (
	"sort"
	"testing"

	"github.com/davidtran641/go-beauty/src/fanin"
	"gopkg.in/go-playground/assert.v1"
)

func TestFanIn(t *testing.T) {
	want := []string{"image-1", "image-2", "text-1", "text-2"}
	got := []string{}

	a := make(chan string, 2)
	b := make(chan string, 2)
	a <- "image-1"
	a <- "image-2"
	b <- "text-1"
	b <- "text-2"

	c := fanin.FanIn(a, b)
	for i := 0; i < 4; i++ {
		got = append(got, <-c)
	}

	sort.Strings(got)
	assert.Equal(t, want, got)
}
