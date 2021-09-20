package main

import (
	"fmt"
	"strings"
)

type String *string

func main() {
	var rstr String
	var pstr string

	fmt.Println(rstr)
	fmt.Println(pstr)

	pstr = "lorem ipsum"
	rstr = &pstr

	fmt.Println(rstr)
	fmt.Println(pstr)

	pstr = "lorem ipsum dolor"

	fmt.Println(*rstr)
	fmt.Println(pstr)

	*rstr = "lorem ipsum dolor sit amet"

	fmt.Println(*rstr)
	fmt.Println(pstr)

	replaceStr(pstr)

	fmt.Println(*rstr)
	fmt.Println(pstr)

	replaceStrP(rstr)

	fmt.Println(*rstr)
	fmt.Println(pstr)
}

func replaceStr(s string) {
	s = strings.Replace(s, "lorem", "LOREM", 1)
}

func replaceStrP(s String) {
	*s = strings.Replace(*s, "lorem", "LOREM", 1)
}
