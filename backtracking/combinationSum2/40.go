package combinationSum2

import "sort"

func ont(nums []int, target int) [][]int {
	var temp []int
	var result [][]int
	var handler func(int, int, int)
	sort.Ints(nums)
	handler = func(start int, index int, tar int) {
		if tar == 0 {
			result = append(result, append([]int(nil), temp...))
			return
		}

		for i := start; i < len(nums); i++ {
			if tar-nums[i] < 0 {
				return
			}
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			temp = append(temp, nums[i])
			handler(i+1, index+1, tar-nums[i])
			temp = temp[:len(temp)-1]
		}
	}
	handler(0, 0, target)
	return result
}
