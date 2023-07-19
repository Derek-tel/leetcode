package subarraySum

func subarraySum(nums []int, k int) int {
	prefix := make(map[int]int)

	prefix[0] = 1
	sum := 0
	count := 0
	for i := 0; i < len(nums); i++ {
		sum = sum + nums[i]
		if _, ok := prefix[sum-k]; ok {
			count = prefix[sum-k] + count
		}
		prefix[sum] += 1
	}
	return count
}

func test(nums []int, k int) int {
	count := 0
	prefix := make(map[int]int)
	prefix[0] = 1
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum = sum + nums[i]
		if _, ok := prefix[sum-k]; ok {
			count = count + prefix[sum-k]
		}
		prefix[sum] += 1
	}
	return count
}

func get(nums []int, k int) int {
	count := 0
	prefix := make(map[int]int)
	prefix[0] = 1
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum = sum + nums[i]
		if _, ok := prefix[sum-k]; ok {
			count += prefix[sum-k]
		}
		prefix[sum] += 1
	}
	return count
}

func four(nums []int, k int) int {
	count := 0
	prefix := make(map[int]int)
	prefix[0] = 1
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum = sum + nums[i]
		if _, ok := prefix[sum-k]; ok {
			count += prefix[sum-k]
		}
		prefix[sum] += 1
	}
	return count
}

func five(nums []int, k int) int {
	count := 0
	prefix := make(map[int]int)
	prefix[0] = 1
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if _, ok := prefix[sum-k]; ok {
			count += prefix[sum-k]
		}
		prefix[sum] += 1
	}
	return count
}

func six(nums []int, k int) int {
	count := 0
	prefix := make(map[int]int)
	prefix[0] = 1
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if _, ok := prefix[sum-k]; ok {
			count += prefix[sum-k]
		}
		prefix[sum] += 1
	}
	return count
}

func seven(nums []int, k int) int {
	count := 0
	prefix := make(map[int]int)
	prefix[0] = 1
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if _, ok := prefix[sum-k]; ok {
			count += prefix[sum-k]
		}
		prefix[sum] += 1
	}
	return count
}

func eight(nums []int, k int) int {
	count := 0
	prefix := make(map[int]int)
	prefix[0] = 1
	sum := 0

	for i := 0; i < len(nums); i++ {
		sum = sum + nums[i]
		if _, ok := prefix[sum-k]; ok {
			count += prefix[sum-k]
		}
		prefix[sum]++
	}
	return count
}
