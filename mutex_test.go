package main_test

import (
	"fmt"
	"sync"
	"testing"
)

var mu sync.Mutex
var chain string

func TestMutex(t *testing.T) {
	chain = "main"
	A()
	fmt.Println(chain)
}

func A() {
	mu.Lock()
	defer mu.Unlock()
	chain = chain + " --> A"
	B()
}

func B() {
	chain = chain + " --> B"
	C()
}

func C() {
	mu.Lock()
	defer mu.Unlock()
	chain = chain + " --> C"
}

// A: 不能编译

// B: 输出 main --> A --> B --> C

// C: 输出 main

// D: panic

/*
死锁,A函数第一行调用Lock,又在函数内调用B，B调用C的，C函数Lock会被阻塞。
*/
