package main

import "fmt"

func removeElement(nums []int) []int {
	for i := 0; i < len(nums)-1; {
		if nums[i] == nums[i+1] {
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			i++
		}
	}

	return nums
}

func main() {
	nums := []int{1, 1, 2, 2, 3, 4, 5, 5, 5, 6, 6, 9}

	fmt.Println(removeElement(nums))
}
