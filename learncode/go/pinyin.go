package main

import (
	"fmt"
	"sort"

	"github.com/mozillazg/go-pinyin"
)

func merge(pin pinyin.Args, s string) string {
	ss := pinyin.Pinyin(s, pin)

	re := ""
	for _, v := range ss {
		re += v[0]
	}
	return re
}

func main() {
	//hans := "中国人"
	a := pinyin.NewArgs()
	//a.Style = pinyin.FirstLetter

	//fmt.Println(merge(pinyin.Pinyin(hans, a)))
	hans := "中国人"
	fmt.Println("default:", pinyin.Pinyin(hans, a))

	strslice := []string{"山东", "成安", "山西", "安徽", "安庆", "北京", "福建", "甘肃", "广东", "贵州", "成都", "重庆", "西安"}

	kvs := make(map[string]string)
	newslice := make([]string, len(strslice))
	for i, v := range strslice {
		gbk := merge(a, v)
		fmt.Printf("%s->%s, ", v, gbk)
		kvs[gbk] = v
		newslice[i] = gbk
	}

	sort.Strings(newslice)
	for _, v := range newslice {
		fmt.Printf("%s, ", kvs[v])
	}
	fmt.Println()
}
