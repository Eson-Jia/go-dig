package main

import (
	"fmt"
	"testing"
)

type Person interface {
	job()
	groupUp()
}

func whatJob(p Person) {
	p.job()
}

func groupUp(p Person) {
	p.groupUp()
}

type Programmer struct {
	age int
}

func (p *Programmer) groupUp() {
	p.age++
}

func (p *Programmer) job() {
	fmt.Println("study")
}

type Student struct {
	age int
}

func (s *Student) groupUp() {
	s.age += 10
}

func (s *Student) job() {
	fmt.Println("coding")
}

func TestInterface(t *testing.T) {
	s := &Student{age: 10}
	groupUp(s)
	whatJob(s)
	p := &Programmer{age: 10}
	whatJob(p)
	groupUp(p)
}
