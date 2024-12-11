package main

import (
	"fmt"
	"sync"
)

func main() {
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

	for n := range numbers {
		fmt.Println(n)
	}

}

func joinChannels(chnls ...<-chan int) <-chan int {

	out := make(chan int)

	go func() {
		wg := sync.WaitGroup{}
		wg.Add(len(chnls))

		for _, ch := range chnls {
			go func() {
				defer wg.Done()
				for n := range ch {
					out <- n
				}
			}()
		}
		wg.Wait()
		close(out)
	}()

	return out
}

// second variant
// func joinChannels(chnls ...<-chan int) <-chan int {

// 	out := make(chan int)
// 	wg := sync.WaitGroup{}

// 	for _, ch := range chnls {
// 		wg.Add(1)
// 		go func() {
// 			defer wg.Done()
// 			for n := range ch {
// 				out <- n
// 			}
// 		}()
// 	}

// 	go func() {
// 		wg.Wait()
// 		close(out)
// 	}()

// 	return out
// }
