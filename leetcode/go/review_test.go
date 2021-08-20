package dance

import (
	"fmt"
	"testing"
)

/**
交替打印数字和字母
使⽤两个 goroutine 交替打印序列，⼀个 goroutine 打印数字， 另外⼀个 goroutine 打印字⺟， 最终效果如下：
解题思路
问题很简单，使⽤ channel 来控制打印的进度。使⽤两个 channel ，来分别控制数字和
字⺟的打印序列， 数字打印完成后通过 channel 通知字⺟打印, 字⺟打印完成后通知数
字打印，然后周⽽复始的⼯作。
源码参考
1 12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728
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
func TestPrintNumAndAB(t *testing.T) {
	printNumAndAB()
}
