package main_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type test struct {
	Num   int
	mutex sync.Mutex
}

func Test(theT *testing.T) {
	t := &test{}
	go func() {
		for i := 0; i < 100; i++ {
			time.Sleep(5 * time.Microsecond)
			num := t.Num
			time.Sleep(5 * time.Microsecond)
			t.mutex.Lock()
			t.Num = num + 1
			t.mutex.Unlock()
		}
	}()
	go func() {
		for i := 0; i < 100; i++ {
			num := t.Num
			time.Sleep(time.Microsecond * 5)
			t.mutex.Lock()
			t.Num = num + 1
			t.mutex.Unlock()
		}
	}()
	time.Sleep(1 * time.Second)
	fmt.Println(t.Num)
}
