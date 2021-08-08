package dance

import (
	"testing"
)

/**
4,9,25,等
可以发现这些都是素数 i 的平方
*/
func isThree(n int) bool {
	if n <= 3 {
		return false
	}
	he := make([]bool, n/2+1)
	for i := 2; i <= n/2; i++ {
		if he[i] {
			if i*i == n {
				return false
			}
			continue
		}
		if i*i == n {
			return true
		}
		for j := i; j <= n/2; j += i {
			he[j] = true
		}
	}
	return false
}

/**
尽可能平分一组数

*/
func numberOfWeeks(milestones []int) int64 {
	return 0
}

func TestIsTree(t *testing.T) {
	t.Log(isThree(81))
}
