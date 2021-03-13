package generator

import (
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestGenerators(t *testing.T) {
	ch := NewOddNumbers(5)

	want := []int{1, 3}
	got := []int{}
	for v := range ch {
		got = append(got, v)
	}
	assert.Equal(t, want, got)
}
