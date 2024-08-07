package main

import "fmt"

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	fmt.Println(removeElement(nums, val))
}

func removeElement(nums []int, val int) int {
	index := 0
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums, i)
		if nums[i] != val {
			nums[index] = nums[i]
			index++
		}
	}

	return index
}
