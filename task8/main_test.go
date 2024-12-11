package main

import (
	"testing"
)

func TestCustomWaitGroup(t *testing.T) {
	wg := NewWaitGroup()
	counter := 0
	expectedCount := 5

	for i := 0; i < expectedCount; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			counter++
		}(i)
	}

	wg.Wait()

	if counter != expectedCount {
		t.Errorf("want %v, got %v", expectedCount, counter)
	}
}
