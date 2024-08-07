package majorityElement

func majorityElement(nums []int) int {
	count := 0
	var p int

	for _, num := range nums {
		if count == 0 {
			p = num
			count++
		} else {
			if num != p {
				count--
			} else {
				count++
			}
		}
	}
	return p
}

func test(nums []int) int {
	count := 0

	var p int

	for i := 0; i < len(nums); i++ {
		if count == 0 {
			p = nums[i]
			count++
		} else {
			if nums[i] == p {
				count++
			} else {
				count--
			}
		}
	}
	return p
}

func get(nums []int) int {
	count := 0
	var p int

	for i := 0; i < len(nums); i++ {
		if count == 0 {
			p = nums[i]
			count++
		} else {
			if nums[i] == p {
				count++
			} else {
				count--
			}
		}
	}
	return p
}

func four(nums []int) int {
	count := 0
	var p int
	for _, num := range nums {
		if count == 0 {
			p = num
			count++
		} else {
			if num == p {
				count++
			} else {
				count--
			}
		}
	}
	return p
}

func five(nums []int) int {
	count := 0
	var p int
	for _, v := range nums {
		if count == 0 {
			p = v
			count++
		} else {
			if p == v {
				count++
			} else {
				count--
			}
		}
	}
	return p
}

func six(nums []int) int {
	count := 0
	var p int
	for _, num := range nums {
		if count == 0 {
			p = num
			count++
		} else {
			if p == num {
				count++
			} else {
				count--
			}
		}
	}
	return p
}

func seven(nums []int) int {
	count := 0
	var p int
	for _, num := range nums {
		if count == 0 {
			p = num
			count++
		} else {
			if p == num {
				count++
			} else {
				count--
			}
		}
	}
	return p
}

func eight(nums []int) int {
	count := 0
	var p int
	for _, num := range nums {
		if count == 0 {
			p = num
			count++
		} else {
			if p == num {
				count++
			} else {
				count--
			}
		}
	}
	return p
}

func nine(nums []int) int {
	count := 0
	var p int
	for _, num := range nums {
		if count == 0 {
			p = num
			count++
		} else {
			if p == num {
				count++
			} else {
				count--
			}
		}
	}
	return p
}

func ten(nums []int) int {
	count := 0
	var p int
	for _, num := range nums {
		if count == 0 {
			p = num
			count++
		} else {
			if p == num {
				count++
			} else {
				count--
			}
		}
	}
	return p
}

func eleven(nums []int) int {
	count := 0
	var p int
	for _, num := range nums {
		if count == 0 {
			p = num
			count++
		} else {
			if p == num {
				count++
			} else {
				count--
			}
		}
	}
	return p
}

func twelve(nums []int) int {
	count := 0
	var p int
	for _, num := range nums {
		if count == 0 {
			p = num
			count++
		} else {
			if p == num {
				count++
			} else {
				count--
			}
		}
	}
	return p
}

func thirteen(nums []int) int {
	count := 0
	var p int
	for _, num := range nums {
		if count == 0 {
			p = num
			count++
		} else {
			if p == num {
				count++
			} else {
				count--
			}
		}
	}
	return p
}

func fourteen(nums []int) int {
	count := 0
	var p int
	for _, num := range nums {
		if count == 0 {
			p = num
			count++
		} else {
			if p == num {
				count++
			} else {
				count--
			}
		}
	}
	return p
}

func fifteen(nums []int) int {
	count := 0
	var p int
	for _, num := range nums {
		if count == 0 {
			p = num
			count++
		} else {
			if p == num {
				count++
			} else {
				count--
			}
		}
	}
	return p
}
