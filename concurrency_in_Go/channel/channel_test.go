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
		doneC := make(chan interface{})
		go func() {
			defer close(doneC)
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
				case <-or(channels[2:]...):
				}
			}
		}()
		return doneC
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
		sig(time.Second),
		sig(time.Minute),
		sig(time.Hour),
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
	OutValue := func(done chan interface{}, inStream <-chan int, prefix string) {
		for value := range inStream {
			select {
			case <-done:
				return
			default:

			}
			fmt.Println(prefix, value)
		}
	}
	done := make(chan interface{})
	defer close(done)
	outStream := generator(done, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	var wg sync.WaitGroup
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go func(index int) {
			defer wg.Done()
			OutValue(done, outStream, fmt.Sprintf("%d", index))
		}(i)
	}
	wg.Wait()
}
