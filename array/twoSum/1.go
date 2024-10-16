package twoSum

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		if _, ok := m[another]; ok {
			return []int{m[another], i}
		}
		m[nums[i]] = i
	}
	return nil
}

func demo(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		if _, ok := m[another]; ok {
			return []int{m[another], i}
		}
		m[nums[i]] = i
	}
	return nil
}

func test(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		if _, ok := m[another]; ok {
			return []int{m[another], i}
		}
		m[nums[i]] = i
	}
	return nil
}

func get(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		if _, ok := m[another]; ok {
			return []int{m[another], i}
		}
		m[nums[i]] = i
	}
	return nil
}

func five(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		if _, ok := m[another]; ok {
			return []int{m[another], i}
		}
		m[nums[i]] = i
	}
	return nil
}

func six(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		if p, ok := m[another]; ok {
			return []int{p, i}
		}
		m[nums[i]] = i
	}
	return nil
}

func seven(nums []int, target int) []int {
	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		if p, ok := m[another]; ok {
			return []int{p, i}
		}
		m[nums[i]] = i
	}

	return nil
}

func eight(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		if p, ok := m[another]; ok {
			return []int{p, i}
		}
		m[nums[i]] = i
	}
	return nil
}

func nine(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		if index, ok := m[another]; ok {
			return []int{index, i}
		}
		m[nums[i]] = i
	}
	return nil
}

func ten(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		if index, ok := m[another]; ok {
			return []int{index, i}
		}
		m[nums[i]] = i
	}
	return nil
}

func eleven(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		if index, ok := m[another]; ok {
			return []int{index, i}
		}
		m[nums[i]] = i
	}
	return nil
}

func twelve(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		if index, ok := m[another]; ok {
			return []int{index, i}
		}
		m[nums[i]] = i
	}
	return nil
}

func thirteen(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		if index, ok := m[another]; ok {
			return []int{index, i}
		}
		m[nums[i]] = i
	}
	return nil
}

func fourteen(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		if index, ok := m[another]; ok {
			return []int{index, i}
		}
		m[nums[i]] = i
	}
	return nil
}

func fifteen(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		if index, ok := m[another]; ok {
			return []int{index, i}
		}
		m[nums[i]] = i
	}
	return nil
}

func sixteen(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		if index, ok := m[another]; ok {
			return []int{index, i}
		}
		m[nums[i]] = i
	}
	return nil
}
