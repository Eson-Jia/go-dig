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
// field tag 可以是任何字符串,但是它一般理解为由空格分割的key:"value" pair 组成的列表,因为这些 pair 包含双引号,
// 所以 field tags 一般都写成原始字符串
// 一般是这种格式
// FieldA int `json:"field_a,omitempty" bson:"fieldA,omitempty"`
// 这里我们 FieldA 的tag包含了两个键json,bson以及分别对用的值:"field_a,omitempty"和"fieldA,omit"

// 结构体成员声明后面跟的字符串是 成员标签定义 ,是由一串由空格分开的标签键值对 (key:"value") 组成的
// json 成员标签有额外的选项 omitempty 和 string

type Test struct {
	TestString bool `json:"test_bool,string"`
	TestNumber bool  `json:"test_number"`
	TestFieldTag bool "json:\"test_field_tag\""
}

func TestTest(t *testing.T) {
	the := Test{
		TestString: false,
		TestNumber: false,
	}
	buff,err:= json.Marshal(the)
	if err!=nil{
		t.Fatal(err)
	}
	t.Logf("json:%s",string(buff))
}

