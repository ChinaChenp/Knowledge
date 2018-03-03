package main

import "fmt"

func parseString(item interface{}, errno int, required bool) string {
	if (item == nil || item == "") && !required {
		return ""
	}
	if s, ok := item.(string); ok {
		if s != "" || !required {
			return s
		}
	}
	panic(errno)
}

func main() {
	var v1 interface{} = ""
	re := parseString(v1, 34, false)
	fmt.Println(re)
}
