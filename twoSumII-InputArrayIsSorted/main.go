package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	for i, num := range numbers {
		// needNum := target - num

		low, heigh := 0, len(numbers)-1
		mid := (low + heigh) / 2

		for low <= heigh {
			if numbers[mid]+num == target && mid != i {
				return []int{i + 1, mid + 1}
			} else if numbers[mid] >= target {
				heigh = mid - 1
				mid = (low + heigh) / 2
			} else if numbers[mid] <= target {
				low = mid + 1
				mid = (low + heigh) / 2
			}
		}

	}

	return []int{}
}

func main() {
	numbers := []int{-1, 0}
	target := -1

	fmt.Println(twoSum(numbers, target))
}
