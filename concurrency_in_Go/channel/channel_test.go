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

func Test_Pepiline(t *testing.T) {
	repeat := func(done <-chan interface{}, fn func() interface{}) <-chan interface{} {
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
	count := 0
	done := make(chan interface{})
	for value := range repeat(done, func() interface{} {
		return 1
	}) {
		count += 1
		if count > 10 {
			close(done)
		}
		fmt.Println(value)
	}

}

func Test_Pipeline1(t *testing.T) {
	generator := func(done <-chan interface{}, values ...int) <-chan int {
		outStream := make(chan int)
		go func() {
			defer close(outStream)
			for _, value := range values {
				select {
				case <-done:
					return
				case outStream <- value:
				}
			}
		}()
		return outStream
	}
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
	for value := range multiple(done, generator(done, 1, 2, 3, 4, 5, 6, 7), 10) {
		fmt.Println(value)
	}
}

func Test_Fan_Out(t *testing.T) {
	generator := func(done <-chan interface{}, until int) <-chan int {
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
	fan_in := func(done <-chan interface{}, inStreams ...<-chan int) <-chan int {
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

	OutValue := func(done chan interface{}, inStream <-chan int, prefix string) {
		for value := range inStream {
			select {
			case <-done:
				return
			default:

			}
			//time.Sleep(time.Second)
			fmt.Println(prefix, value)
		}
	}

	fan_out := func(done <-chan interface{}, inStream <-chan int, outSize int) []<-chan int {
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

	done := make(chan interface{})
	defer close(done)
	outStream := generator(done, 50)
	OutValue(done, fan_in(done, fan_out(done, outStream, 10)...), "only_one")
}

func TestChanSlice(t *testing.T) {
	channels := make([]chan int, 5)
	//
	channels[0] = make(chan int)
	go func() {
		channels[0] <- 12
	}()
	fmt.Println(<-channels[0])
}
