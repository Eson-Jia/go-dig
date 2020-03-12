package channel_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func Test_Or(t *testing.T) {
	var or func(...<-chan interface{}) <-chan interface{}
	or = func(channels ...<-chan interface{}) <-chan interface{} {
		switch len(channels) {
		case 0:
			return nil
		case 1:
			return channels[0]
		}
		orDone := make(chan interface{})
		go func() {
			defer close(orDone)
			switch len(channels) {
			case 2:
				select {
				case <-channels[0]:
				case <-channels[1]:
				}
			default:
				select {
				case <-channels[0]:
				case <-channels[1]:
				case <-channels[2]:
				case <-or(append(channels[3:], orDone)...):
				}
			}
		}()
		return orDone
	}
	sig := func(duration time.Duration) <-chan interface{} {
		done := make(chan interface{})
		go func() {
			defer fmt.Printf("return sig inner routine\n")
			defer close(done)
			time.Sleep(duration)
		}()
		return done
	}
	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	duration := time.Since(start)
	fmt.Printf("duration is :%v\n", duration)
}

// 这个测试有 panic 因为 close of closed channel
// 因为 repeat 中的 for 所在的 goroutine 和 main goroutine并发运行
// 当 repeat 往 chan 中传入超过10个数字的时候,close(done)就会被调用至少两次
// 使用 sync.Once 包着 close(done)可以避免这个问题
func Test_Pipeline(t *testing.T) {
	var once sync.Once
	count := 0
	done := make(chan interface{})
	for value := range repeat(done, func() interface{} {
		return 1
	}) {
		count += 1
		if count > 10 {
			once.Do(func() {
				close(done)
			})
		}
		fmt.Println(value)
	}
}

func Test_Pipeline1(t *testing.T) {
	multiple := func(done <-chan interface{}, inStream <-chan int, multiper int) <-chan int {
		outStream := make(chan int)
		go func() {
			defer close(outStream)
			for value := range inStream {
				select {
				case <-done:
					return
				case outStream <- value * multiper:
				}
			}
		}()
		return outStream
	}
	done := make(chan interface{})
	for value := range multiple(done, Generator(done, 10), 10) {
		fmt.Println(value)
	}
}

func echoValue(done chan interface{}, inStream <-chan int, prefix string) <-chan int {
	outStream := make(chan int)
	go func() {
		defer close(outStream)
		for value := range inStream {
			select {
			case <-done:
			case outStream <- value:
				time.Sleep(time.Second)
				fmt.Println(prefix, value)
			}
		}
	}()
	return outStream
}

func Test_Fan_Out_In_Chain(t *testing.T) {
	done := make(chan interface{})
	defer close(done)
	outStream := Generator(done, 50)
	echoValue(done, FanIn(done, FanOut(done, outStream, 10)...), "only_one")
}

func Test_Fan_Out_Fan_In(t *testing.T) {
	// 使用 fan-out 拆成多个 channel
	theSize := 0
	var mutex sync.Mutex
	finished := make(chan interface{})
	blackHole := func(done <-chan interface{}, inStream <-chan int, amount int) {
		go func() {
			for v := range inStream {
				select {
				case <-done:
					return
					// 注意,如果不加 default 就会一直阻塞 select 直到 done 被 close
				default:
					fmt.Sprintln(v)
				}
				mutex.Lock()
				theSize++
				if theSize >= amount {
					close(finished)
				}
				mutex.Unlock()
			}
		}()
	}
	size := 100
	done := make(chan interface{})
	defer close(done)
	outStreams := FanOut(done, Generator(done, size), 10)
	for index, outStream := range outStreams {
		blackHole(done, echoValue(done, outStream, fmt.Sprintf("%d", index)), size)
	}
	<-finished
}

func TestChanSlice(t *testing.T) {
	channels := make([]chan int, 5)
	//channels 里面的 chan 没有调用(make)初始化, 值都是 nil,需要初始化
	//值为 nil 的 chan 读取和写入值都会阻塞,close 会 panic
	channels[0] = make(chan int)
	go func() {
		close(channels[0])
		channels[0] <- 12
	}()
	fmt.Println(<-channels[0])
}

func repeat(done <-chan interface{}, fn func() interface{}) <-chan interface{} {
	outStream := make(chan interface{})
	go func() {
		defer close(outStream)
		for {
			select {
			case <-done:
				return
			case outStream <- fn():
			}
		}
	}()
	return outStream
}

func Generator(done <-chan interface{}, until int) <-chan int {
	outStream := make(chan int)
	go func() {
		defer close(outStream)
		for i := 0; i < until; i++ {
			select {
			case <-done:
				return
			case outStream <- i:
			}
		}
	}()
	return outStream
}

func FanOut(done <-chan interface{}, inStream <-chan int, outSize int) []<-chan int {
	var streams []chan int
	var outStreams []<-chan int
	for i := 0; i < outSize; i++ {
		stream := make(chan int)
		streams = append(streams, stream)
		outStreams = append(outStreams, stream)
	}
	for i := 0; i < outSize; i++ {
		go func(outStream chan int) {
			defer close(outStream)
			for v := range inStream {
				select {
				case <-done:
					return
				case outStream <- v:
				}
			}
		}(streams[i])
	}
	return outStreams
}

func FanIn(done <-chan interface{}, inStreams ...<-chan int) <-chan int {
	multiplexedStream := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(inStreams))
	go func() {
		defer close(multiplexedStream)
		for _, inStream := range inStreams {
			go func(stream <-chan int) {
				defer wg.Done()
				for value := range stream {
					select {
					case <-done:
						return
					case multiplexedStream <- value:
					}
				}
			}(inStream)
		}
		wg.Wait()
	}()
	return multiplexedStream
}
