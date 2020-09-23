package dance

import (
	"fmt"
	"testing"
)

func TestButtonUpCutRod(t *testing.T) {
	price := []int{0, 1, 5, 8, 9, 10, 17, 17, 20, 24, 30}
	for i := 1; i <= 10; i++ {
		result := ButtonUpCutRod(price, i)
		if result != cutRod(price, i) {
			t.Error("test failed")
		}
		fmt.Println("result:", result)
	}
}

func TestCutRod(t *testing.T) {
	price := []int{0, 1, 5, 8, 9, 10, 17, 17, 20, 24, 30}
	for i := 1; i <= 10; i++ {
		result := cutRod(price[:], i)
		fmt.Println("result:", result)
	}
}

func TestPrintButtonUpCutRod(t *testing.T) {
	price := []int{0, 1, 5, 8, 9, 10, 17, 17, 20, 24, 30}
	expectResult := struct {
		Amount []int
		First  []int
	}{
		[]int{0, 1, 5, 8, 10, 13, 17, 18, 22, 25, 30},
		[]int{0, 1, 2, 3, 2, 2, 6, 1, 2, 3, 10},
	}
	amount, first := ExtendedButtonUpCutRod(price, 10)
	for i := 0; i < 11; i++ {
		if expectResult.Amount[i] != amount[i] || expectResult.First[i] != first[i] {
			t.Fatal("unexpect result")
		}
	}
	PrintButtonUpCutRod(price[:], 10)
}

func TestMemoizedCutRod(t *testing.T) {
	price := []int{0, 1, 5, 8, 9, 10, 17, 17, 20, 24, 30}
	for i := 0; i <= 10; i++ {
		if MemoizedCutRod(price, i) != cutRod(price, i) {
			t.Errorf("length:%d,MemoizedCutRod=%d,cutRod=%d", i, MemoizedCutRod(price, i), cutRod(price, i))
		}
	}
}

func TestWaysToStep(t *testing.T) {
	suits := []struct {
		Input  int
		Expect int
	}{
		{
			3,
			4,
		},
	}
	for _, suit := range suits {
		if result := waysToStep(suit.Input); suit.Expect != result {
			t.Errorf("input:%x,expect:%x,got:%x", suit.Input, suit.Expect, result)
		}
	}
}
