package main

import "fmt"

func splitSlice(all []int64, maxSize int) [][]int64 {
	n := len(all) / maxSize
	vecN := n
	if len(all)%maxSize != 0 {
		vecN = n + 1
	}

	total := 0
	out := make([][]int64, vecN)
	for i := 0; i < n; i++ {
		for j := i * maxSize; j < (i+1)*maxSize; j++ {
			out[i] = append(out[i], all[j])
		}
		total += len(out[i])
	}

	if total < len(all) {
		for j := total; j < len(all); j++ {
			out[n] = append(out[n], all[j])
		}
	}
	return out
}

func print(in [][]int64) {
	fmt.Println(len(in))
	for i := 0; i < len(in); i++ {
		fmt.Println(in[i])
	}
}

func main() {
	in := []int64{}
	v := splitSlice(in, 10)

	print(v)
}
