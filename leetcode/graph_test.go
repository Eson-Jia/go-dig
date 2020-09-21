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

// 1577. 数的平方等于两数乘积的方法数
func numTriplets(nums1 []int, nums2 []int) int {
	multi := func(num []int) [][]int {
		result := make([][]int, len(num))
		for i := 0; i < len(num); i++ {
			result[i] = make([]int, len(num))
			for j := i; j < len(num); j++ {
				result[i][j] = num[i] * num[j]
			}
		}
		return result
	}
	multiNums1, multiNums2 := multi(nums1), multi(nums2)
	countNums := func(multiNums1 [][]int, multiNums2 [][]int) int {
		allCount := 0
		countMap := make(map[int]int)
		for i := 0; i < len(multiNums1); i++ {
			count, ok := countMap[multiNums1[i][i]]
			if ok {
				allCount += count
				continue
			}
			for j := 0; j < len(multiNums2)-1; j++ {
				for k := j + 1; k < len(multiNums2); k++ {
					if multiNums1[i][i] == multiNums2[j][k] {
						count++
					}
				}
			}
			allCount += count
			countMap[multiNums1[i][i]] = count
		}
		return allCount
	}
	return countNums(multiNums1, multiNums2) + countNums(multiNums2, multiNums1)
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
			if canUse {
				colored[currentVertex] = colorIndex
				break
			}
		}
	}
	return colored[1:]
}

// 172. 阶乘后的零
// 给定一个整数 n，返回 n! 结果尾数中零的数量。
// 尾数中零的数量就是10的多少次方，也就是多少个10相乘,10=10*1=5*2,分解成质数就是5*2
// 那么我们要做的就是统计（1....n）中5和2因子的个数，设5的个数为n,2的个数为m.那么min(n,m)即为尾数中零的个数，
// 由常识可知，n < m，所以我们只需要统计 5 因子的个数
// 在 5,10,15,20,25数中，(5,10,15,20) 分别包含了一个5因子，而25包含了2个
// 由规律看出一个数包含的因子数为该数可以连续整除5多少次
// 而(1...n)递增数列中的5因子个数为所有数5因子个数之和
func trailingZeroes(n int) int {
	num5 := 0
	for n > 1 {
		num5 += n / 5
		n /= 5
	}
	return num5
}

func TestTrailingZeroes(t *testing.T) {
	t.Log(trailingZeroes(30))
}

// 645. 错误的集合
// (1-n) 先找出重复的数字，再找出缺失的数字
// 额
func findErrorNums(nums []int) []int {
	length := len(nums)
	marked := make([]bool, length+1)
	duplicate := 0
	missing := 0
	sum := 0
	for _, num := range nums {
		sum += num
		if !marked[num] {
			marked[num] = true
		} else {
			duplicate = num
		}
	}
	missing = duplicate - (sum - ((length+1)*length)/2)
	return []int{duplicate, missing}
}

func TestFindErrorNums(t *testing.T) {
	t.Log(findErrorNums([]int{1, 2, 2, 4}))
}

//1583. 统计不开心的朋友

func unhappyFriends(n int, preferences [][]int, pairs [][]int) int {
	count := 0
	pairsMap := make(map[int]int)

	for _, pair := range pairs {
		pairsMap[pair[0]] = pair[1]
		pairsMap[pair[1]] = pair[0]
	}
	PreferencesIndex := make([][]int, n)
	for i, preference := range preferences {
		PreferencesIndex[i] = make([]int, n)
		for index, friend := range preference {
			PreferencesIndex[i][friend] = n - 1 - index
		}
	}
	for a, b := range pairsMap {
		for _, betterFriend := range preferences[a] {
			//break 之前的朋友都是关系好过 b 的朋友
			if betterFriend == b {
				break
			}
			ABetterFriendPair := pairsMap[betterFriend]
			if PreferencesIndex[betterFriend][a] > PreferencesIndex[betterFriend][ABetterFriendPair] {
				count += 1
				// 这里必须添加因为一个不开心的朋友可以不开心很多次
				break
			}
		}
	}
	return count
}

func TestUnhappyFriends(t *testing.T) {
	suit := struct {
		N           int
		Preferences [][]int
		Pairs       [][]int
	}{
		4,
		[][]int{[]int{1, 2, 3}, []int{3, 2, 0}, []int{3, 1, 0}, []int{1, 2, 0}},
		[][]int{[]int{0, 1}, []int{2, 3}},
	}
	t.Log(unhappyFriends(suit.N, suit.Preferences, suit.Pairs))
}

// 黑白画

func paintingPlan(n, k int) int {
	count := k / n
	rest := k % n
	C23 := 0
	if count == 0 {
		return 0
	}
	if rest == 0 {
		return 2 * C23
	}
	if rest == n-(count-1) {

	}
	return 0
}
