package main

import (
	"encoding/json"
	"fmt"
)

type Result struct {
	Uid *json.Number `json:"uid"`
	//ReturnLoc *int         `json:"returnloc"`
}

func CheckUid(uid **json.Number, required bool, def string) bool {
	if *uid == nil {
		if required {
			fmt.Printf("1\n")
			return false
		} else {
			re := json.Number(def)
			*uid = &re
			fmt.Printf("2\n")
			return true
		}
	} else {
		_, err := (*uid).Int64()
		if err != nil {
			re := json.Number(def)
			*uid = &re
			fmt.Printf("3\n")
			return false
		}

		//fmt.Printf("ok----------=%v==\n", (*uid).String())
		//out := json.Number((*uid).String())
		//*uid = &out
		fmt.Printf("4\n")
		return true
	}

	return false
}

func main() {
	m := `{"uid":2222222222222222}`

	fuck := &Result{}
	if err := json.Unmarshal([]byte(m), fuck); err != nil {
		fmt.Printf("unmarshal result failed, err:%v", err)
		return
	}
	fmt.Printf("Unmarshal %v\n", fuck)

	//var ptr *json.Number
	ok := CheckUid(&fuck.Uid, false, "-1")
	fm1, _ := fuck.Uid.Int64()
	fmt.Printf("out ----%v---===%d===-\n", ok, fm1)

	//re1 := json.Number("-5")
	//fm1, _ = re1.Int64()
	//fmt.Println(fm1)
	//ptr = nil
	//ok = CheckUid(&ptr, true, "1")
	//fmt.Printf("out ----%v--%v-\n", ok, ptr)

	//ptr = new(json.Number)
	//ok = CheckUid(&ptr, true, "2")
	//fmt.Printf("out ----%v--%v-\n", ok, ptr)

	//re1 := json.Number("222222222222222222222222222222222")
	//ptr = &re1
	//ok = CheckUid(&ptr, true, "3")
	//fmt.Printf("out ----%v--%v-\n", ok, ptr)

	//re1 = json.Number("986921652806590222")
	//ptr = &re1
	//ok = CheckUid(&ptr, true, "5")
	//fm, _ := ptr.Int64()
	//fmt.Printf("out ----%v--%v-\n", ok, fm)
}
