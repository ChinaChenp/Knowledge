package main

import (
	"encoding/json"
	"fmt"
)

type EncodeMessage struct {
	MsgId   string `json:"msg_id"`
	Content string `json:"content"`
	Age     int    `json:"age"`
}

type DecodeMessage struct {
	MsgId   string  `json:"msg_id"`
	Content string  `json:"content"`
	Age     *int    `json:"age"`
	Name    *string `json:"name"`
}

func main() {
	var age int = 35
	encode := EncodeMessage{"01", "test", age}
	str, err := json.Marshal(encode)
	fmt.Println(string(str), err)

	var decode DecodeMessage
	err = json.Unmarshal(str, &decode)
	fmt.Println(decode)
	if decode.Age != nil {
		fmt.Println(*decode.Age)
	}

	if decode.Name != nil {
		fmt.Println(*decode.Name)
	}
}
