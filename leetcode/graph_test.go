package dance_test

import (
	"testing"
)

func TestT(t *testing.T) {
	mats := [][][]int{
		{
			{0, 0, 1},
			{1, 0, 0},
			{1, 0, 0},
		},
		{
			{1, 0, 0},
			{0, 1, 0},
			{0, 0, 1},
		},
		{
			{0, 0, 0, 1},
			{1, 0, 0, 0},
			{0, 1, 1, 0},
			{0, 0, 0, 0},
		},
	}
	for _, mat := range mats {
		t.Log(numSpecial(mat))
	}
}

func numSpecial(mat [][]int) int {
	count := 0
	rowNums := len(mat)
	colNums := len(mat[0])
	rowSum := make([]int, rowNums)
	colsSum := make([]int, colNums)
	for i := 0; i < rowNums; i++ {
		for j := 0; j < colNums; j++ {
			rowSum[i] += mat[i][j]
			colsSum[j] += mat[i][j]
		}
	}
	for i := 0; i < rowNums; i++ {
		for j := 0; j < colNums; j++ {
			if mat[i][j] == 1 {
				if rowSum[i]+colsSum[j] == 2 {
					count++
				}
			}
		}
	}
	return count
}

func TestModifyString(t *testing.T) {
	//t.Log(modifyString("?zs"))
	//
	//t.Log(modifyString("ubv?w"))
	////"??yw?ipkj?"
	//t.Log(modifyString("??yw?ipkj?"))
	//"j?qg??b"
	t.Log(modifyString("j?qg??b"))
}

//205
func modifyString(s string) string {
	theSlice := []byte(s)
	theSet := make([]bool, 26)
	reset := func() {
		for i := 0; i < 26; i++ {
			theSet[i] = false
		}
	}
	for i := 0; i < len(theSlice); i++ {
		if theSlice[i] == '?' {
			if i >= 1 {
				theSet[theSlice[i-1]-'a'] = true
			}
			if i < len(theSlice)-1 && theSlice[i+1] != '?' {
				theSet[theSlice[i+1]-'a'] = true
			}
			var origin = -1
			for theByte, good := range theSet {
				if !good {
					origin = theByte
					break
				}
			}
			if origin == -1 {
				panic("not found")
			}
			theSlice[i] = byte(origin + 'a')
			reset()
		}
	}
	return string(theSlice)
}

func TestTriplets(t *testing.T) {
	t.Log(numTriplets([]int{7, 7, 8, 3}, []int{1, 2, 9, 7}))
}

func numTriplets(nums1 []int, nums2 []int) int {
	count := 0
	theFunc := func(nums1 []int, nums2 []int) {
		for i := 0; i < len(nums1); i++ {
			for j := 0; j < len(nums2)-1; j++ {
				for k := j + 1; k < len(nums2); k++ {
					if (nums1[i] > nums2[j] && nums1[i] > nums2[k]) || (nums1[i] < nums2[j] && nums1[i] < nums2[k]) {
						continue
					}
					if nums1[i]*nums1[i] == nums2[j]*nums2[k] {
						count++
					}
				}
			}
		}
	}
	theFunc(nums1, nums2)
	theFunc(nums2, nums1)
	return count
}

// 1041. 困于环中的机器人
func isRobotBounded(instructions string) bool {
	const (
		Min = 0
		N   = 1
		E   = 2
		S   = 3
		W   = 4
		Max = 5
	)
	x, y := 0, 0
	direct := N
	G := func() {
		switch direct {
		case N:
			y += 1
		case E:
			x += 1
		case S:
			y -= 1
		case W:
			x -= 1
		}
	}
	L := func() {
		direct -= 1
		if direct == Min {
			direct = W
		}
	}
	R := func() {
		direct += 1
		if direct == Max {
			direct = N
		}
	}
	theInstructions := []byte(instructions)
	for _, instruction := range theInstructions {
		switch instruction {
		case 'G':
			G()
		case 'L':
			L()
		case 'R':
			R()
		}
	}
	if (x == 0 && y == 0) || direct != N {
		return true
	}
	return false
}

func TestGardenNoAdj(t *testing.T) {
	suits := []struct {
		N     int
		Paths [][]int
	}{
		{
			3,
			[][]int{
				[]int{1,2},
				[]int{2,3},
				[]int{1,3},
			},
		},
		{
			4,
			[][]int{
				[]int{1,2},
				[]int{3,4},
			},
		},
		{
			4,
			[][]int{
				[]int{1, 2},
				[]int{2, 3},
				[]int{3, 4},
				[]int{4, 1},
				[]int{1, 3},
				[]int{2, 4},
			},
		},
	}
	for _, suit := range suits {
		t.Log(gardenNoAdj(suit.N, suit.Paths))
	}
}

// 1042. 不邻接植花
// version 1.0
// 这里的想法是先挑一个点(A)，染成第一种颜色(1)，然后找到与这个点相邻的点(比如B,C,D)，分别染成其他三种颜色(2,3,4)，
// 然后再找到跟这些点相邻的点，再染色，就这样广度优先遍历染色。
// 这个方法是有问题的，问题就是在给B染色的时候是依据A的颜色，同时还依据了(C,D)的颜色，而B与(C,D)可能并不相邻，如果(B,C,D)不相邻的话，
// 可以将(B,C,D)全都都染成颜色 2
func gardenNoAdj(N int, paths [][]int) []int {
	graph := make([][]int, N+1)
	for _, path := range paths {
		a, b := path[0], path[1]
		if graph[a] == nil {
			graph[a] = make([]int, 0)
		}
		if graph[b] == nil {
			graph[b] = make([]int, 0)
		}
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}
	colored := make([]int, N+1)
	isOpened := make([]bool, N+1)
	wf := func(point int) {
		queue := make([]int, 0)
		queue = append(queue, point)
		colored[point] = 1
		for len(queue) > 0 {
			head := queue[0]
			queue = queue[1:]
			allColor := []bool{false, true, true, true, true}
			allColor[colored[head]] = false
			if isOpened[head] {
				continue
			}
			allPoints := graph[head]
			for _, point := range allPoints {
				if colored[point] == 0 {
					for colorIndex := 0; colorIndex < len(allColor); colorIndex++ {
						if allColor[colorIndex] {
							allColor[colorIndex] = false
							colored[point] = colorIndex
							break
						}
					}
				}
				queue = append(queue, point)
			}
			isOpened[head] = true
		}
	}
	for i := 1; i <= N; i++ {
		if !isOpened[i] {
			wf(i)
		}
	}
	return colored[1:]
}

// 1042. 不邻接植花
func gardenNoAdjSecond(N int, paths [][]int) []int {
	graph := make([][]int, N+1)
	for _, path := range paths {
		a, b := path[0], path[1]
		if graph[a] == nil {
			graph[a] = make([]int, 0)
		}
		if graph[b] == nil {
			graph[b] = make([]int, 0)
		}
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}
	colored := make([]int, N+1)
	for i := 1; i < N; i++ {
	}
	return colored[1:]
}