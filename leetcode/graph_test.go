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
				[]int{1, 2},
				[]int{2, 3},
				[]int{1, 3},
			},
		},
		{
			4,
			[][]int{
				[]int{1, 2},
				[]int{3, 4},
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
		t.Log(gardenNoAdjSecond(suit.N, suit.Paths))
	}
}

// 1042. 不邻接植花
// version 1.0
// 这里的想法是先挑一个点(A)，染成第一种颜色(1)，然后找到与这个点相邻的点(比如B,C,D)，分别染成其他三种颜色(2,3,4)，
// 然后再找到跟这些点相邻的点，再染色，就这样广度优先遍历染色。
// 这个方法是有问题的，问题就是在给B染色的时候依据与其相邻的A颜色的同时，还依据了(C,D)的颜色，而B与(C,D)可能并不相邻，如果(B,C,D)不相邻的话，
// 可以将(B,C,D)全都都染成颜色 2。我们在给一个顶点染色的时候，只需要参考与其相邻的顶点的颜色，这个我们没有做到这一点。

// version 2.0
// 因为顶点最多有三个条边，而有四种颜色，假设相邻的三个顶点颜色各不相同，那么该顶点也有第四个颜色可以使用。
// 所以无论如何都可以在某个顶点选择一种颜色不与相邻顶点的颜色重合
// 这里采用广度优先算法遍历无向图，因为该图不一定是强连通性的，所有需要遍历顶点来使用广度优先算法。

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
	// 广度优先
	breadthFirst := func(point int) {
		queue := make([]int, 0)
		queue = append(queue, point)
		for len(queue) > 0 {
			head := queue[0]
			queue = queue[1:]
			if colored[head] != 0 {
				continue
			}
			allPoints := graph[head]
			canUseColor := []bool{false, true, true, true, true}
			for _, point := range allPoints {
				if colored[point] != 0 {
					canUseColor[colored[point]] = false
				} else {
					queue = append(queue, point)
				}
			}
			for colorIndex, canUse := range canUseColor {
				if canUse {
					colored[head] = colorIndex
					break
				}
			}
		}
	}
	for i := 1; i <= N; i++ {
		if colored[i] == 0 {
			breadthFirst(i)
		}
	}
	return colored[1:]
}

// 1042. 不邻接植花
// 这里没有采用深度优先或者广度优先算法，因为不需要保存路径
// 直接遍历所有顶点，在为顶点染色的时候考虑邻接顶点的颜色
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
	for currentVertex := 1; currentVertex < N+1; currentVertex++ {
		canUseColor := []bool{false, true, true, true, true}
		for _, vertex := range graph[currentVertex] {
			if colored[vertex] != 0 {
				canUseColor[colored[vertex]] = false
			}
		}
		for colorIndex, canUse := range canUseColor {
			if canUse{
				colored[currentVertex]=colorIndex
				break
			}
		}
	}
	return colored[1:]
}
