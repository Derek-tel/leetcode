package canFinish

func canFinish(numCourses int, prerequisites [][]int) bool {
	var (
		edge   = make([][]int, numCourses)
		visit  = make([]int, numCourses) //0 未搜索 1 搜索中 2 已搜索
		result []int
		valid  = true
		dfs    func(int)
	)
	dfs = func(i int) {
		visit[i] = 1
		for _, site := range edge[i] {
			if visit[site] == 0 {
				dfs(site)
				if !valid {
					return
				}
			} else if visit[site] == 1 {
				valid = false
				return
			}
		}
		visit[i] = 2
		result = append(result, i)
		return
	}
	for _, site := range prerequisites {
		edge[site[1]] = append(edge[site[1]], site[0])
	}
	for i := 0; i < numCourses && valid; i++ {
		if visit[i] == 0 {
			dfs(i)
		}
	}
	return valid
}

func test(numCourses int, prerequisites [][]int) bool {
	var (
		edge   = make([][]int, numCourses)
		visit  = make([]int, numCourses) //0 为搜索 1 搜索中 2 已搜索
		dfs    func(int)
		valid  = true
		result []int
	)
	dfs = func(course int) {
		visit[course] = 1
		for _, site := range edge[course] {
			if visit[site] == 0 {
				dfs(site)
				if !valid {
					return
				}
			} else if visit[site] == 1 {
				valid = false
				return
			}
		}
		visit[course] = 2
		result = append(result, course)
		return
	}
	for _, site := range prerequisites {
		edge[site[1]] = append(edge[site[1]], site[0])
	}

	for i := 0; i < numCourses; i++ {
		if visit[i] == 0 {
			dfs(i)
		}
	}
	return valid
}

func can(numCourse int, prerequisites [][]int) bool {
	var (
		valid = true
		edge  = make([][]int, numCourse)
		visit = make([]int, numCourse) // 0 not yet  1 ing 2 done
		dfs   func(int)
	)

	dfs = func(i int) {
		visit[i] = 1
		for _, site := range edge[i] {
			if visit[site] == 0 {
				dfs(site)
				if !valid {
					return
				}
			} else if visit[site] == 1 {
				valid = false
				return
			}
		}
		visit[i] = 2
		return
	}

	for _, prerequisite := range prerequisites {
		edge[prerequisite[1]] = append(edge[prerequisite[1]], prerequisite[0])
	}

	for i := 0; i < numCourse && valid; i++ {
		if visit[i] == 0 {
			dfs(i)
		}
	}
	return valid
}

func four(numCourses int, prerequisites [][]int) bool {
	var (
		valid  = true
		edge   = make([][]int, numCourses)
		visit  = make([]int, numCourses) //0 not yet 1 ing 2 done
		helper func(int)
	)
	helper = func(i int) {
		visit[i] = 1
		for _, site := range edge[i] {
			if visit[site] == 0 {
				helper(site)
				if !valid {
					return
				}
			} else if visit[site] == 1 {
				valid = false
				return
			}
		}
		visit[i] = 2
		return
	}
	for _, prerequisite := range prerequisites {
		edge[prerequisite[1]] = append(edge[prerequisite[1]], prerequisite[0])
	}
	for i := 0; i < numCourses && valid; i++ {
		if visit[i] == 0 {
			helper(i)
		}
	}
	return valid
}

func five(numCourses int, prerequisites [][]int) bool {
	var (
		valid  = true
		edge   = make([][]int, numCourses)
		visit  = make([]int, numCourses) //0 not yet 1 ing 2 done
		helper func(int)
	)
	helper = func(i int) {
		visit[i] = 1
		for _, site := range edge[i] {
			if visit[site] == 0 {
				helper(site)
				if !valid {
					return
				}
			} else if visit[site] == 1 {
				valid = false
				return
			}
		}
		visit[i] = 2
		return
	}
	for _, prerequisite := range prerequisites {
		edge[prerequisite[1]] = append(edge[prerequisite[1]], prerequisite[0])
	}
	for i := 0; i < numCourses && valid; i++ {
		if visit[i] == 0 {
			helper(i)
		}
	}
	return valid
}

func six(numCourses int, prerequisites [][]int) bool {
	var valid = true
	var edge = make([][]int, numCourses)
	var visit = make([]int, numCourses)
	var helper func(int)
	helper = func(i int) {
		visit[i] = 1
		for _, site := range edge[i] {
			if visit[site] == 0 {
				helper(site)
				if !valid {
					return
				}
			} else if visit[site] == 1 {
				valid = false
				return
			}
		}
		visit[i] = 2
	}
	for _, site := range prerequisites {
		edge[site[1]] = append(edge[site[1]], site[0])
	}
	for i := 0; i < numCourses && valid; i++ {
		if visit[i] == 0 {
			helper(i)
		}
	}
	return valid
}

func seven(numCourses int, prerequisites [][]int) bool {
	var valid = true
	var edge = make([][]int, numCourses)
	var visit = make([]int, numCourses)
	var helper func(int)
	helper = func(index int) {
		visit[index] = 1
		for _, next := range edge[index] {
			if visit[next] == 0 {
				helper(next)
				if !valid {
					return
				}
			} else if visit[next] == 1 {
				valid = false
				return
			}
		}
		visit[index] = 2
	}
	for _, site := range prerequisites {
		edge[site[1]] = append(edge[site[1]], site[0])
	}
	for i := 0; i < numCourses && valid; i++ {
		if visit[i] == 0 {
			helper(i)
		}
	}
	return valid
}

func eight(numCourses int, prerequisites [][]int) bool {
	var valid = true
	var visit = make([]int, numCourses)
	var edge = make([][]int, numCourses)
	var helper func(int)
	helper = func(index int) {
		visit[index] = 1
		for _, next := range edge[index] {
			if visit[next] == 0 {
				helper(next)
				if !valid {
					return
				}
			} else if visit[next] == 1 {
				valid = false
				return
			}
		}
		visit[index] = 2
	}

	for _, site := range prerequisites {
		edge[site[1]] = append(edge[site[1]], site[0])
	}
	for i := 0; i < numCourses && valid; i++ {
		if visit[i] == 0 {
			helper(i)
		}
	}
	return valid
}

func nine(numCourses int, prerequisites [][]int) bool {
	var valid = true
	var visit = make([]int, numCourses)
	var edge = make([][]int, numCourses)
	var helper func(int)
	helper = func(index int) {
		visit[index] = 1
		for _, site := range edge[index] {
			if visit[site] == 0 {
				helper(site)
				if !valid {
					return
				}
			} else if visit[site] == 1 {
				valid = false
				return
			}
		}
		visit[index] = 2
	}
	for _, site := range prerequisites {
		edge[site[1]] = append(edge[site[1]], site[0])
	}
	for i := 0; i < numCourses && valid; i++ {
		if visit[i] == 0 {
			helper(i)
		}
	}
	return valid
}

func ten(numCourses int, prerequisites [][]int) bool {
	var valid = true
	var visit = make([]int, numCourses)
	var edge = make([][]int, numCourses)
	var helper func(int)

	helper = func(course int) {
		visit[course] = 1
		for _, site := range edge[course] {
			if visit[site] == 0 {
				helper(site)
				if !valid {
					return
				}
			} else if visit[site] == 1 {
				valid = false
				return
			}
		}
		visit[course] = 2
	}

	for _, prerequisite := range prerequisites {
		edge[prerequisite[1]] = append(edge[prerequisite[1]], prerequisite[0])
	}

	for i := 0; i < numCourses && valid; i++ {
		if visit[i] == 0 {
			helper(i)
		}
	}
	return valid
}
