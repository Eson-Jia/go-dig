package channel_test

import (
	"fmt"
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
