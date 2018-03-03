package main

import "mymath"
import "strconv"
import "fmt"
import "os"

func main() {
	args := os.Args
	if args == nil || len(args) < 2 {
		fmt.Println("args error")
		return
	}

	switch args[1] {
	case "add":
		if len(args) != 4 {
			fmt.Println("add intput args error.")
			return
		}

		v1, err1 := strconv.Atoi(args[2])
		v2, err2 := strconv.Atoi(args[3])
		if err1 != nil && err2 != nil {
			fmt.Println("add input args error argv1 %s, argv2 %s", args[2], args[3])
			return
		}

		ret := mymath.Add(v1, v2)
		fmt.Println("add result: ", ret)
	case "sqrt":
		if len(args) != 3 {
			fmt.Println("sqrt intput args error.")
			return
		}

		v1, err1 := strconv.Atoi(args[2])
		if err1 != nil {
			fmt.Println("sqrt input args error argv1 %s", args[2])
			return
		}

		ret := mymath.Sqrt(v1)
		fmt.Println("sqrt result: ", ret)
	default:
		fmt.Println("nothing..............")
	}
}
