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

func TestRange(t *testing.T) {
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
