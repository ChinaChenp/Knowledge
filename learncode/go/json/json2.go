package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//str := `{"errno": 0, "errmsg":"SUCCESS","data": {"aDriverBasic": {"name": "阿斯蒂芬","id_no": "130421199601060257"}}}`
	str := `{"errno": 0, "errmsg":"SUCCESS","data": {"aDriverBasic": []}}`

	type Result struct {
		Errno  int64  `json:"errno"`
		Errmsg string `json:"errmsg"`
		Data   struct {
			AdriverBasic struct {
				IdNo string `json:"id_no"`
				Name string `json:"name"`
			}
		}
	}
	fuck := &Result{}
	if err := json.Unmarshal([]byte(str), fuck); err != nil {
		fmt.Printf("unmarshal result failed, err:%v", err)
		return
	}
	fmt.Printf("unmarshal %v\n", fuck)
}
