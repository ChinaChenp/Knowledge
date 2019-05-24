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

func JsonEncoder(in interface{}) ([]byte, error) {
	var out bytes.Buffer
	encoder := json.NewEncoder(&out)
	err := encoder.Encode(in)
	if err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}

func main() {
	data := Data{B: 5764607589339758604, C: 1152921610250748151}
	var in interface{}
	in = data

	b, _ := json.Marshal(in)

	out := Data{}
	//err := JsonDecoder(b, &out)
	err := json.Unmarshal(b, &out)
	if err != nil {
		return
	}

	fmt.Println(out, out.B)

}
