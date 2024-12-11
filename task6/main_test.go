package main

import (
	"testing"
)

func TestRandomNumber(t *testing.T) {
	n := 10
	numbers := randomNumber(n)
	slice := make([]int, 0, n)
	for n := range numbers {
		slice = append(slice, n)
	}
	if len(slice) != 10 {
		t.Errorf("want:%d random numbers, got:%d", n, len(slice))
	}
}
