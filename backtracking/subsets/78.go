package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func subsets(nums []int) [][]int {
	var dfs func(int)
	set := []int{}
	res := [][]int{}
	dfs = func(cur int) {
		if cur == len(nums) {
			res = append(res, append([]int(nil), set...))
			return
		}

		//选
		set = append(set, nums[cur])
		dfs(cur + 1)
		set = set[:len(set)-1]
		//不选
		dfs(cur + 1)
	}
	dfs(0)
	return res
}

func demo1(nums []int) [][]int {
	var helper func(int)
	set := []int{}
	var result [][]int
	helper = func(index int) {
		if index == len(nums) {
			result = append(result, append([]int(nil), set...))
			return
		}

		set = append(set, nums[index])
		helper(index + 1)
		set = set[:len(set)-1]

		helper(index + 1)
	}
	helper(0)
	return result
}

func test1(nums []int) [][]int {
	var dfs func(int)
	res := [][]int{}
	temp := []int{}
	dfs = func(index int) {
		if index == len(nums) {
			res = append(res, append([]int(nil), temp...))
			return
		}
		//不选
		dfs(index + 1)
		//选择
		temp = append(temp, nums[index])
		dfs(index + 1)
		temp = temp[:len(temp)-1]
	}
	dfs(0)
	return res
}

func demo(nums []int) [][]int {
	var helper func(int)
	var result [][]int
	var temp []int
	helper = func(index int) {
		if index == len(nums) {
			result = append(result, append([]int(nil), temp...))
			return
		}
		helper(index + 1)
		temp = append(temp, nums[index])
		helper(index + 1)
		temp = temp[:len(temp)-1]
	}
	helper(0)
	return result
}

func five(nums []int) [][]int {
	var result [][]int
	var temp []int
	var helper func(int)
	helper = func(index int) {
		if index == len(nums) {
			result = append(result, append([]int(nil), temp...))
			return
		}
		//do not pick
		helper(index + 1)
		//pick
		temp = append(temp, nums[index])
		helper(index + 1)
		temp = temp[:len(temp)-1]
	}
	helper(0)
	return result
}

func six(nums []int) [][]int {
	var result [][]int
	var com []int
	var handler func(int)
	handler = func(index int) {
		if index == len(nums) {
			result = append(result, append([]int(nil), com...))
			return
		}
		handler(index + 1)

		com = append(com, nums[index])
		handler(index + 1)
		com = com[:len(com)-1]
	}
	handler(0)
	return result
}

func seven(nums []int) [][]int {
	var result [][]int
	var com []int
	var handler func(int)
	handler = func(index int) {
		if index == len(nums) {
			result = append(result, append([]int(nil), com...))
			return
		}

		handler(index + 1)

		com = append(com, nums[index])
		handler(index + 1)
		com = com[:len(com)-1]
	}
	handler(0)
	return result
}

func eight(nums []int) [][]int {
	var result [][]int
	var com []int
	var handler func(int)
	handler = func(index int) {
		if index == len(nums) {
			result = append(result, append([]int(nil), com...))
			return
		}

		handler(index + 1)

		com = append(com, nums[index])
		handler(index + 1)
		com = com[:len(com)-1]
	}
	handler(0)
	return result
}

func nine(nums []int) [][]int {
	var result [][]int
	var com []int
	var handler func(int)
	handler = func(index int) {
		if index == len(nums) {
			result = append(result, append([]int(nil), com...))
			return
		}

		handler(index + 1)

		com = append(com, nums[index])
		handler(index + 1)
		com = com[:len(com)-1]
	}
	handler(0)
	return result
}

func ten(nums []int) [][]int {
	var result [][]int
	var temp []int
	var handler func(int)
	handler = func(index int) {
		if index == len(nums) {
			result = append(result, append([]int(nil), temp...))
			return
		}

		handler(index + 1)

		temp = append(temp, nums[index])
		handler(index + 1)
		temp = temp[:len(temp)-1]
	}
	handler(0)
	return result
}

func eleven(nums []int) [][]int {
	var result [][]int
	var temp []int
	var handler func(int)
	handler = func(index int) {
		if index == len(nums) {
			result = append(result, append([]int(nil), temp...))
			return
		}

		handler(index + 1)

		temp = append(temp, nums[index])
		handler(index + 1)
		temp = temp[:len(temp)-1]
	}
	handler(0)
	return result
}

func twelve(nums []int) [][]int {
	var result [][]int
	var temp []int
	var handler func(int)

	handler = func(index int) {
		if index == len(nums) {
			result = append(result, append([]int(nil), temp...))
			return
		}

		handler(index + 1)

		temp = append(temp, nums[index])
		handler(index + 1)
		temp = temp[:len(temp)-1]
	}
	handler(0)
	return result
}

func main() {
	test := []int{1, 2, 3}
	fmt.Println(test1(test))

	blePhoneKeyByte, _ := base64.StdEncoding.DecodeString("ZGQMRnE0CAeu/Av+Vz2glzo3OyXKgQ/ew3y7mII5Ft0=")

	//ble phone key and para
	blePhoneKey := hex.EncodeToString(blePhoneKeyByte)
	fmt.Println(blePhoneKey)
	origin := "746800000300202303240000000005821735ffb991ef8d0069aa9a4e"
	origin = origin + "80"
	originByteLen := len(origin) / 2
	n := 16
	mod := originByteLen % n
	for i := 0; i < n-mod; i++ {
		origin += "00"
	}
	fmt.Println(origin)
	blePhoneKeyDataByte, _ := hex.DecodeString(origin)
	fmt.Println(base64.StdEncoding.EncodeToString(blePhoneKeyDataByte))
}
