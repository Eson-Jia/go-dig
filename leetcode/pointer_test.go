package dance

import (
	"fmt"
	"testing"
)

func xor() int {
	result := 0
	for i := 0; i < 10; i++ {
		fmt.Printf("%08b^%08b=", result, i)
		result ^= i
		fmt.Printf("%08b\n", result)
	}
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Printf("%08b^%08b=", result, i)
		result ^= i
		fmt.Printf("%08b\n", result)
	}
	return result
}

func TestXor(t *testing.T) {
	t.Log(xor())
}
