package main

import "fmt"

func intersect(num1 []int, num2 []int) []int {
	m0 := map[int]int{}

	for _, v := range num1 {
		m0[v] += 1
	}

	k := 0
	for _, v := range num2 {
		if m0[v] > 0 {
			m0[v] -= 1
			num2[k] = v
			k++
		}
	}

	return num2[0:k]
}

func main() {

	num1 := []int{1, 2, 2, 1}
	num2 := []int{2, 2}

	fmt.Print(intersect(num1, num2))
}
