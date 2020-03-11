package error_test

import (
	"fmt"
	"net/http"
	"testing"
)

func Test_ErrorHandle(t *testing.T) {
	type Result struct {
		err      error
		response *http.Response
	}
	statucCheck := func(urls []string) <-chan Result {
		resultChan := make(chan Result)
		go func() {
			defer close(resultChan)
			for _, url := range urls {
				resp, err := http.Get(url)
				resultChan <- Result{
					err:      err,
					response: resp,
				}
			}
		}()
		return resultChan
	}
	urls := []string{
		"https://www.baidu.com",
		"https://cn.bing.com",
	}
	for result := range statucCheck(urls) {
		if result.err != nil {
			fmt.Printf("something wrong:%s \n", result.err)
			continue
		}
		fmt.Printf("response is %v\n", result.response)
	}
}
