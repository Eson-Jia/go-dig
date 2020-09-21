package cond

import (
	"sync"
	"testing"
)

func TestGoRoutine(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		//wg.Add(1)             2
		go func() {
			wg.Add(1) //   1
			defer wg.Done()
			t.Log("done")
		}()
	}
	wg.Wait()
}
