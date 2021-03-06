package main

import (
	"fmt"
	"testing"
)

type I1 interface {
	F1()
	F2()
}

type I2 interface {
	F3()
	F4()
}

type S1 struct {
	field1 string
}

func (s S1) F1() {
	fmt.Println("F1")
}

func (s S1) F2() {
	fmt.Println("F2")
}

func (s S1) F3() {
	fmt.Println("F3")
}

func (s S1) F4() {
	fmt.Println("F4")
}

func TestTypeAssertion(t *testing.T) {
	var eface interface{} = S1{}
	I1face := eface.(I1)
	I1face.F1()
	I1face.F2()
	I2face := I1face.(I2)
	I2face.F3()
	I2face.F4()
	s1 := I2face.(S1)
	s1.F1()
	s1.F2()
	var einterface interface{}
	_, ok := interface{}(einterface).(interface{})
	fmt.Println(ok)
}

func TestInterfaceTypeAssertion(t *testing.T) {
	key := "timeout"
	options := map[string]interface{}{
		key: int(100),
	}
	if v, ok := options[key]; ok {
		fmt.Println("timeout:", v.(uint)) // panic: interface conversion: interface {} is int, not uint
	}
}
