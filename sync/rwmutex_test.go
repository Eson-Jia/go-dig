package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var mu sync.RWMutex
var count int

func TestRW(t *testing.T) {
	go A()
	time.Sleep(2 * time.Second)
	mu.Lock()
	defer mu.Unlock()
	count++
	fmt.Println(count)
}
func A() {
	mu.RLock()
	defer mu.RUnlock()
	B()
}
func B() {
	time.Sleep(5 * time.Second)
	C()
}
func C() {
	mu.RLock()
	defer mu.RUnlock()
}

// A: 不能编译

// B: 输出 1

// C: 程序hang住

// D: panic

/*
选择B

*/

/*
答案错误，正确是panic，知识点如下：
如果已经有一组reader持有了读写锁，这个时候如果writer调用Lock,它会被阻塞。
接着如果有reader调用RLock, 等前面那组readerUnlock后， writer优先获取锁
*/
