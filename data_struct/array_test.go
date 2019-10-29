package data_struct

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	origin := make([]byte, 100)
	var model1 []byte
	//需要传入 model1 的指针,因为需要对其做修改, model1 的值最好为 nil
	n, err := fmt.Sscanf("123", "%s", &model1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(model1)
	fmt.Println(n, origin[:100])
}
