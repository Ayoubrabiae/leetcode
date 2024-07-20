package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	for i, num := range numbers {
		low, heigh := 0, len(numbers)-1
		mid := (low + heigh) / 2
		needNum := target - num

		for low <= heigh {
			if numbers[mid]+num == target && mid != i {
				return []int{i + 1, mid + 1}
			} else if numbers[mid] > needNum {
				heigh = mid - 1
				mid = (low + heigh) / 2
			} else if numbers[mid] <= needNum {
				low = mid + 1
				mid = (low + heigh) / 2
			}
		}

	}

	return []int{}
}

func main() {
	numbers := []int{1, 2, 3, 4, 4, 9, 56, 90}
	target := 8

	fmt.Println(twoSum(numbers, target))
}
