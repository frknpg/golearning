package main

import "fmt"

func main() {
	//channel initialization
	unbufferedChan := make(chan int)

	//channel declaration
	var unbufferedChan2 chan int

	fmt.Println(unbufferedChan)
	fmt.Println(unbufferedChan2) //nil

	//only can read from channel
	go func(unbufChan <-chan int) {
		//blocks until data arrives
		value := <-unbufChan
		fmt.Println(value)
	}(unbufferedChan)

	unbufferedChan <- 10
}
