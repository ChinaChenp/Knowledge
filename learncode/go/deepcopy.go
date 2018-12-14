package main

import (
	"encoding/json"
	"fmt"
)

func Copy(dst interface{}, src interface{}) error {
	if dst == nil {
		return fmt.Errorf("dst cannot be nil")
	}
	if src == nil {
		return fmt.Errorf("src cannot be nil")
	}
	bytes, err := json.Marshal(src)
	if err != nil {
		return fmt.Errorf("Unable to marshal src: %s", err)
	}
	err = json.Unmarshal(bytes, dst)
	if err != nil {
		return fmt.Errorf("Unable to unmarshal into dst: %s", err)
	}
	return nil
}

type List struct {
	A int
	B float64
}

type Map struct {
	K int
	V string
}

type From struct {
	C []string
	L *List
	M *Map
}

func main() {
	from := &From{
		C: []string{"1", "2", "3", "4"},
		L: &List{
			A: 5,
			B: 2.34567,
		},
		M: &Map{
			K: 99,
			V: "99",
		},
	}

	to := &From{}
	Copy(to, from)
	fmt.Println(to, to.L, to.M)
}
