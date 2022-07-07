package main

import "fmt"

func solution(numbers []int) int {
	allNumbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var resultNumbers []int
	var result int
	for _, allNumber := range allNumbers {
		flag := false
		for _, n := range numbers {
			if allNumber == n {
				flag = true
			}
		}

		if !flag {
			resultNumbers = append(resultNumbers, allNumber)
		}

	}

	for _, rn := range resultNumbers {
		result = result + rn
	}
	return result
}

func main() {
	numbers := []int{5, 8, 4, 0, 6, 7, 9}
	fmt.Println(solution(numbers))
}
