package main

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	MsgId   string `json:"msg_id"`
	Content string `json:"content"`
}

type Message1 struct {
	MsgId   string `json:"msg_id"`
	Content string `json:"content"`
	Age     string `json:"age"`
}

//json 序列号反序列化
func main() {
	msg := Message{"msgid_001", "contente2222222222"}
	str, err := json.Marshal(msg)
	fmt.Println(string(str), err)

	var msg1 Message1
	err = json.Unmarshal(str, &msg1)
	fmt.Println(msg1)
}
