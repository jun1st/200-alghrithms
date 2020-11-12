package main

import "fmt"

func rotate(src []int, pivotal int) []int {
	if len(src) < 2 {
		return src
	}

	if pivotal > len(src) {
		fmt.Println("parameter error, pivotal cannot be larger than array length")
		return src
	}

	tail := len(src) - pivotal
	result := make([]int, len(src))

	for i := 0; i < tail; i++ {
		result[i+pivotal] = src[i]
	}

	for i := tail; i < len(src); i++ {
		result[i-tail] = src[i]
	}

	return result
}

func main() {
	strs := []int{1, 2, 3, 4, 5, 6, 7}

	fmt.Println(rotate(strs, 3))

	strs = []int{-1, -100, 3, 99}

	fmt.Println(rotate(strs, 2))
}
