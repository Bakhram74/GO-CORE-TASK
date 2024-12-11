package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestJoinChannels(t *testing.T) {
	a := make(chan int)
	b := make(chan int)
	c := make(chan int)

	go func() {
		for _, n := range []int{1, 2, 3} {
			a <- n
		}
		close(a)
	}()
	go func() {
		for _, n := range []int{10, 20, 30} {
			b <- n
		}
		close(b)
	}()
	go func() {
		for _, n := range []int{100, 200, 300} {
			c <- n
		}
		close(c)
	}()

	numbers := joinChannels(a, b, c)

	slice := make([]int, 0, 10)

	for n := range numbers {
		slice = append(slice, n)
	}

	expected := []int{1, 2, 3, 10, 20, 30, 100, 200, 300}

	sort.Ints(slice)

	if !reflect.DeepEqual(expected, slice) {
		t.Errorf("Expected %v, got %v", expected, slice)
	}
}
