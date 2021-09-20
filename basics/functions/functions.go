package main

import (
	"fmt"
	"time"
)

func main() {
	cikti, _ := timer(time.After(5 * time.Second))
	fmt.Println(cikti)

	timerAdvenced(time.After(5 * time.Second))
}

func timer(c <-chan time.Time, message ...string) (s string, e error) {
	defer fmt.Println("bir sonraki asamaya geciyor")
	defer fmt.Println("timer bitti") // last in first out

	s = "timer executed"
	e = nil
	for {
		select {
		case <-c:
			return
		default:
			fmt.Println(time.Now())
			time.Sleep(1 * time.Second)
		}
	}
}

func timerAdvenced(c <-chan time.Time) {
	defer fmt.Println("bir sonraki asamaya geciyor")
	defer fmt.Println("timer bitti") // last in first out

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	for {
		select {
		case <-c:
			return
		default:
			fmt.Println(time.Now())
			time.Sleep(1 * time.Second)

			if time.Now().Day() == 21 {
				panic("Beklenmeyen bir hata olustu")
			}
		}
	}
}
