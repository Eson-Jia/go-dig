package dance

import (
	"fmt"
	"testing"
)

func TestButtonUpCutRod(t *testing.T) {
	price := [...]int{0, 1, 5, 8, 9, 10, 17, 17, 20, 24, 30}
	for i := 1; i <= 10; i++ {
		result := ButtonUpCutRod(price[:], i)
		fmt.Println("result:", result)
	}
}

func TestPrintButtonUpCutRod(t *testing.T) {
	price := [...]int{0, 1, 5, 8, 9, 10, 17, 17, 20, 24, 30}
	expectResult := struct {
		Amount [11]int
		First  [11]int
	}{
		[11]int{0, 1, 5, 8, 10, 13, 17, 18, 22, 25, 30},
		[11]int{0, 1, 2, 3, 2, 2, 6, 1, 2, 3, 10},
	}
	amount, first := ExtendedButtonUpCutRod(price[:], 10)
	for i := 0; i < 11; i++ {
		if expectResult.Amount[i] != amount[i] || expectResult.First[i] != first[i] {
			t.Fatal("unexpect result")
		}
	}
	PrintButtonUpCutRod(price[:], 10)
}

func TestMemoizedCutRod(t *testing.T) {
	price := [...]int{0, 1, 5, 8, 9, 10, 17, 17, 20, 24, 30}
	for i := 0; i < 11; i++ {
		fmt.Printf("len:%d result:%d\n", i, MemoizedCutRod(price[:], i))
	}
}
