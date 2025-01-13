package main

import (
	"fmt"
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

func main() {
	// fmt.Println(strStr("sadbutsad", "sad"))
	fmt.Println(strStr("hello", "ll"))
}
