package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func(param string) {
		fmt.Println("Hello from", param)
		wg.Done()
	}("goroutine")

	wg.Wait()

	for i := 0; i < 10; i++ {

		// bu i referansini routinede isler ve inin son degerini ekrana tum routineler icin yazar
		// go func() {
		// 	fmt.Println(i)
		// }()

		// bu i degerini ger routine icin kopyalar ve 1 den 10 a tum sayilar yazilir
		go func(val int) {
			fmt.Println(val)
		}(i)
	}

	fmt.Println("Hello from main")
	time.Sleep(3 * time.Second)
}
