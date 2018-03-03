package main

import (
	"bytes"
	"fmt"
	s "strings"
)

var p = fmt.Println

func main() {
	p("Contains: ", s.Contains("test", "es"))      //是否包含 true
	p("Count: ", s.Count("test", "t"))             //字符串出现字符的次数 2
	p("HasPrefix: ", s.HasPrefix("test", "te"))    //判断字符串首部 true
	p("HasSuffix: ", s.HasSuffix("test", "st"))    //判断字符串结尾 true
	p("Index: ", s.Index("test", "e"))             //查询字符串位置 1
	p("Join: ", s.Join([]string{"a", "b"}, "-"))   //字符串数组 连接 a-b
	p("Repeat: ", s.Repeat("a", 5))                //重复一个字符串 aaaaa
	p("Replace: ", s.Replace("foo", "o", "0", -1)) //字符串替换 指定起始位置为小于0,则全部替换 f00
	p("Replace: ", s.Replace("foo", "o", "0", 1))  //字符串替换 指定起始位置1 f0o
	p("Split: ", s.Split("a-b-c-d-e", "-"))        //字符串切割 [a b c d e]
	p("ToLower: ", s.ToLower("TEST"))              //字符串 小写转换 test
	p("ToUpper: ", s.ToUpper("test"))              //字符串 大写转换 TEST
	p("Len: ", len("hello"))                       //字符串长度
	p("Char:", "hello"[1])                         //标取字符串中的字符，类型为byte

	str := "ruckoyf"
	index := str[len(str)-2]
	p("len is ", str[0], len(str), str[len(str)-1])
	if index == 'y' {
		p("ok", index, 'y', 'f')
	} else {
		p("no", index, 'y', 'f')
	}

	str2 := "88989"
	by1 := []byte(str2)

	in := bytes.LastIndexByte(by1, byte('8'))

	p("index is ", in, str2[:in+1])
}
