package main

import "fmt"

func main() {
	nums1 := []int{2, 0}
	nums2 := []int{1}
	m := 1
	n := 1

	merge(nums1, m, nums2, n)
	fmt.Println("Hello Nihal", nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := n - 1
	k := m + n - 1

	for j >= 0 {
		if i >= 0 && nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			k--
			i--
		} else {
			nums1[k] = nums2[j]
			k--
			j--
		}
	}
}
