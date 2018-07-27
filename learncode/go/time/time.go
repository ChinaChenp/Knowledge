package main

import (
	"fmt"
	"strings"
)

func getMaxDuration() int {
	return 5
}

func forMatTime(t string) string {
	if t == "" {
		return ""
	}

	re := strings.Split(t, ".")
	if len(re) != 2 {
		return ""
	}

	return strings.Replace(re[0], "T", " ", -1)
}

func main() {
	//startTime := time.Now().Add(0 - time.Minute * time.Duration(getMaxDuration())).Format("2006-01-02 15:04:05")
	//fmt.Println(startTime)
	//fmt.Println(time.Now().Format("2006-01-02 15:04:05"))

	str := "2018-07-05T14:42:00.000+08:00"

	//t, _ := time.Parse("2006-01-02 15:04:05", "2018-07-05T14:43:00")
	t := forMatTime(str)
	fmt.Println(t)
}