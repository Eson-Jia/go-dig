package dance

import (
	"testing"
)

func TestRecursiveActivitySelector(t *testing.T) {
	s := []int{0, 1, 3, 0, 5, 3, 5, 6, 8, 8, 2, 12}
	f := []int{0, 4, 5, 6, 7, 9, 9, 10, 11, 12, 14, 16}
	result := []int{1, 4, 8, 11}
	got := recursiveActivitySelector(s, f, 0, 11)
	got1 := greedyActivitySelector(s, f, 11)
	if !IntEqual(result, got) {
		t.Error("failed in test")
	}
	if !IntEqual(result, got1) {
		t.Error("failed in test")
	}
}

func TestMaxSUmRangeQuery(t *testing.T) {
	suits := []struct {
		Nums    []int
		Request [][]int
	}{
		{
			[]int{1, 2, 3, 4, 5},
			[][]int{
				[]int{
					1, 3,
				},
				[]int{0, 1},
			},
		},
	}
	for _, suit := range suits {
		maxSumRangeQuerySecond(suit.Nums, suit.Request)
	}
}
