package main

import (
	"fmt"
	"sync"
)

const MaxBatchCount  =  3

func getOdInfoStatusImpl(grid string, ids []int64) ([]int64, error) {
	out := make([]int64, 0)
	for _, v := range ids {
		out = append(out, v + 100)
	}
	return out, nil
}

func GetOdInfoStatus(grid string, ids []int64) {
	total := len(ids)
	//if total == 0 {
	//	return
	//}

	maxRange := total/MaxBatchCount
	if total % MaxBatchCount != 0 {
		maxRange = total / MaxBatchCount + 1
	}

	var wg sync.WaitGroup
	response := make(chan []int64)
	wg.Add(maxRange)
	for i := 0; i < maxRange; i++ {
		max := total
		if (i + 1) * MaxBatchCount < total {
			max = (i + 1) * MaxBatchCount
		}

		tmp := ids[i*MaxBatchCount : max]
		go func(gid string, id []int64, i int) {
			defer wg.Done()
			fmt.Println(i, "------->", id)
			re, err := getOdInfoStatusImpl(grid, id)
			if err == nil && len(re) != 0 {
				response <- re
			}
		}(grid, tmp, i)
	}

	go func() {
		wg.Wait()
		close(response)
	}()

	out := make([]int64, 0)
	for resp := range response {
		out = append(out, resp...)
	}
	fmt.Println("re ---->", out)
}

func main() {
	in := make([]int64, 0)
	for i := 0 ;i < 12; i++ {
		in = append(in, int64(i))
	}

	GetOdInfoStatus("a", in)
}