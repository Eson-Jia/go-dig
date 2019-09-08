package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func Square(in <-chan int, out chan<- int) {
	for num := range in {
		out <- num * num
	}
	close(out)
}

func Print(in <-chan int) {
	for num := range in {
		fmt.Println("get a num:", num)
	}
}

func RandomProduce(output chan<- int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 10; i++ {
		output <- r.Int() % 10
	}
	close(output)
}

func TestPipeline(t *testing.T) {
	ch1, ch2 := make(chan int), make(chan int)
	go RandomProduce(ch1)
	go Square(ch1, ch2)
	Print(ch2)
}
