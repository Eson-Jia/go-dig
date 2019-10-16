package pearls

import (
	"fmt"
	"testing"
)

func TestBSearch(t *testing.T) {
	data := make([]int, 5)
	for i := 0; i < 5; i++ {
		data[i] = i * 5
	}
	fmt.Println("result:", bSearch(data, 40))
}
