package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Data struct {
	B int64   `json:"b"`
	C float64 `json:"c"`
}

type Info struct {
	A    int         `json:"a"`
	Data interface{} `json:"data"`
}

func JsonDecoder(b []byte, out interface{}) error {
	decoder := json.NewDecoder(bytes.NewBuffer(b))
	decoder.UseNumber()

	err := decoder.Decode(out)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	data := Data{B: 1152921610250748151, C: 1152921610250748151}
	info := Info{A: 1, Data: data}

	b, _ := json.Marshal(info)

	out := Info{}
	err := JsonDecoder(b, &out)
	if err != nil {
		return
	}

	fmt.Println(out)

}
