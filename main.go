package main

import "fmt"

func RemoveElement(nums []int, val int) int {
	// var lenArray = len(nums)
	var k int = 0
	for i := 0; i <= len(nums)-1; i++ {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}

	}
	// result := len(lenNums)
	return k
}

func main() {
	fmt.Println(RemoveElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
}
