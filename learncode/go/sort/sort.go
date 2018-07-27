package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"sort"
	"unicode/utf8"

	"golang.org/x/text/collate"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/language"
	"golang.org/x/text/transform"
)

func GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func Utf8ToGbk(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func main() {
	c := collate.New(language.Chinese)
	strings := []string{"山东", "山西", "安徽", "安庆", "北京", "福建", "甘肃", "广东", "贵州", "成都", "重庆", "西安"}
	c.SortStrings(strings)
	fmt.Println(strings)

	c1 := collate.New(language.TraditionalChinese)
	ss1 := []string{"中國", "印尼", "印度", "日本", "香港", "哥倫比亞", "泰國", "馬來西亞", "智利", "朝鮮", "菲律賓", "新加坡", "新西蘭", "墨西哥", "澳大利亞"}
	c1.SortStrings(ss1)
	fmt.Println(ss1)

	c2 := collate.New(language.English)
	ss2 := []string{"surface", "ipad", "mac-pro", "mac-air", "think-pad", "idea-pad"}
	c2.SortStrings(ss2)
	fmt.Println(ss2)
	//re := collate.OptionsFromTag(language.SimplifiedChinese)
	//fmt.Printf("%v", *re)

	//a := []int{3, 5, 4, -1, 9, 11, -14}
	//sort.Ints(a)
	//fmt.Println(a)

	//ss := []string{"surface", "ipad", "mac-pro", "mac-air", "think-pad", "idea-pad"}
	//sort.Strings(sort.StringSlice(ss))
	//fmt.Println(ss)
	//ss1 := []string{"中國", "印尼", "印度", "日本", "香港", "哥倫比亞", "泰國", "馬來西亞", "智利", "朝鮮", "菲律賓", "新加坡", "新西蘭", "墨西哥", "澳大利亞"}
	//sort.Strings(ss1)
	//fmt.Println(ss1)
	//strslice := []string{"中国", "美国", "美人", "中国人"}
	strslice := []string{"山东", "山西", "安徽", "安庆", "北京", "福建", "甘肃", "广东", "贵州", "成都", "重庆", "西安"}

	kvs := make(map[string]string)
	newslice := make([]string, len(strslice))
	for i, v := range strslice {
		gbk, err := Utf8ToGbk([]byte(v))
		if err != nil {
			fmt.Printf("shit ~~~~~~~~~~~~ %s", v)
		}
		kvs[string(gbk)] = v
		newslice[i] = string(gbk)
	}

	sort.Strings(newslice)
	for _, v := range newslice {
		fmt.Printf("%s, ", kvs[v])
	}
	fmt.Println()
}

//sort.Sort(sort.Reverse(sort.StringSlice(ss)))
//fmt.Printf("After reverse: %v\n", ss)
func firstLetter(s string) string {
	_, size := utf8.DecodeRuneInString(s)
	return s[:size]
}
