package dance

import (
	"fmt"
	"sync"
	"testing"
)

/**
交替打印数字和字母
使⽤两个 goroutine 交替打印序列，⼀个 goroutine 打印数字， 另外⼀个 goroutine 打印字⺟， 最终效果如下：
12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728

2021/8/20
额,没有认真读题,上来就干,结果题目要求都理解错了.题目要求两个 goroutine 交替打印序列,注意打印.就是在两个 goroutine 中调用打印函数.
printNumAndAB 完全错误,没有理解清楚题意.

15:51
一个问题,两个 goroutine 中包不包括 main goroutine?
安全起见不使用 main goroutine
实现了 interleavePrint,但是这个实现还是不够简洁,可以看出打印出来的字母和数组都是相互交替一一对应的,所以只需要在一个 goroutine 实现计数即可
*/
func printNumAndAB() {
	numChan, ABChan := make(chan int, 0), make(chan byte, 0)
	go func() {
		defer close(numChan)
		i := 1
		for i <= 28 {
			select {
			case numChan <- i:
				i++
			}
		}
	}()
	go func() {
		defer close(ABChan)
		A := byte('A')
		for A <= 'Z' {
			select {
			case ABChan <- A:
				A++
			}
		}
	}()
	done := false
	for {
		num, ok := <-numChan
		if !ok {
			break
		}
		fmt.Print(num)
		fmt.Print(<-numChan)
		if !done {
			ab, ok := <-ABChan
			if !ok {
				continue
				done = true
			}
			fmt.Printf("%c", ab)
			fmt.Printf("%c", <-ABChan)
		}

	}
}

func interleavePrint() {
	var wg sync.WaitGroup
	wg.Add(1)
	numSingal := make(chan struct{})
	alphaSignal := make(chan struct{})
	wg.Add(1)
	go func() {
		i := 1
		for {
			fmt.Print(i)
			i++
			fmt.Print(i)
			i++
			if i < 29 {
				alphaSignal <- struct{}{}
				<-numSingal
			} else {
				break
			}
		}
		defer wg.Done()
	}()
	go func() {
		defer wg.Done()
		for a := byte('A'); a <= byte('Z'); {
			<-alphaSignal
			fmt.Printf("%c", a)
			a++
			fmt.Printf("%c", a)
			a++
			numSingal <- struct{}{}
		}
	}()
	wg.Wait()
}
func TestPrintNumAndAB(t *testing.T) {
	//printNumAndAB()
	interleavePrint()
}
