func f() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func f() (r int) {
	t := 5
	r = t
	func() {
		t = t + 5
	}()
}


<!-- package main

import "fmt"

func modifyReturnValue(in int) (r int) {
	r = in
	fmt.Println("r:", r)
	defer func() {
		r++
	}()
	return in * 2
}

func main() {
	fmt.Println("after", modifyReturnValue(100))
}

// func f() (r int) {
// 	defer func(r int) {
// 		r = r + 5
// 	}(r)
// 	return 1
// } -->
