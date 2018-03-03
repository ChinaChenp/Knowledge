package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{3, 5, 4, -1, 9, 11, -14}
	sort.Ints(a)
	fmt.Println(a)

	ss := []string{"surface", "ipad", "mac-pro", "mac-air", "think-pad", "idea-pad"}
	sort.Strings(sort.StringSlice(ss))
	fmt.Println(ss)
	ss1 := []string{"中國", "印尼", "印度", "日本", "香港", "哥倫比亞", "泰國", "馬來西亞", "智利", "朝鮮", "菲律賓", "新加坡", "新西蘭", "墨西哥", "澳大利亞"}
	sort.Strings(ss1)
	fmt.Println(ss1)
	//strslice := []string{"中国", "美国", "美人", "中国人"}
	strslice := []string{"山东", "山西", "安徽", "北京", "福建", "甘肃", "广东", "贵州"}
	//sort.Sort(sort.StringSlice(strslice))
	sort.Strings(strslice)
	fmt.Println(strslice)
	//sort.Sort(sort.Reverse(sort.StringSlice(ss)))
	//fmt.Printf("After reverse: %v\n", ss)
}
