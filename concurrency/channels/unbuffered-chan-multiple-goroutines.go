package main

import "fmt"

func main() {
	//channel initialization
	unbufferedChan := make(chan int)

	//reader goroutuine
	go func(unbufChan chan int) {
		//blocks until data arrives
		value := <-unbufChan
		fmt.Println(value)
	}(unbufferedChan)

	//writer goroutuine
	go func(unbufChan chan int) {
		unbufChan <- 1
	}(unbufferedChan)

	fmt.Println("Multiple goroutines")
}
