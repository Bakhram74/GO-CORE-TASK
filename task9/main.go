package main

import (
	"fmt"
	"math"
)

func main() {
	input := make(chan uint8)
	output := make(chan float64)
	go pipeline(input, output)

	go func() {
		for i := 1; i <= 5; i++ {
			input <- uint8(i)
		}
		close(input)
	}()
	for n := range output {
		fmt.Println(n)
	}
}

func pipeline(input <-chan uint8, output chan<- float64) {
	for n := range input {
		output <- math.Pow(float64(n), 3)
	}
	close(output)
}

// func main() {
// 	ch1 := make(chan uint8)

// 	go func() {
// 		for i := 1; i <= 5; i++ {
// 			ch1 <- uint8(i)
// 		}
// 		close(ch1)
// 	}()

// 	ch2 := pipeline(ch1)
// 	for n := range ch2 {
// 		fmt.Println(n)
// 	}
// }

// func pipeline(ch chan uint8) <-chan float64 {
// 	ch2 := make(chan float64)
// 	go func() {
// 		wg := sync.WaitGroup{}
// 		for n := range ch {
// 			wg.Add(1)
// 			go func() {
// 				defer wg.Done()

// 				ch2 <- math.Pow(float64(n), 3)
// 			}()
// 		}
// 		wg.Wait()
// 		close(ch2)
// 	}()

// 	return ch2
// }
