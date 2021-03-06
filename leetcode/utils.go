package dance

func getMax(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func IntEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
