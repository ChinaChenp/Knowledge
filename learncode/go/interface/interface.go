package main

import (
	"bytes"
	"fmt"
	"io"
)

const debug = true

func main() {
	var buf io.Writer
	if buf == nil {
		fmt.Println("buf is nil")
	} else {
		fmt.Println("buf is not nil")
	}

	if debug {
		buf = new(bytes.Buffer) // enable collection of output
	}

	f(buf) // NOTE: subtly incorrect!
	if debug {
		// ...use buf...
	}

	in1 := make(map[string]interface{})
	in2 := make(map[string]interface{})
	in1["chenping1"] = "chenping1"
	in2["chenping"] = "chenping1"

	in1["chenping"] = in2["chenping"]
	fmt.Println("in - >", in1["chenping"].(string))

	key := in1["chenping2"]
	fmt.Println("key - >", key)

	f1(1, 1.25, "fdjsa", true, []byte{'a', 'b', 'f'})
	f1(1, 2)
}

// If out is non-nil, output will be written to it.
func f(out io.Writer) {
	// ...do something...
	if out != nil {
		out.Write([]byte("done!\n"))
	}
}

func f1(args ...interface{}) {
	str := ""
	length := len(args)
	for i := 0; i < length; i++ {
		switch args[i].(type) {
		case []byte:
			str += string(args[i].([]byte))
		default:
			str += fmt.Sprintf("%v", args[i])
		}

		if i+1 != length {
			str += ","

		}
	}
	fmt.Printf("%s\n", str)
}
