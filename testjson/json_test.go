package testjson

import (
	"encoding/json"
	"fmt"
	"strings"
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

func TestJsonOmitEmptyMarshal(t *testing.T) {
	data := struct {
		FiledA string `json:"fileda,omitempty"`
		FiledB string `json:"filedb"`
	}{}
	body, err := json.Marshal(data)
	if err != nil {
		t.Errorf("%s\n", err)
	}
	// FiledA有 omitempty tag 因为是零值,所以被省略了
	if strings.Compare(`{"filedb":""}`, string(body)) != 0 {
		t.Errorf("failed in compare\n")
	}
}
