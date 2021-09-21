package main

import "fmt"

func main() {
	//channel initialization
	unbufferedChan := make(chan int, 5)

	unbufferedChan <- 1
	unbufferedChan <- 2
	unbufferedChan <- 3
	unbufferedChan <- 4
	unbufferedChan <- 5

	fmt.Println(<-unbufferedChan)
	fmt.Println(<-unbufferedChan)
	fmt.Println(<-unbufferedChan)
	fmt.Println(<-unbufferedChan)
	fmt.Println(<-unbufferedChan)
}
