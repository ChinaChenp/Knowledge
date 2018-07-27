package main

import (
	"fmt"
	"net/url"
	"strings"
)

func Print(ss []string) {
	fmt.Println("len = ", len(ss))
	fmt.Println(ss)

	fmt.Println("---------------------------")
}

func Split() {
	s := "aa bb cc\tdd"
	vecstr := strings.Split(s, " ")
	Print(vecstr)

	s = "11\tbb\tdd\t1 dfsalj\t"
	vecstr = strings.FieldsFunc(s, func(s rune) bool {
		if s == '\t' || s == ' ' {
			return true
		} else {
			return false
		}
	})
	Print(vecstr)

}

func TrimSpace() {
	s := "\t\n a lone \t gopher \n\t\r\n "
	fmt.Println(len(s), s)
	fmt.Println("=============================")
	re := strings.TrimSpace(s)
	fmt.Println(len(re), re)
}

func main() {
	Split()

	TrimSpace()

	ticketStr := "%%20shit%20%21cao"
	for i := 0; i < 10; i++ {
		if strings.Contains(ticketStr, "%") {
			fmt.Println("====================", i, ticketStr)
			ticketStr, _ = url.QueryUnescape(ticketStr)
			fmt.Println("====================", i, ticketStr)
		}
	}

	fmt.Println(ticketStr)
}
