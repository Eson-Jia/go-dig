package pearls

import (
	"fmt"
	"sort"
)

func bSearch(sorted []int, target int) int {
	if !sort.SliceIsSorted(sorted, func(i, j int) bool { return sorted[i] < sorted[j] }) {
		panic("slice is not sorted")
	}
	l, h := 0, len(sorted)-1
	for l <= h {
		m := (l + h) / 2
		fmt.Println(l, m, h)
		switch {
		case target < sorted[m]:
			h = m - 1
		case target > sorted[m]:
			l = m + 1
		case sorted[m] == target:
			return m
		}
	}
	return -1
}
