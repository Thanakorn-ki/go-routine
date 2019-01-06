package main

import (
	"fmt"
)

func main() {
	/*
		This is Channel receipt and send
		var chOut <-chan int
		var chIn chan<- int
	*/

	// Ex 1
	ch := make(chan struct{})
	go echo(ch)
	<-ch

	// Ex 2
	chNum := make(chan int)
	chQ := make(chan struct{})
	go fibonacci(chNum, chQ)

	for i := 0; i < 10; i++ {
		fmt.Println(<-chNum)
	}
	chQ <- struct{}{}
}

func echo(ch chan struct{}) {
	println("Hi")
	ch <- struct{}{}
}

func fibonacci(chNum chan int, chQ chan struct{}) {
	a, b := 0, 1
	for {
		select {
		case chNum <- a:
			a, b = b, a+b
		case <-chQ:
			return
		}
	}
}
