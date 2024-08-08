package main

import "fmt"

func main() {
	nums := []int{0, 0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(nums))
	fmt.Println(nums)
}

func removeDuplicates(nums []int) int {
	j := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[j] = nums[i]
			j++
		}
	}

	return j
}

// func removeDuplicates(nums []int) int {
// 	index := 0
// 	myMap := make(map[int]int)

// 	for _, num := range nums {
// 		if count, exists := myMap[num]; exists {
// 			myMap[num] = count + 1
// 		} else {
// 			myMap[num] = 1
// 			nums[index] = num
// 			index++
// 		}
// 	}

// 	return index
// }
