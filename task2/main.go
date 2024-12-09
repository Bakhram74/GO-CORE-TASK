package main

import (
	"fmt"
	"math/rand"
	"time"
)

var originalSlice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func randomNumbers(length, min, max int) []int {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	slice := make([]int, length)
	for i := range slice {
		slice[i] = rand.Intn(max-min) + min
	}
	return slice
}

func sliceExample(slice []int) []int {
	var newSlice []int
	for _, n := range slice {
		if n%2 == 0 {
			newSlice = append(newSlice, n)
		}
	}
	return newSlice
}

func addElements(slice []int, number int) []int {
	newSlice := append([]int{}, slice...)
	return append(newSlice, number)
}

func copySlice(slice []int) []int {
	return append([]int{}, slice...)
}

func removeElement(slice []int, idx int) []int {
	newSlice := make([]int, 0, len(slice)-1)
	for i, n := range slice {
		if i != idx {
			newSlice = append(newSlice, n)
		}
	}
	return newSlice
}

func main() {
	fmt.Printf("randomNumbers %v\n", randomNumbers(10, 0, 9))
	fmt.Printf("sliceExample %v\n", sliceExample(originalSlice))
	fmt.Printf("addElements %v\n", addElements(originalSlice, 80))
	fmt.Printf("copySlice %v\n", copySlice(originalSlice))
	fmt.Printf("removeElement %v\n", removeElement(originalSlice, 5))

}
