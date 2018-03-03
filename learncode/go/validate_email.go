package main

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

type ParamError struct {
	Errno int
}

func generr(errno int) ParamError {
	return ParamError{errno}
}

func ParseEmail(item interface{}) string {
	if item == nil || item == "" {
		return ""
	}

	fmt.Println("fuck 3")
	if s, ok := item.(string); ok {
		fmt.Println("fuck 4", s)
		email := ""
		if strings.Contains(s, "%") {
			email, _ = url.QueryUnescape(s)
			fmt.Println("fuck 5, ", s)
			if email == "" {
				panic(generr(12345))
			}
		} else {
			if s == "" {
				return ""
			}
		}

		fmt.Println("fuck 1")
		if len(s) > 128 {
			panic(generr(12345))
		}

		fmt.Println("fuck 2")
		if IsEmailFormatValid(s) {
			return s
		}
	}
	panic(generr(1234567))
}

func IsEmailFormatValid(email string) bool {
	if !strings.ContainsAny(email, "@") {
		return false
	}
	emailregexp := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if !emailregexp.MatchString(email) {
		return false
	}
	return true
}

func main() {
	var email interface{}
	email = "sk@1234"
	re := ParseEmail(email)
	fmt.Println()
	fmt.Println(re)
}
