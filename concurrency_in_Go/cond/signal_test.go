package cond_test

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
		before := len(queue)
		queue = queue[1:]
		fmt.Printf("before len:%d,remove from queue,after len:%d\n", before, len(queue))
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

func TestBroadcast(t *testing.T) {
	type Button struct {
		Clicked *sync.Cond
	}
	button := Button{Clicked: &sync.Cond{L: &sync.Mutex{}}}
	subscribe := func(button *Button, fn func()) {
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			wg.Done()
			button.Clicked.L.Lock()
			defer button.Clicked.L.Unlock()
			button.Clicked.Wait()
			fn()
		}()
		wg.Wait()
	}
	var finished sync.WaitGroup
	finished.Add(3)
	subscribe(&button, func() {
		defer finished.Done()
		fmt.Println("maximizing windows")
	})
	subscribe(&button, func() {
		defer finished.Done()
		fmt.Println("displaying annoying dialog box!")
	})
	subscribe(&button, func() {
		defer finished.Done()
		fmt.Println("Mouse clicked.")
	})
	button.Clicked.Broadcast()
	finished.Wait()
}
