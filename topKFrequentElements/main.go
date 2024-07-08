package main

import "fmt"

func topKFrequent(nums []int, k int) []int {
	if len(nums) == 0 || k == 0 {
		return []int{}
	}

	numsMap := map[int]int{}
	for _, num := range nums {
		numsMap[num]++
	}

	reversedMap := map[int]int{}
	for key, num := range numsMap {
		reversedMap[num] = key
	}

	res := []int{}
	for _, num := range reversedMap {
		if k == 0 {
			break
		}
		res = append(res, num)
		k--
	}

	fmt.Println(numsMap)
	fmt.Println(reversedMap)

	return res
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 2, 2, 2, 1, 3}
	k := 2
	res := topKFrequent(nums, k)

	fmt.Println(res)
}
