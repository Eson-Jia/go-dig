package pearls

import (
	"fmt"
	"testing"
)

func searchElement(pos int, source []int) int {
	smallest := make([]int, 0)
	for _, ele := range source {
		smallest = append(smallest, ele)
		for i := len(smallest) - 2; i >= 0; i-- {
			if smallest[i] > ele {
				smallest[i+1] = smallest[i]
			} else {
				smallest[i-1] = ele
			}
		}
		smallest = smallest[:pos]
	}
	return smallest[pos]
}

func TestSearch(t *testing.T) {

}

func TestSlice(t *testing.T) {
	fmt.Println([]int{1, 2, 3, 4}[:3])
}
