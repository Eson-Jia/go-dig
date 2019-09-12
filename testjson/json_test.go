package testjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type TestStruct struct {
	FieldA string `json:"FieldA"`
	FieldB string
}

func TestJsonUnMarshal(t *testing.T) {
	str := []byte("{\"fielda\":\"test\"}")
	var test TestStruct
	err := json.Unmarshal(str, &test)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(test.FieldA)
}

func TestJosnMarshal(t *testing.T) {

}
