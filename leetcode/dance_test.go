package dance_test

import (
	"log"
	"testing"
)

func TestMapOK(t *testing.T) {
	cache := make(map[int32]int)
	value, ok := cache[100]
	log.Println(value, ok)
	if false {
		value, ok = cache[100]
		log.Println(value, ok)
	} else {
		value = cache[100]
		log.Println(value, ok)
	}
	cache[100] = 0
	value, ok = cache[100]
	log.Println(value, ok)
}

// TestLongestSubStr 无重复字符的最长子串
func TestLongestSubStr(t *testing.T) {
	theSuits := []struct {
		origin string
		result int
	}{
		{
			origin: "abcabcbb",
			result: 3,
		},
		{
			origin: "bbbbb",
			result: 1,
		},
		{
			origin: "pwwkew",
			result: 3,
		},
		{
			origin: " ",
			result: 1,
		},
		{
			origin: "dvdf",
			result: 3,
		},
	}
	theFunc := func(s string) int {
		var max int
		var begin int
		cache := make(map[int32]int, 0)
		var loopIndex int
		var theChar int32
		for loopIndex, theChar = range s {
			if charIndex, ok := cache[theChar]; ok {
				begin = charIndex
				if loopIndex-begin > max {
					max = loopIndex - begin
				}
				for theChar, position := range cache {
					if position < begin {
						delete(cache, theChar)
					}
				}
			}
			cache[theChar] = loopIndex
		}
		if loopIndex-begin > max {
			return loopIndex - begin
		}
		return max
	}
	for _, suit := range theSuits {
		if theFunc(suit.origin) != suit.result {
			log.Fatalln(suit.origin, "result:", theFunc(suit.origin), "expect:", suit.result)
		}
	}
	log.Print()
}

// TestLongestCommonPrefix 最长公共前缀
func TestLongestCommonPrefix(t *testing.T) {
	theFunc := func(strs []string) string {
		if len(strs) == 0 {
			return ""
		}
		common := strs[0]
		for _, str := range strs {
			str1, str2 := common, str
			buff := make([]byte, 0)
			if len(str1) > len(str2) {
				str1, str2 = str2, str1
			}
			for i := 0; i < len(str1); i++ {
				if str1[i] != str2[i] {
					break
				}
				buff = append(buff, str1[i])
			}
			common = string(buff)
		}
		return common
	}
	suits := []struct {
		Input  []string
		Output string
	}{
		{
			Input:  []string{"flower", "flow", "flight"},
			Output: "fl",
		},
		{
			Input:  []string{"dog", "racecar", "car"},
			Output: "",
		},
	}
	for _, suit := range suits {
		if theFunc(suit.Input) != suit.Output {
			log.Fatalln("input:", suit.Input, "result:", theFunc(suit.Input), "output:", suit.Output)
		}
	}
}
