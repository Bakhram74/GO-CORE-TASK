package main

import "testing"

func TestIntersection(t *testing.T) {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}

	isIntersection, slice := intersection(a, b)

	if !isIntersection {
		t.Errorf("Got %t want %t", isIntersection, true)
	}

	if len(slice) != 2 {
		t.Errorf("Got %d want %d", len(slice), 2)
	}
}
