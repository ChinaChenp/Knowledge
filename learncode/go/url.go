package main

import "fmt"

type Values map[string]string

func (v *Values) Get(key string) string {
	if _, ok := (*v)[key]; ok {
		return (*v)[key]
	}
	return ""
}

func (v Values) Add(key, value string) {
	v[key] = value
}

func main() {
	url := &Values{"a": "aa", "b": "bb"}
	fmt.Println(url.Get("a"))
	fmt.Println(url.Get("b"))
	url.Add("c", "cc")
	fmt.Println(url.Get("c"))
	fmt.Println(url.Get("d"))
}
