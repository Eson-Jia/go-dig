package pearls

import (
	"fmt"
	"net/http"
	"testing"
)

var buffer = make([]byte, 100000)
var word = make([][]byte, 2000)

var k = 2

func wordCompare(q, p []byte) int {
	n, index := k, 0
	for q[index] == p[index] {
		if q[index] == 0 && n > 0 {
			n--
		}
		index++
		if n == 0 {
			return 0
		}
	}
	return int(q[index] - p[index])
}

func skip(phrase []byte, n int) {
	index := 0
	for n > 0 {
		if phrase[index] == 0 {
			n--
		}
		index++
	}
}

func TestMakov(t *testing.T) {
	i := 0
	word[i] = buffer
	for err := error(nil); err == nil; _, err = fmt.Scanf("%s", word[i]) {
		word[i] = buffer[len(word[i-1])]
		fmt.Print(word[i])
	}
}

func TestHttp(t *testing.T) {
	http.ListenAndServe()
}
