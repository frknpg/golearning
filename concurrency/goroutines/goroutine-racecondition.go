package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// race condition farkli threadlerin ayni memorye erismeye calismasi
func main() {
	// raceExample()
	// raceExampleFixed()
	raceExampleFixedWithAtomic()
}

func raceExample() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	//shared value
	val := 0

	go func() {
		for i := 0; i < 100000000; i++ {
			val++
		}

		wg.Done()
	}()

	go func() {
		for i := 0; i < 100000000; i++ {
			val++
		}

		wg.Done()
	}()

	wg.Wait()

	fmt.Println(val)
}

func raceExampleFixed() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	lock := sync.Mutex{}

	//shared value
	val := 0

	go func() {
		for i := 0; i < 100000000; i++ {
			lock.Lock()
			val++
			lock.Unlock()
		}

		wg.Done()
	}()

	go func() {
		for i := 0; i < 100000000; i++ {
			lock.Lock()
			val++
			lock.Unlock()
		}

		wg.Done()
	}()

	wg.Wait()

	fmt.Println(val)
}

func raceExampleFixedWithAtomic() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	//shared value
	var val int32 = 0

	loopFunc := func() {
		for i := 0; i < 100000000; i++ {
			atomic.AddInt32(&val, 1)
		}

		wg.Done()
	}

	go loopFunc()
	go loopFunc()

	wg.Wait()

	fmt.Println(val)
}
