package channel_test

import (
	"fmt"
	"testing"
	"time"
)

func Tee(done <-chan interface{}, inStream <-chan interface{}) (_, _ <-chan interface{}) {
	chan1, chan2 := make(chan interface{}), make(chan interface{})
	go func() {
		defer close(chan1)
		defer close(chan2)
		for value := range inStream {
			select {
			case <-done:
				return
			default:
			}
			chan1, chan2 := chan1, chan2
			for i := 0; i < 2; i++ {
				select {
				case chan1 <- value:
					chan1 = nil
				case chan2 <- value:
					chan2 = nil
				}
			}
		}
	}()
	return chan1, chan2
}

func Test_Tee(t *testing.T) {
	Generator := func(done <-chan interface{}, until int) <-chan interface{} {
		out := make(chan interface{})
		go func() {
			defer close(out)
			for i := 0; i < until; i++ {
				select {
				case out <- i:
				case <-done:
					return
				}
			}
		}()
		return out
	}
	printValue := func(done <-chan interface{}, inStream <-chan interface{}, prefix string) {
		go func() {
			for value := range inStream {
				select {
				case <-done:
				default:
					t.Log(prefix, value)
				}
			}
		}()
	}
	done := make(chan interface{})
	defer close(done)

	chan1, chan2 := Tee(done, Generator(done, 10))
	printValue(done, chan1, "first")
	printValue(done, chan2, "second")
	time.Sleep(time.Second)
}

func TestFuncParamOrder(t *testing.T) {
	param1 := func() string {
		fmt.Println("param1")
		return "param1"
	}
	param2 := func() string {
		fmt.Println("param2")
		return "param2"
	}
	Sum := func(param1, param2 string) {
		fmt.Println(param1, param2)
	}
	Sum(param1(), param2())
}
