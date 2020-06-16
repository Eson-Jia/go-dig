package data_struct

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"testing"
	"unicode"
	"unicode/utf8"
)

func TestSliceAsKey(t *testing.T) {
	var s SliceMap = make(map[string]int)
	suits := []struct {
		List   []string
		Result int
	}{
		{
			List:   []string{"1", "2", "34"},
			Result: 3,
		},
		{
			List:   []string{"1", "3", "34"},
			Result: 3,
		},
		{
			List:   []string{"this is a ", "d", "34"},
			Result: 3,
		},
	}
	for _, suit := range suits {
		for i := 0; i < suit.Result; i++ {
			s.Add(suit.List)
		}
	}
	for _, suit := range suits {
		if result := s.Count(suit.List); suit.Result != result {
			t.Fatalf("List:%q, result: %d\n", suit.List, result)
		}
	}
}
func k(list []string) string {
	return fmt.Sprintf("%q", list)
}

type SliceMap map[string]int

func (s SliceMap) Add(list []string) { s[k(list)]++ }

func (s SliceMap) Count(list []string) int { return s[k(list)] }

func TestCount(t *testing.T) {
	counts := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int
	invalid := 0
	in := bufio.NewReader(os.Stdin)
	for {
		// ReadRune 读取一个 UTF-8 编码的 Unicode 字符,并返回该字符的字节数.
		// 如果不是有效的 utf-8 编码,那么会消耗一个字节并返回 unicode.ReplacementChar 字符
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		// 不是有效的 utf-8 编码
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
