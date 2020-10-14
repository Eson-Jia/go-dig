package dance

import "fmt"

//840. 矩阵中的幻方
func numMagicSquaresInside(grid [][]int) int {
	rows, columns := len(grid), len(grid[0])
	count := 0
	if rows < 3 || columns < 3 {
		return 0
	}
	for i := 2; i < rows; i++ {
		for j := 2; j < columns; j++ {
			if isMagic(grid, i, j) {
				count += 1
			}
		}
	}
	return count
}

func isMagic(grid [][]int, maxRow, maxColumn int) bool {
	set := make([]int, 10)
	for i := maxRow - 2; i <= maxRow; i++ {
		for j := maxColumn - 2; j <= maxColumn; j++ {
			if v := grid[i][j]; 1 <= v && v <= 9 {
				set[v]++
				fmt.Println("v", v)
			} else {
				// fmt.Println("v",v)
				return false
			}
		}
	}
	fmt.Println(set)
	for i := 1; i < 10; i++ {
		if set[i] != 1 {
			fmt.Println("222")
			return false
		}
	}
	if grid[maxRow-1][maxColumn-1] != 5 {
		return false
	}
	for j := maxColumn - 2; j <= maxColumn; j++ {
		if grid[maxRow-2][j]+grid[maxRow][maxColumn-j] != 10 {
			return false
		}
	}
	for i := maxRow - 2; i <= maxRow; i++ {
		if grid[i][maxColumn-2]+grid[maxColumn-i][maxColumn] != 10 {
			return false
		}
	}
	for j := maxColumn - 2; j <= maxColumn; j++ {

	}
	return true
}
