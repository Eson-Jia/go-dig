package dance

import "testing"

func coinChange(coins []int, amount int) int {
	return DPTop2Button(coins, amount)
}

func coinChangeTop2ButtonWithNote(coins []int, amount int) int {
	note := make(map[int]int, 0)
	return DPTop2ButtonWithNode(coins, note, amount)
}

func coinChangeButton2TopWithNote(coins []int, amount int) int {
	note := map[int]int{0: 0}
	for i := 1; i <= amount; i++ {
		min := amount + 1
		for _, coin := range coins {
			if i < coin {
				break
			}
			min = getMin(min, note[i-coin]+1)
		}
		note[i] = min
	}
	if note[amount] == amount+1 {
		return -1
	}
	return note[amount]
}

func DPTop2Button(coins []int, amount int) int {
	if amount < 0 {
		return -1
	}
	if amount == 0 {
		return 0
	}
	min := amount + 1
	for _, coin := range coins {
		if res := DPTop2Button(coins, amount-coin); res < 0 {
			continue
		}
		min = getMin(DPTop2Button(coins, amount-coin)+1, min)
	}
	if min == amount+1 {
		return -1
	}
	return min
}

func DPTop2ButtonWithNode(coins []int, note map[int]int, amount int) int {
	if amount < 0 {
		return 0
	}
	if amount == 0 {
		return 0
	}
	if v, ok := note[amount]; ok {
		return v
	}
	min := amount + 1
	for _, coin := range coins {
		result := DPTop2ButtonWithNode(coins, note, amount-coin)
		if result < 0 {
			continue
		}
		min = getMin(result+1, min)
	}
	if min == amount+1 {
		min = -1
	}
	note[amount] = min
	return min
}

func TestCoinChange(t *testing.T) {
	coins := []int{1, 2, 5}
	if result := coinChange(coins, 11); result != 3 {
		t.Errorf("unexpect result")
	}
	if result := coinChangeTop2ButtonWithNote(coins, 11); result != 3 {
		t.Errorf("unexpect result")
	}
	if result := coinChangeButton2TopWithNote(coins, 11); result != 3 {
		t.Errorf("unexpect result")
	}
}
