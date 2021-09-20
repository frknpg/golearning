package main

import (
	"fmt"
	"sync"
)

type RaceTest struct {
	Val int
}

func main() {
	raceTest := &RaceTest{}

	locker := &sync.Mutex{}

	wg := &sync.WaitGroup{}
	wg.Add(10000)

	for i := 0; i < 10000; i++ {
		go increment(raceTest, wg, locker)
	}

	wg.Wait()
	fmt.Println(*raceTest)
}

func increment(rt *RaceTest, wg *sync.WaitGroup, locker *sync.Mutex) {
	locker.Lock()
	rt.Val += 1
	wg.Done()
	locker.Unlock()
}
