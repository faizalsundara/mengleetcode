package main

import (
	"fmt"
	"sort"
	"strings"
)

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

// strings.Index is a built-in function in Go that returns the index of the first occurrence of needle in haystack, or -1 if needle is not found.
func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	return strings.Index(haystack, needle)
}

func threeSum(nums []int) [][]int {
	result := [][]int{}

	// Sort the array to make it easier to avoid duplicates
	sort.Ints(nums)

	n := len(nums)
	for i := 0; i < n-2; i++ {
		// Skip duplicates for the first element
		if i > 0 && nums[i] == nums[i-1] {
			fmt.Println("---", i)
			fmt.Println("---", nums[i])
			fmt.Println("---", nums[i-1])
			continue
		}
		left, right := i+1, n-1
		// -4 -1 -1 0 1 2
		for left < right {
			fmt.Println("looping--2", i, nums[i])
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})

				// Skip duplicates for the second element
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				// Skip duplicates for the third element
				for left < right && nums[right] == nums[right-1] {
					right--
				}

				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}

	return result
}

func lengthOfLastWord(s string) int {
	result := 0

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == byte(' ') {
			if result != 0 {
				break
			}
			continue
		}
		result++

	}
	return result
}

func plusOne(digits []int) []int {
	n := len(digits)

	for i := n - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			return digits
		}
		// untuk i != 9 set 0
		digits[i] = 0
	}
	// diappend ditambah element 1 di digit depannya
	return append([]int{1}, digits...)
}

func main() {
	// fmt.Println(strStr("sadbutsad", "sad"))
	// fmt.Println(strStr("hello", "ll"))
	// a := []int{-1, 0, 1, 2, -1, -4}
	// fmt.Println(plusOne([]int{1, 2, 3}))
	fmt.Println(plusOne([]int{9, 2}))
}
