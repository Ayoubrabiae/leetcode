package main

import "fmt"

func topKFrequent(nums []int, k int) []int {
	if len(nums) == 0 || k == 0 {
		return []int{}
	}

	numsMap := map[int]int{}
	reversedMap := map[int][]int{}
	keysSlc := []int{}
	res := []int{}

	for _, num := range nums {
		numsMap[num]++
	}

	for key, num := range numsMap {
		reversedMap[num] = append(reversedMap[num], key)
	}

	for key := range reversedMap {
		keysSlc = append(keysSlc, key)
	}

	sortize(keysSlc)

	fmt.Println(numsMap)
	fmt.Println(reversedMap)
	fmt.Println(keysSlc)

	for i := 0; i < len(keysSlc); i++ {
		res = append(res, reversedMap[keysSlc[i]]...)
	}

	return res[:k]
}

func sortize(nums []int) {
	for i := 0; i < len(nums); i++ {
		for k := 0; k < len(nums); k++ {
			if nums[k] < nums[i] {
				temp := nums[k]
				nums[k] = nums[i]
				nums[i] = temp
			}
		}
	}
}

func main() {
	nums := []int{4, 1, -1, 2, -1, 2, 3}
	k := 2
	res := topKFrequent(nums, k)

	fmt.Println(res)
}
