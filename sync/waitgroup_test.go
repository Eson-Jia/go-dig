package main

import (
	"sync"
	"testing"
	"time"
)

func TestWaitGroup(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		time.Sleep(time.Millisecond)
		wg.Done()
		wg.Add(1)
	}()
	// time.Sleep(time.Second)    1
	wg.Wait()
}

// A: 不能编译

// B: 无输出，正常退出

// C: 程序hang住

// D: panic

/*
我的答案是 程序hang住 因为 Add(1)会导致 Wait()一直等待
*/

/*
正确答案是 panic 因为 WaitGroup is reused before previous Wait has returned.
反思：不会WaitGroup正确操作方法,没认真读代码，没有注意主goroutine会阻塞在Wait
*/
