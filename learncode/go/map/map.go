package main

import "fmt"

type PersonInfo struct {
	Id   string
	Name string
}

func PrintMap(mp map[string]PersonInfo) {
	fmt.Println("-------------------\n")
	for k, v := range mp {
		fmt.Printf("key %s, value %v\n", k, v)
	}
}

func do1() (map[string]int, bool) {
	return map[string]int{"chenp1": 1, "xunbo": 2}, true
}

func do() (rep map[string]int, re bool) {
	rep, err := do1()
	return rep, err
}

func main() {
	mp := make(map[string]PersonInfo, 5)
	mp["1"] = PersonInfo{"11", "chenping"}
	mp["2"] = PersonInfo{"22", "gongyuan"}
	PrintMap(mp)

	if v, ok := mp["1"]; ok {
		fmt.Printf("-----------%v, %v\n", v, ok)
	} else {
		fmt.Printf("-----------%v, %v\n", v, ok)
	}

	delete(mp, "1")
	if _, ok := mp["1"]; ok {
		fmt.Println("find key 1\n")
	} else {
		fmt.Println("not find key 1\n")
	}

	mp1 := make(map[string]interface{})
	if mp1 == nil {
		fmt.Println("mp1 is nil")
	} else {
		fmt.Println("mp1 is not nil", mp1)
	}

	mp1 = nil
	if mp1 == nil {
		fmt.Println("mp1 is nil")
	} else {
		fmt.Println("mp1 is not nil", mp1)
	}

	if v, ok := mp1["chenping"]; ok {
		fmt.Println("mp1 chenping ok", v)
	} else {
		fmt.Println("mp1 chenping not ok")
	}

	var mp4 map[string]interface{}
	delete(mp4, "chenp")
	//mp4["chenp"] = "ping"
	if v, ok := mp4["chenping"]; ok {
		fmt.Println("not find chenping", v)
	} else {
		fmt.Println("not find chenping")
	}

	if mp4 == nil {
		fmt.Println("mp4 is nil")
	} else {
		fmt.Println("mp4 is not nil", mp1)
	}

	re, err := do()
	fmt.Println(re, err)
	//mp1["chenping"] = "good" // will core

	map4 := make(map[string]interface{})
	map4["chenping"] = "chongqing"
	map4["age"] = "28"
	fmt.Printf("map not exist %v\n", map4["chenping1"])

	//ad := map4["chenping"].(string)
	//age := map4["age1"].(string)

	//fmt.Println(ad, age)

	// map sort
	var mp2 = map[int]string{
		1:  "11",
		2:  "22",
		8:  "88",
		5:  "55",
		3:  "11",
		4:  "22",
		11: "111",
		9:  "99",
		6:  "66",
	}
	for k, v := range mp2 {
		fmt.Println(k, v)
	}
}
