package main

import "fmt"

var slc_1 []int

func main() {
	slc_2 := make([]int, 0, 3)

	//0 len oldugun icin atama yapilamiyor append gerekli
	// slc_2[0] = 2

	slc_2 = append(slc_2, 1)
	slc_2 = append(slc_2, 2)
	slc_2 = append(slc_2, 3)
	// capacity doldugunga capacity iki kartina cikarir
	// slc_2 = append(slc_2, 4)

	fmt.Println(slc_1, slc_2)
	//length and capacity
	fmt.Printf("slc_1 len: %d cap: %d \n", len(slc_1), cap(slc_1))
	fmt.Printf("slc_2 len: %d cap: %d \n", len(slc_2), cap(slc_2))
}
