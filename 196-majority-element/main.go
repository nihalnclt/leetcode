package main

import "fmt"

func main() {
	nums := []int{3, 1, 2, 2, 4, 8, 5, 6, 7}
	fmt.Println(majorityElement(nums))
	fmt.Println(nums)
}

// We can solve this by 3 methods

// 1 - first sort the array and take nums[n/2].
// because the majority element should come at center of the array

// 2 - Hashmap

// 3 - Boyer-Moore majority vote algorithm (below) -
// The majority element is the element that appears more than ⌊n / 2⌋ times

func majorityElement(nums []int) int {
	count := 0
	candidate := 0

	for _, num := range nums {
		if count == 0 {
			candidate = num
		}

		if num == candidate {
			count++
		} else {
			count--
		}
	}

	return candidate
}
