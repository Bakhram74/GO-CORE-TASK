package main

import (
	"reflect"
	"testing"
	"unsafe"
)

func TestRandomNumbers(t *testing.T) {
	length := 10
	min := 1
	max := 10
	numbers := randomNumbers(length, min, max)

	if length != len(numbers) {
		t.Errorf("invalid numbers length %d", len(numbers))
	}
	for _, n := range numbers {
		if n < min {
			t.Errorf("invalid min value of numbers  %d", n)
			return
		}
	}
	for _, n := range numbers {
		if n > max {
			t.Errorf("invalid max value of numbers  %d", n)
			return
		}
	}

}

func TestSliceExample(t *testing.T) {
	var originalSlice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := []int{2, 4, 6, 8, 10}
	slice := sliceExample(originalSlice)

	if !reflect.DeepEqual(expected, slice) {
		t.Error("slices not equal")
	}
}

func TestAddElements(t *testing.T) {
	slice := make([]int, 1, 4)

	element := 80
	result := addElements(slice, element)

	ptr1 := unsafe.Pointer(&slice[0])
	ptr2 := unsafe.Pointer(&result[0])

	if ptr1 == ptr2 {
		t.Error("slices must not be equal")
	}
}

func TestCopySlice(t *testing.T) {
	slice := make([]int, 5)

	newSlice := copySlice(slice)
	newSlice[0] = 1
	if slice[0] == newSlice[0] {
		t.Error("slices must not be equal")
	}
}

func TestRemoveElement(t *testing.T) {
	slice := make([]int, 0, 10)
	for i := 0; i < 10; i++ {
		slice = append(slice, i)
	}

	idx := 3

	length := len(slice)

	newSlice := removeElement(slice, idx)

	if length == len(newSlice) {
		t.Errorf("length %d and len(newSlice) %d is equal", length, len(newSlice)-1)
	}

	if newSlice[idx] == slice[idx] {
		t.Errorf("slices must not contain %d", slice[idx])
	}
}
