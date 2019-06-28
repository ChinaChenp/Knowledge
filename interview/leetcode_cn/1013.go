package main

import "fmt"

// 注意条件 i+1 < j 所以数组会是连续的
//https://leetcode-cn.com/problems/partition-array-into-three-parts-with-equal-sum/comments/
func canThreePartsEqualSum(A []int) bool {
	sum := 0
	for _, v := range A {
		sum += v
	}

	avg := sum / 3
	if sum % 3 != 0 {
		return false
	}

	beg, end := 0, len(A) - 1
	leftSum, rightSum := 0, 0
	for beg <= end {
		if leftSum != avg {
			leftSum += A[beg]
			beg += 1
		}

		if rightSum != avg {
			rightSum += A[end]
			end -= 1
		}

		fmt.Println(leftSum, rightSum)
		if leftSum == avg && rightSum == avg {
			// 中间合多余的，从平均数理解
			return true
		}
	}
	return false
}

func main() {
	arr := []int{0,2,1,-6,6,-7,9,1,2,0,1}
	re := canThreePartsEqualSum(arr)
	fmt.Println(re)
}
