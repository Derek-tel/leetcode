package main

import (
	"fmt"
	"sort"
)

func reconstructQueue(people [][]int) [][]int {
	ans := [][]int{}
	sort.Slice(people, func(i, j int) bool {
		a, b := people[i], people[j]
		return a[0] > b[0] || a[0] == b[0] && a[1] < b[1]
	})
	for _, person := range people {
		index := person[1]
		ans = append(ans[:index], append([][]int{person}, ans[index:]...)...)
		fmt.Println(ans)
	}
	return ans
}

func test(people [][]int) [][]int {
	ans := [][]int{}

	sort.Slice(people, func(i, j int) bool {
		a, b := people[i], people[j]
		return a[0] > b[0] || a[0] == b[0] && a[1] < b[1]
	})
	fmt.Println(people)
	for _, person := range people {
		index := person[1]
		ans = append(ans[:index], append([][]int{person}, ans[index:]...)...)

	}
	return ans
}

func get(people [][]int) (ans [][]int) {
	sort.Slice(people, func(i, j int) bool {
		return people[i][0] > people[j][0] || people[i][0] == people[j][0] && people[i][1] < people[j][1]
	})
	fmt.Println(people)

	for i := 0; i < len(people); i++ {
		index := people[i][1]
		ans = append(ans[:index], append([][]int{people[i]}, ans[index:]...)...)
	}

	return
}

func four(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		return people[i][0] > people[j][0] || (people[i][0] == people[j][0] && people[i][1] < people[j][1])
	})
	var result [][]int
	for i := 0; i < len(people); i++ {
		index := people[i][1]
		result = append(result[:index], append([][]int{people[i]}, result[index:]...)...)
	}
	return result
}

func five(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		return people[i][0] > people[j][0] || (people[i][0] == people[j][0] && people[i][1] < people[j][1])
	})
	var result [][]int

	for i := 0; i < len(people); i++ {
		index := people[i][1]
		result = append(result[:index], append([][]int{people[i]}, result[index:]...)...)
	}
	return result
}

func six(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		return people[i][0] > people[j][0] || people[i][0] == people[j][0] && people[i][1] < people[j][1]
	})

	var result [][]int

	for i := 0; i < len(people); i++ {
		index := people[i][1]
		result = append(result[:index], append([][]int{people[i]}, result[index:]...)...)
	}
	return result
}

func main() {

	test1 := [][]int{
		{7, 0},
		{4, 4},
		{7, 1},
		{5, 0},
		{6, 1},
		{5, 1},
	}
	fmt.Println(get(test1))

}
