package main

import "testing"

func TestRandomNumber(t *testing.T) {
	ch1 := make(chan int)
	go randomNumber(ch1)
	if <-ch1 == 0 {
		t.Errorf("want random numbers, got: %d", 0)
	}

	ch2 := make(chan int)
	go randomNumber(ch2)
	if <-ch1 == <-ch2 {
		t.Errorf("ch1:%d equal to ch2:%d", ch1, ch2)
	}

}
