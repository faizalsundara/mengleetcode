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

func addBinary(a string, b string) string {
	// for len(a) < len(b) {
	// 	a = "0" + a
	// }
	// for len(b) < len(a) {
	// 	b = "0" + b
	// }
	lenA := len(a)
	lenB := len(b)

	// Ensure both strings are of the same length by padding the shorter one with leading zeros
	for lenA < lenB {
		fmt.Println("a0", a)
		a = "0" + a
		fmt.Println("a1", a)
		// lenA++
		lenA = len(a)
		fmt.Println("a2", a)
	}
	for lenB < lenA {
		fmt.Println("b0", a)
		b = "0" + b
		fmt.Println("b1", a)
		// lenB++
		fmt.Println("b2", a)
	}

	carry := 0
	result := ""

	// Perform addition from the last digit to the first
	fmt.Println("a10", a)
	fmt.Println("b10", b)
	for i := len(a) - 1; i >= 0; i-- {
		bitA := int(a[i] - '0') // Convert character to integer (0 or 1)
		bitB := int(b[i] - '0')
		sum := bitA + bitB + carry
		result = string((sum%2)+'0') + result // Append the current bit result
		fmt.Println("--", sum)
		carry = sum / 2 // Update carry for the next iteration
		fmt.Println("--Cary", carry)
	}

	// If there is a carry left, prepend it to the result
	if carry > 0 {
		result = "1" + result
	}

	return result
}

func customSorting(strArr []string) []string {
	var result []string
	var resultGanjil []string
	var resultGenap []string
	fmt.Println()
	for i := 0; i <= len(strArr)-1; i++ {
		// if i == len(strArr) -1
		if len(strArr[i])%2 != 0 {
			fmt.Println(strArr[i])
			resultGanjil = append(resultGanjil, strArr[i])
		} else if len(strArr[i])%2 == 0 {
			fmt.Println(strArr[i])
			resultGenap = append(resultGenap, strArr[i])
		}
	}
	result = append(result, resultGanjil...)
	result = append(result, resultGenap...)
	fmt.Println("---", result)
	return result
}

func main() {
	// fmt.Println(strStr("sadbutsad", "sad"))
	// fmt.Println(strStr("hello", "ll"))
	// a := []int{-1, 0, 1, 2, -1, -4}
	// fmt.Println(plusOne([]int{1, 2, 3}))
	// fmt.Println(plusOne([]int{9, 2}))
	fmt.Println(addBinary("1", "111"))
	// fmt.Println(customSorting([]string{"hayu", "ayu"}))
}
