package main

import (
	"math/rand"
	"time"
)

func randomNumber(ch chan<- int) {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	ch <- rand.Int()
	close(ch)

}
