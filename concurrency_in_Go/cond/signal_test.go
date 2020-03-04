package cond__test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestSignal(t *testing.T) {
	queue := make([]struct{}, 0, 10)
	theCond := sync.Cond{L: &sync.Mutex{}}
	removeFromQueue := func(delay time.Duration) {
		time.Sleep(delay)
		theCond.L.Lock()
		queue = queue[1:]
		fmt.Println("remove from queue")
		theCond.L.Unlock()
		theCond.Signal()
	}
	for i := 0; i < 10; i++ {
		theCond.L.Lock()
		for len(queue) == 2 {
			theCond.Wait()
		}

		queue = append(queue, struct{}{})
		fmt.Println("add to queue")
		go removeFromQueue(time.Second)
		theCond.L.Unlock()
	}

}
