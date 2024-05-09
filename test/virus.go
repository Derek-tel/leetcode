package main

import (
	"errors"
	"fmt"
)

/*
input set=[1,2,3,4,5],target=10, output=[[1,2,3,4],[2,3,5],[1,4,5]]
*/

func combine(nums []int, target int) [][]int {
	var result [][]int
	var temp []int
	var handler func(int, int)
	handler = func(index int, tar int) {
		if tar == 0 {
			result = append(result, append([]int(nil), temp...))
			return
		}

		if index == len(nums) {
			return
		}

		//不选
		handler(index+1, tar)
		//选
		if nums[index] <= tar {
			temp = append(temp, nums[index])
			handler(index+1, tar-nums[index])
			temp = temp[:len(temp)-1]
		}
	}
	handler(0, target)
	return result
}

func main() {
	fmt.Println(combine([]int{1, 2, 3, 4, 5}, 10))
	var ErrKeyNotExistOrContentEmpty = errors.New("config key not exist or config value empty")

	fmt.Println(fmt.Errorf("key %s not exist %w", "hah", ErrKeyNotExistOrContentEmpty))
}
