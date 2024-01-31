package main

func one(nums []int) []int {
	tag := 0
	for i := 0; i < len(nums); i++ {
		tag ^= nums[i]
	}

	flag := tag & (-tag)

	left, right := 0, 0
	for _, num := range nums {
		if num&flag > 0 {
			left ^= num
		} else {
			right ^= num
		}
	}
	return []int{left, right}
}