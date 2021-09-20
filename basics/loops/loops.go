package main

import (
	"fmt"
	"time"
)

func main() {

	slc_1 := []int{1, 2, 3}
	slc_2 := []int{}

	for i, value := range slc_1 {
		fmt.Printf("index: %d value: %d \n", i, value)
	}

	for i := 0; i < 10; i++ {
		slc_2 = append(slc_2, i)
	}

	for i := range slc_2 {
		fmt.Printf("index: %d \n", i)
	}

	for _, value := range slc_2 {
		fmt.Printf("value: %d \n", value)
	}

	c := time.After(5 * time.Second)

	for {
		b := false

		select {
		case <-c:
			b = true
		default:
			fmt.Println(time.Now())
			time.Sleep(1 * time.Second)
		}

		if b {
			break
		}
	}
}
