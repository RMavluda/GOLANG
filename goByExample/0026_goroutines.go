package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var mu = sync.Mutex{}
	start := time.Now()
	for i := 0; i < 10; i++ {
		go func() {

			mu.Lock()
			for j := 0; j < 10; j++ {
				time.Sleep(time.Second)
			}
			mu.Unlock()
		}()
	}

	wg.Wait()
	diff := time.Now().Sub(start)
	fmt.Println(diff)

}
