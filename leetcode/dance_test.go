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
		var max, current int
		cache := make(map[int32]struct{}, 0)
		for _, b := range s {
			if _, ok := cache[b]; ok {
				if current > max {
					max = current
				}
				current = 0
				cache = make(map[int32]struct{}, 0)
			}
			current += 1
			cache[b] = struct{}{}
		}
		if max < current {
			return current
		}
		return max
	}
	for _, suit := range theSuits {
		if theFunc(suit.origin) != suit.result {
			log.Fatalln(suit.origin, theFunc(suit.origin), suit.result)
		}
	}
	log.Print()
}
