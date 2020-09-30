package dance_test

import (
	"fmt"
	"testing"
)

// 1593. 拆分字符串使唯一子字符串的数目最大
// 在切分的时候需要考虑切分之后剩余的字符串是否合法
// 例如 aa 如果我第一个子字符串选为 a，那么下一个子字符串也只能是 a，这个切分方式是不合法的。
// 下面的代码在输入为 addbsd 的时候就不能正确的切分

func maxUniqueSplit(s string) int {
	sByte := []byte(s)
	record := make(map[string]struct{})
	lastCut := 0
	count := 0
	for i := 0; i < len(sByte); i++ {
		if _, ok := record[string(sByte[lastCut:i+1])]; !ok {
			record[string(sByte[lastCut:i+1])] = struct{}{}
			lastCut = i + 1
			count += 1
		}
	}
	return count
}

// 1593. 拆分字符串使唯一子字符串的数目最大
// 采用回溯算法
func maxUniqueSplitSecond(s string) int {
	sByte := []byte(s)
	record := make(map[string]struct{})
	setRecord := func(key []byte) {
		record[string(key)] = struct{}{}
	}
	deleteRecord := func(key []byte) {
		delete(record, string(key))
	}
	cutSlice := []int{0}
	lastCut, i := 0, 0
	length := len(sByte)
	for i < length && lastCut < length {
		if _, ok := record[string(sByte[lastCut:i+1])]; !ok {
			setRecord(sByte[lastCut : i+1])
			lastCut = i + 1
			cutSlice = append(cutSlice, lastCut)
			i += 1
		} else if i == length-1 {
			cutSliceLength := len(cutSlice)
			c1, c2 := cutSliceLength-2, cutSliceLength-1
			deleteRecord(sByte[c1:c2])
			i = lastCut + 1
			lastCut = c2
		} else {
			i += 1
		}
	}
	previous := 0
	for _, cut := range cutSlice {
		if cut != previous {
			fmt.Println(string(sByte[previous:cut]))
			previous = cut
		}
	}
	return 0

}

func TestMaxUniqueSplit(t *testing.T) {
	suits := []struct {
		Input  string
		Result int
	}{
		{
			"addbsd",
			5,
		},
	}
	for _, suit := range suits {
		if result := maxUniqueSplitSecond(suit.Input); result != suit.Result {
			t.Errorf("expect %x,got %x\n", suit.Result, result)
		}
	}
}
