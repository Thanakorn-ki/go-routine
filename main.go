package main

import (
	"fmt"
)

var asset int // for Ex 3

func main() {

	/*
		This is Channel receipt and send
		var chOut <-chan int
		var chIn chan<- int
	*/

	// Ex 1
	// ch := make(chan struct{})
	// go echo(ch)
	// <-ch

	// Ex 2
	// chNum := make(chan int)
	// chQ := make(chan struct{})
	// go fibonacci(chNum, chQ)

	// for i := 0; i < 10; i++ {
	// 	fmt.Println(<-chNum)
	// }
	// chQ <- struct{}{}

	// ex 3 Rece condition
	go setAsset()
	for i := 0; i < 20; i++ {
		fmt.Println(asset)
	}
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

func setAsset() {
	for i := 0; i < 100; i++ {
		asset = 1
	}
}
