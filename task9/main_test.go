package main

import (
	"math"
	"testing"
)

func TestPipeliner(t *testing.T) {

	input := make(chan uint8, 5)
	output := make(chan float64, 5)

	go pipeline(input, output)

	expectedValues := []uint8{1, 2, 3, 4, 5}
	for _, val := range expectedValues {
		input <- val
	}
	close(input)

	var results []float64
	for res := range output {
		results = append(results, res)
	}

	for i, val := range expectedValues {
		expected := math.Pow(float64(val), 3)
		if results[i] != expected {
			t.Errorf("Expected %.2f, but got %.2f", expected, results[i])
		}
	}
}
