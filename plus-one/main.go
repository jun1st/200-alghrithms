package main

import "fmt"

func plus(nums []int) []int {

	for i := len(nums) - 1; i >= 0; i-- {

		if nums[i] != 9 {
			nums[i] = nums[i] + 1
			break
		} else {
			nums[i] = 0
		}
	}

	if nums[0] == 0 {
		new := []int{1}
		nums = append(new, nums[:]...)
	}

	return nums
}

func main() {
	a := []int{1, 2, 3}

	fmt.Println(plus(a))

	a = []int{1, 2, 9}
	fmt.Println(plus(a))

	a = []int{1, 9, 9}
	fmt.Println(plus(a))

	a = []int{9, 9, 9}
	fmt.Println(plus(a))
}
