package main

import "fmt"

var count int

// todo : 翻转单词
func reverseWords(str []byte) string {
	end := len(str) - 1
	reverse(str, 0, end)
	slow, fast := 0, 0
	for slow != end {
		if str[slow] == ' ' {
			slow++
			fast++
		} else if str[fast] == ' ' {
			reverse(str, slow, fast-1)
			slow = fast
		} else if fast == end {
			reverse(str, slow, fast)
			slow = fast
		} else {
			fast++
		}
	}

	return string(str)
}

func reverse(str []byte, beg, end int) {
	for beg < end {
		str[beg], str[end] = str[end], str[beg]
		beg++
		end--
	}
}

func main() {
	str := []byte("  abd12 defg d bedf")
	fmt.Printf("%s --> ", string(str))
	reverseWords(str)
	fmt.Println(string(str))
}
