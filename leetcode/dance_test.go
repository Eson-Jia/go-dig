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
func TestLengthOfLongestSubstring(t *testing.T) {
	theSuits := []struct {
		input  string
		output int
	}{
		{
			input:  "abcabcbb",
			output: 3,
		},
		{
			input:  "bbbbb",
			output: 1,
		},
		{
			input:  "pwwkew",
			output: 3,
		},
		{
			input:  " ",
			output: 1,
		},
		{
			input:  "dvdf",
			output: 3,
		},
		{
			input:  "",
			output: 0,
		},
		{
			input:  "cdd",
			output: 2,
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
				if loopIndex-1-begin+1 > max {
					max = loopIndex - 1 - begin + 1
				}
				begin = charIndex + 1
				for theChar, position := range cache {
					if position < begin {
						delete(cache, theChar)
					}
				}
			}
			cache[theChar] = loopIndex
		}
		if loopIndex-begin+1 > max && len(s) > 0 {
			return loopIndex - begin + 1
		}
		return max
	}
	for _, suit := range theSuits {
		if theFunc(suit.input) != suit.output {
			log.Fatalln("input:", suit.input, "output:", theFunc(suit.input), "expect:", suit.output)
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
			log.Fatalln("input:", suit.Input, "output:", theFunc(suit.Input), "output:", suit.Output)
		}
	}
}
