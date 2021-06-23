package dance

import (
	"fmt"
	"math"
	"sort"
	"testing"
)

func reverseWords(s string) string {
	sByte := []byte(s)
	out := make([]byte, 0)
	beginIndex := 0
	previousSpace := false
	for i := 0; i < len(sByte); i++ {
		if sByte[i] == ' ' {
			for j := i - 1; i >= 0 && j >= beginIndex; j-- {
				out = append(out, sByte[j])
			}
			out = append(out, ' ')
			previousSpace = true
			beginIndex = 0
		} else {
			if previousSpace {
				beginIndex = i
				previousSpace = false
			}
		}
	}
	for j := len(sByte) - 1; j >= beginIndex; j-- {
		out = append(out, sByte[j])
	}
	return string(out)
}

//16. 最接近的三数之和
func threeSumClosest(nums []int, target int) int {
	length := len(nums)
	closestSum := nums[0] + nums[1] + nums[2]
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			for k := j + 1; k < length; k++ {
				if sum := nums[i] + nums[j] + nums[k]; abs(sum-target) < abs(closestSum-target) {
					closestSum = sum
				}
			}
		}
	}
	return closestSum
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}
func TestTreeSumClosest(t *testing.T) {
	threeSumClosest([]int{-1, 2, 1, -4}, 1)
}

//237. 删除链表中的节点
func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

//148. 排序链表
//对链表进行排序，而且要求常量级空间和 n log n 级的空间复杂度。
//链表排序与数组排序不同的是数组的游标可以左右移动，但是链表只有从头部向尾部移动
//有了这个限制，归并排序和快排这些有双向游标的排序算法就无法发挥作用了
//常量级空间的要求也决定了无法将数据读到数组中再进行排序
//现在使用排除法，先列出时间复杂度为 n log n 的排序算法
//归并排序，希尔排序，快速排序，堆排序

//看到相似题目中有合并两个有序链表，突然来了灵感
//我们可以将链表像拉链一样分开成两个子链表，然后再对子链表进行同样的操作，直到每个链表只剩一个元素，然后再将这些链表归并排序

func sortList(head *ListNode) *ListNode {
	next := head
	var first, firstLast, second, secondLast *ListNode = nil, nil, nil, nil
	position := 0
	for next != nil {
		if position%2 == 0 {
			if position == 0 {
				first = next
			} else {
				firstLast.Next = next
			}
			firstLast = next
		} else {
			if position == 1 {
				second = next
			} else {
				secondLast.Next = next
			}
			secondLast = next
		}
		position += 1
		next = next.Next
	}
	if firstLast != nil {
		firstLast.Next = nil
	}
	if secondLast != nil {
		secondLast.Next = nil
	}
	if position > 2 {
		first = sortList(first)
		second = sortList(second)
	}
	var sorted, sortedLast *ListNode = nil, nil
	for first != nil || second != nil {
		var currentMin *ListNode = nil
		if first == nil {
			currentMin = second
			second = second.Next
		} else if second == nil {
			currentMin = first
			first = first.Next
		} else {
			if first.Val > second.Val {
				currentMin = second
				second = second.Next
			} else {
				currentMin = first
				first = first.Next
			}
		}
		if sorted == nil {
			sorted = currentMin
		} else {
			sortedLast.Next = currentMin
		}
		sortedLast = currentMin
	}
	return sorted
}

func TestSortList(t *testing.T) {
	suits := [][]int{
		{4, 2, 1, 3},
		{-1, 5, 3, 4, 0},
	}
	head := constructorList(suits[1])
	sortedHead := sortList(head)
	fmt.Println(sortedHead)
}

//78. 子集
//这是一个 2的n次幂的问题，每一个字符在子集中有出现或者不出现两个选项，所以有 2 的 n 次幂
func subsets(nums []int) [][]int {
	ret := make([][]int, 0)
	length := len(nums)
	lengthPow := int(math.Pow(2, float64(length)))
	for i := 0; i < lengthPow; i++ {
		array := make([]int, 0)
		for j := 0; j < length; j++ {
			if ((1 << uint(j)) & i) > 0 {
				array = append(array, nums[j])
			}
		}
		ret = append(ret, array)
	}
	return ret
}

//78. 子集

func subsetsSecond(nums []int) [][]int {
	return nil
}

//238. 除自身以外数组的乘积
func productExceptSelf(nums []int) []int {
	return nil
}

//136. 只出现一次的数字
//将所有数字与非到一起
func singleNumber(nums []int) int {
	for i := 1; i < len(nums); i++ {
		nums[0] ^= nums[i]
	}
	return nums[0]
}

//61. 旋转链表
// 先遍历一边获取长度 length  k = k % length，
// 将链表首尾相连，将 head 指针 向右移动 length - k 个位置
func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 {
		return head
	}
	node, latest, length := head, head, 0
	for node != nil {
		latest = node
		length += 1
		node = node.Next
	}
	//注意这里，如果长度为 0 就直接返回 nil
	if length == 0 {
		return nil
	}
	k %= length
	if k == 0 {
		return head
	}
	latest.Next = head
	for i := 0; i < length-k; i++ {
		latest = head
		head = head.Next
	}
	latest.Next = nil
	return head
}

func TestRotateRight(t *testing.T) {
	head := constructorList([]int{1, 2, 3, 4, 5})
	ret := rotateRight(head, 2)
	fmt.Println(ret)
}

//15. 三数之和
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	length := len(nums)
	if length < 3 {
		return nil
	}
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			//third := 0 - (nums[i] + nums[j])
		}
	}
	return nil
}

//11. 盛最多水的容器
//2020年10月14日09:46:06
//感觉动态规划在这并不适用，因为没有找到最优子结构，那么同样贪心算法也不适用。
//那么回溯算法呢，回溯算法可以用来测试某个选择是否适合，却不能用来寻找最优解，所以回溯算法也不适用。
//暴力尝试的方法
func maxArea(height []int) int {
	length := len(height)
	max := 0
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if theArea := (j - i) * getMin(height[i], height[j]); theArea > max {
				max = theArea
			}
		}
	}
	return max
}

//双指针解法
// 好吧，官方提示使用双指针
//我的思路是双指针分别指向数组的两端，计算出一个初始面积和左右指针两者的较大值，判断较大值在左边还是在右边
//然后移动相反方向的指针，如果该指针下一位大于该指针前一个值就计算这时的面积并与最大值比较
func maxAreaWithDoublePointer(height []int) int {
	length := len(height)
	left, right, max := 0, length-1, (length-1)*getMin(height[0], height[length-1])
	for left != right {
		if height[left] < height[right] {
			for before := height[left]; height[left] <= before; left++ {
				if right == left {
					return max
				}
			}
		} else {
			for before := height[right]; height[right] <= before; right-- {
				if right == left {
					return max
				}
			}
		}
		if theArea := (right - left) * getMin(height[left], height[right]); theArea > max {
			max = theArea
		}
	}
	return max
}

func maxAreaRegular(height []int) int {
	length := len(height)
	left, right, max := 0, length-1, 0
	for left != right {
		l, r, area := height[left], height[right], 0
		if l > r {
			area = (right - left) * r
			right--
		} else {
			area = (right - left) * l
			left++
		}
		if max < area {
			max = area
		}
	}
	return max
}

func TestMaxAreaWithDoublePointer(t *testing.T) {
	t.Log(maxAreaWithDoublePointer([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

//46. 全排列
// 这里要使用回溯算法

func permute(nums []int) [][]int {
	return nil
}

func dfs1(nums []int, left int) {
	for i := 0; i < left; i++ {

	}
}

func specialArray(nums []int) int {
	//sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		count := 0
		for j := 0; j < len(nums); j++ {
			if i < nums[j] {
				count += 1
			}
		}
		if len(nums)-count == i {
			return i
		}
	}
	return -1
}

func isEvenOddTree(root *TreeNode) bool {
	floor := 0
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	next := make([]*TreeNode, 0)
	isOdd := func(a int) bool { return a%2 == 1 }
	isEven := func(a int) bool { return a%2 == 0 }
	isLess := func(a, b int) bool { return a < b }
	isGreat := func(a, b int) bool { return a > b }
	for {
		var LessGreat func(int, int) bool
		var OddEven func(int) bool
		if floor%2 == 1 {
			OddEven = isEven
			LessGreat = isGreat

		} else {
			OddEven = isOdd
			LessGreat = isLess
		}
		previous, isFirst := 0, true
		for i := 0; i < len(queue); i++ {
			theVal := queue[i].Val
			if !OddEven(theVal) {
				return false
			}
			if isFirst {
				isFirst = false
			} else if !LessGreat(previous, theVal) {
				return false
			}
			if queue[i].Left != nil {
				next = append(next, queue[i].Left)
			}
			if queue[i].Right != nil {
				next = append(next, queue[i].Right)
			}
			previous = theVal
		}
		if len(next) == 0 {
			break
		}
		queue = next
		next = make([]*TreeNode, 0)
		floor += 1
	}
	return true
}

func TestOdd(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 10}
	root.Right = &TreeNode{Val: 4}
	root.Left.Left = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 7}
	root.Right.Right = &TreeNode{Val: 9}
	root.Left.Left.Left = &TreeNode{Val: 12}
	root.Left.Left.Right = &TreeNode{Val: 8}
	root.Right.Left.Left = &TreeNode{Val: 6}
	root.Right.Right.Right = &TreeNode{Val: 2}
	t.Log(isEvenOddTree(root))
}

//4. 寻找两个正序数组的中位数
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	return 0
}

//89. 格雷编码
//官方提示使用回溯算法
//需要注意的是，两个连续的数值仅有一个位数的差异，例如：000 和 011 是不合法的，而且必须以 0 开头，即开头必须是 0
//
func grayCode(n int) []int {
	buff := []int{0}
	for i := 1; i <= n; i++ {
		length := len(buff)
		for j := length - 1; j >= 0; j-- {
			v := buff[j] | (1 << (i - 1))
			buff = append(buff, v)
		}
	}
	return buff
}

func TestGrayCode(t *testing.T) {
	t.Log(grayCode(2))
}

func pow(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 2
	}
	if n%2 == 1 {
		return pow(n/2) * pow(n/2) * 2
	} else {
		return pow(n/2) * pow(n/2)
	}
}

//54. 螺旋矩阵
func spiralOrder(matrix [][]int) []int {
	return nil
}

//7. 整数反转
func reverse(x int) int {
	return 0
}

func TestBitPrint(t *testing.T) {
	t.Logf("%b", -2147483648)
	fmt.Printf("%b", -2147483648)
}

//23. 合并K个升序链表
//第一个思路是使用索引优先队列进行归并排序
//因为指针的特性，索引优先队列不是必须的，优先队列就可以满足需求了
//现在需要先实现优先队列
//优先队列使用的数据结构是堆，这里我们使用的是最小堆，堆的结构原理是完全二叉树
//如何初始化原始堆： 1.使用一个数组初始化; 2.连续使用 insert
//现在尝试使用连续 insert
type Dump struct {
	FirstPoints []*ListNode
	size        int
}

func New(maxSize int) *Dump {
	dump := &Dump{
		FirstPoints: make([]*ListNode, maxSize+1),
		size:        0,
	}
	return dump
}

func (d *Dump) less(i, j int) bool {
	return d.FirstPoints[i].Val < d.FirstPoints[j].Val
}

func (d *Dump) exchange(i, j int) {
	d.FirstPoints[i], d.FirstPoints[j] = d.FirstPoints[j], d.FirstPoints[i]
}

func (d *Dump) swim(k int) {
	for k >= 2 && d.less(k, k/2) {
		d.exchange(k, k/2)
		k /= 2
	}
}

func (d *Dump) sink(k int) {
	for k*2 <= d.size {
		min := k * 2
		if k*2+1 <= d.size && d.less(k*2+1, k*2) {
			min = k*2 + 1
		}
		if d.less(min, k) {
			d.exchange(min, k)
			k = min
		} else {
			break
		}
	}
	//fmt.Print("sink d:")
	//for i := 1; i <= d.size; i++ {
	//	fmt.Printf(" %d", d.FirstPoints[i].Val)
	//}
	//fmt.Print("\n")
}

func (d *Dump) Pop() *ListNode {
	min := d.FirstPoints[1]
	d.FirstPoints[1] = d.FirstPoints[d.size]
	d.size--
	d.sink(1)
	return min
}

func (d *Dump) Size() int {
	return d.size
}

//Insert 将数据放入数组尾部，并
func (d *Dump) Insert(node *ListNode) {
	d.size += 1
	d.FirstPoints[d.size] = node
	d.swim(d.size)
	//fmt.Print("swim d:")
	//for i := 1; i <= d.size; i++ {
	//	fmt.Printf(" %d", d.FirstPoints[i].Val)
	//}
	//fmt.Print("\n")
}

func mergeKLists(lists []*ListNode) *ListNode {
	var newListHead, newListTail *ListNode
	minDump := New(len(lists))
	for _, list := range lists {
		if list != nil {
			minDump.Insert(list)
		}
	}
	for minDump.Size() > 0 {
		min := minDump.Pop()
		if min.Next != nil {
			minDump.Insert(min.Next)
		}
		if newListHead == nil {
			newListHead = min
			newListTail = min
		} else {
			newListTail.Next = min
			newListTail = min
		}
	}
	if newListTail != nil {
		newListTail.Next = nil
	}
	return newListHead
}

func TestMergeKLists(t *testing.T) {
	theLists := []*ListNode{
		//constructorList([]int{-8, -7, -7, -5, 1, 1, 3, 4}),
		//constructorList([]int{-2}),
		//constructorList([]int{-10, -10, -7, 0, 1, 3}),
		//constructorList([]int{2}),
	}
	result := mergeKLists(theLists)

	for result != nil {
		t.Log(result.Val)
		result = result.Next
	}
}

//977. 有序数组的平方
func sortedSquares(A []int) []int {
	//index, v := 0, 0
	//for index, v = range A {
	//	if v > 0 {
	//		break
	//	}
	//}
	//for i := 0; i < len(A)-1; i++ {
	//	A[i] = A[i] * A[i]
	//}
	//newArray := make([]int, len(A))
	//
	//i, j, count := index-1, index, 0; i >= 0 || j < len(A)
	return nil
}

//236. 二叉树的最近公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	pStack, qStack := prefixOrder(root, p), prefixOrder(root, q)
	pLength, qLength := len(pStack), len(qStack)
	var ret *TreeNode
	for pIndex, qIndex := pLength-1, qLength-1; pIndex >= 0 && qIndex >= 0 && pStack[pIndex] == qStack[qIndex]; {
		ret = pStack[pIndex]
		pIndex -= 1
		qIndex += 1
	}
	return ret
}

func prefixOrder(root, p *TreeNode) []*TreeNode {
	if root == nil {
		return nil
	}
	if root == p {
		return []*TreeNode{root}
	}
	if stack := prefixOrder(root.Left, p); stack != nil {
		return append(stack, root)
	}
	if stack := prefixOrder(root.Right, p); stack != nil {
		return append(stack, root)
	}
	return nil
}

// 官方递归
func lowestCommonAncestorRegular(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left, right := lowestCommonAncestorRegular(root.Left, p, q), lowestCommonAncestorRegular(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if root == p || root == q {
		return root
	}
	if left == nil {
		return right
	}
	return left
}
