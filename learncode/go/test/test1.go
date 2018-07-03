package main

import "fmt"

type ParamError struct {
	Errno int
}

func generr(errno int) ParamError {
	return ParamError{errno}
}

func parseString(item interface{}, errno int, required bool) string {
	if (item == nil || item == "") && !required {
		return ""
	}
	if s, ok := item.(string); ok {
		if s != "" || !required {
			return s
		}
	}
	panic(generr(errno))
}

func main() {
	defer func() {
		if re := recover(); re != nil {
			fmt.Println("recover: ", re)
		}
	}()
	mp := make(map[string]interface{})

	mp["chenping"] = "chenping"

	re := parseString(mp["chenping1"], -1, true)
	fmt.Printf("------%s-------\n", re)
}
