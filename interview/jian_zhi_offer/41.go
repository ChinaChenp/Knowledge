package main

import "fmt"

// todo ： 和为s的两个数字
func FindNumbersWithSum(arr []int, sum int) (int, int) {
	beg, end := 0, len(arr)-1

	for beg < end {
		now := arr[beg] + arr[end]
		if now == sum {
			return beg, end
		} else if now > sum {
			end--
		} else {
			beg++
		}
	}
	return -1, -1
}

// todo : 和为s的连续正数序列
func FindContinuousSequence(sum int) {
	if sum < 3 {
		return
	}

	small, big := 0, 1
	curSum := small + big
	for small <= sum/2 {
		curSum = (small + big) * (big - small + 1) / 2
		//fmt.Println("==>", curSum, small, big)
		if curSum == sum {
			print(small, big)
			small++
		} else if curSum < sum {
			big++
		} else {
			small++
		}
	}
}

func print(small, big int) {
	for i := small; i <= big; i++ {
		fmt.Printf("%d,", i)
	}
	fmt.Println()
}

func main() {
	arr := []int{1, 2, 4, 7, 11, 15}

	sum := 18
	a, b := FindNumbersWithSum(arr, sum)
	fmt.Printf("%d + %d = %d\n", arr[a], arr[b], sum)

	FindContinuousSequence(9)
}
