package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomNumber(n int) <-chan int {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	ch := make(chan int)

	go func() {
		for i := 0; i < n; i++ {
			ch <- rand.Intn(n)
		}
		close(ch)
	}()
	return ch
}

func main() {
	numbers := randomNumber(10)
	for n := range numbers {
		fmt.Println(n)
	}

}
