package pearls

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"sort"
	"testing"
)

func TestSortSlice(t *testing.T) {
	a := []byte("post")
	sort.Slice(a, func(x, y int) bool { return a[x] < a[y] })
	fmt.Printf("%s\n", a)
}

func TestFindTransform(t *testing.T) {
	resp, err := http.Get("https://raw.githubusercontent.com/dwyl/english-words/master/words_alpha.txt")
	if err != nil {
		t.Errorf("http get error:%s", err)
		return
	}
	defer resp.Body.Close()
	read := bufio.NewReader(resp.Body)
	theColleciton := make(map[string][]string)
	theKeys := make([]string, 0)
	for {
		var line []byte
		buff, isPrefix, err := read.ReadLine()
		// fmt.Printf("%s\n", buff)
		if err == io.EOF {
			break
		}
		line = append(line, buff...)
		if !isPrefix {
			data := string(line)
			sort.Slice(line, func(i, j int) bool { return line[i] < line[j] })
			sig := string(line)
			theColleciton[sig] = append(theColleciton[sig], data)
			theKeys = append(theKeys, sig)
			line = line[0:0]
		}
	}
	sort.Slice(theKeys, func(i, j int) bool {
		return len(theKeys[i]) < len(theKeys[j])
	})
	for _, v := range theKeys {
		for _, alpha := range theColleciton[v] {
			fmt.Println(alpha)
		}
		fmt.Println("")
	}
}

func TestCopy(t *testing.T) {
	a := []byte("asdfd")
	var b []byte
	b = append(b, a...)
	fmt.Println(a, b)
	b = append(b, a...)
	fmt.Println(a, b)
}
