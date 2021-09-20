package main

import (
	"fmt"
)

func main() {
	go func(param string) {
		fmt.Println("Hello from", param)
	}("goroutine")

	fmt.Println("Hello from main")
}
