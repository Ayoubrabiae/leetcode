package main

import "fmt"

func productExceptSelf(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}

	prefixSlc := make([]int, len(nums))
	copy(prefixSlc, nums)
	suffixSlc := make([]int, len(nums))
	copy(suffixSlc, nums)
	left, right, i := -1, len(nums)-1, 0

	for left < len(nums)-1 {
		if left == i-1 && right == i+1 {
			if left >= 0 && right < len(suffixSlc) {
				nums[i] = prefixSlc[left] * suffixSlc[right]
			} else if left < 0 && right < len(suffixSlc) {
				nums[i] = suffixSlc[right]
			} else if right >= len(suffixSlc) && left >= 0 {
				nums[i] = prefixSlc[left]
			}
			i++
		}

		if left < i-1 {
			left++
			if left < len(prefixSlc) && left > 0 {
				prefixSlc[left] = prefixSlc[left] * prefixSlc[left-1]
			}
		}
		if right > i+1 {
			right--
			if right < len(suffixSlc) && right > 0 {
				suffixSlc[right] = suffixSlc[right] * suffixSlc[right+1]
			}
		} else if right == i {
			right++
		}
	}

	return nums
}

func main() {
	slc := []int{1, 2, 3, 4}
	res := productExceptSelf(slc)

	fmt.Println(res)
}
