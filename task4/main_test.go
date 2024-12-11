package main

import (
	"reflect"
	"testing"
)

func TestDifference(t *testing.T) {
	tests := []struct {
		slice1   []string
		slice2   []string
		expected []string
	}{
		{
			slice1:   []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"},
			slice2:   []string{"banana", "date", "fig"},
			expected: []string{"apple", "cherry", "43", "lead", "gno1"},
		},
		{
			slice1:   []string{"a", "b", "c"},
			slice2:   []string{"a", "b"},
			expected: []string{"c"},
		},
	}

	for _, test := range tests {
		result := sliceDifference(test.slice1, test.slice2)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("difference(%v, %v) = %v; want %v", test.slice1, test.slice2, result, test.expected)
		}
	}
}
