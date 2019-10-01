package programingpears

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"
)

func RandBigNum() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	res := r.Int()
	return res
}

func RandRangeInt(i, j int) int {
	if i == j {
		return i
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	res := r.Intn(j - i + 1)
	return res + i
}

// GenerateSelectArray1 在 [0...n-1] 之间挑出 m 个随机不重复且有序的数
func GenerateSelectArray1(m, n int) []int {
	var selected []int
	for i := 0; i < n; i++ {
		if (RandBigNum() % (n - i)) < m {
			selected = append(selected, i)
			m--
		}
	}
	return selected
}

// shuffledArray 重新洗牌有序的队列使其变得无序
func shuffledArray(source []int) []int {
	var swap = func(i, j int) {
		source[i], source[j] = source[j], source[i]
	}
	length := len(source)
	for i := 0; i < length; i++ {
		swap(i, RandRangeInt(i, length-1))
	}
	return source
}

func GenerateSelectArray2(m, n int) []int {
	theOrigin := make([]int, 0)
	for i := 0; i < n; i++ {
		theOrigin = append(theOrigin, i)
	}
	// shuffleshuffledArray 使队列变得无序
	theUnsorted := shuffledArray(theOrigin)
	//在无序队列中挑选前 m 个
	theUnsorted = theUnsorted[:m]
	// 对 m 个元素进行排序
	sort.Slice(theUnsorted, func(i, j int) bool { return theUnsorted[i] < theUnsorted[j] })
	return theUnsorted
}

func GenerateSelectArray3(m, n int) []int {
	origin := make([]int, 0)
	for i := 0; i < n; i++ {
		origin = append(origin, i)
	}
	for i := 0; i < m; i++ {
		j := RandRangeInt(i, n-1)
		origin[i], origin[j] = origin[j], origin[i]
	}
	selected := origin[:m]
	sort.Slice(selected, func(i, j int) bool { return selected[i] < selected[j] })
	return selected
}

func testGenerateSelectArray(theAlgo func(m, n int) []int) map[int]uint {
	statistics := make(map[int]uint)
	for i := 0; i < 1000; i++ {
		res := theAlgo(5, 10)
		if !sort.SliceIsSorted(res, func(i, j int) bool { return res[i] < res[j] }) {
			panic("the slice is unsorted")
		}
		// fmt.Println("result:", res)
		for _, resultElement := range res {
			statistics[resultElement]++
		}
	}
	return statistics
}

func TestGenerateSelectArray1(t *testing.T) {
	statistics := testGenerateSelectArray(GenerateSelectArray1)
	var sum uint
	for k, v := range statistics {
		fmt.Printf("%v-->%v\n", k, v)
		sum += v
	}
	fmt.Println("sum:", sum)
}

func TestGenerateSelectArray2(t *testing.T) {
	statistics := testGenerateSelectArray(GenerateSelectArray2)
	var sum uint
	for k, v := range statistics {
		fmt.Printf("%v-->%v\n", k, v)
		sum += v
	}
	fmt.Println("sum:", sum)
}

func TestGenerateSelectArray3(t *testing.T) {
	statistics := testGenerateSelectArray(GenerateSelectArray3)
	var sum uint
	for k, v := range statistics {
		fmt.Printf("%v-->%v\n", k, v)
		sum += v
	}
	fmt.Println("sum:", sum)
}
