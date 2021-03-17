package search

import (
	"sort"
	"testing"
	"time"

	"gopkg.in/go-playground/assert.v1"
)

func mockSearch(results []Result, timeout time.Duration) Search {
	return func(query string) []Result {
		time.Sleep(timeout)
		return results
	}
}

var (
	mockWeb   = mockSearch([]Result{{"web-1"}}, time.Millisecond*2)
	mockImage = mockSearch([]Result{{"image-1"}}, time.Millisecond*2)
	mockVideo = mockSearch([]Result{{"video-1"}}, time.Millisecond*2)
)

func TestSearch(t *testing.T) {
	want := []Result{
		{"image-1"},
		{"video-1"},
		{"web-1"},
	}

	c := make(chan []Result)

	go func() {
		c <- All("a", time.Millisecond*3, mockWeb, mockImage, mockVideo)
	}()

	var got []Result
	select {
	case got = <-c:
		break
	case <-time.After(3 * time.Millisecond):
		break
	}

	sort.Slice(got, func(i, j int) bool {
		return got[i].Title < got[j].Title
	})

	assert.Equal(t, want, got)
}

func TestSearchTimeout(t *testing.T) {
	want := []Result{}

	c := make(chan []Result)

	go func() {
		c <- All("a", time.Millisecond*1, mockWeb, mockImage, mockVideo)
	}()

	var got []Result
	select {
	case got = <-c:
		break
	case <-time.After(3 * time.Millisecond):
		break
	}

	sort.Slice(got, func(i, j int) bool {
		return got[i].Title < got[j].Title
	})

	assert.Equal(t, want, got)
}

func TestSearchReplica(t *testing.T) {

	want := []Result{
		{"image-2"},
		{"video-1"},
		{"web-1"},
	}

	web1 := mockSearch([]Result{{"web-1"}}, time.Millisecond*1)
	web2 := mockSearch([]Result{{"web-2"}}, time.Millisecond*5)

	image1 := mockSearch([]Result{{"image-1"}}, time.Millisecond*4)
	image2 := mockSearch([]Result{{"image-2"}}, time.Millisecond*2)

	video1 := mockSearch([]Result{{"video-1"}}, time.Millisecond*1)
	video2 := mockSearch([]Result{{"video-2"}}, time.Millisecond*6)

	c := make(chan []Result)

	go func() {
		web := first(web1, web2)
		image := first(image1, image2)
		video := first(video1, video2)
		c <- All("a", time.Millisecond*4, web, image, video)
	}()

	var got []Result
	select {
	case got = <-c:
	case <-time.After(3 * time.Millisecond):
	}

	sort.Slice(got, func(i, j int) bool {
		return got[i].Title < got[j].Title
	})

	assert.Equal(t, want, got)

}
