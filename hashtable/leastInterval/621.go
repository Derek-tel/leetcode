package main

import (
	"fmt"
	"sort"
)

func leastInterval(tasks []byte, n int) int {
	counter := make([]int, 26)
	for _, task := range tasks {
		counter[task-'A']++
	}
	sort.Ints(counter)
	most := counter[len(counter)-1]
	minLength := (n+1)*(most-1) + 1

	for i := len(counter) - 2; i >= 0; i-- {
		if counter[i] == most {
			minLength++
		}
	}

	return max(len(tasks), minLength)
}

func test(task []byte, n int) int {
	counter := make([]int, 26)
	for _, b := range task {
		counter[b-'A']++
	}
	sort.Ints(counter)
	most := counter[len(counter)-1]
	minLength := (n+1)*(most-1) + 1
	for i := 0; i < len(counter)-1; i++ {
		if counter[i] == most {
			minLength++
		}
	}
	return max(len(task), minLength)
}

func get(task []byte, n int) int {
	counter := make([]int, 26)
	for _, b := range task {
		counter[b-'A']++
	}
	sort.Ints(counter)
	most := counter[len(counter)-1]
	minLength := (n+1)*(most-1) + 1
	for i := 0; i < len(counter)-1; i++ {
		if counter[i] == most {
			minLength++
		}
	}
	return max(len(task), minLength)
}

func four(task []byte, n int) int {
	counter := make([]int, 26)
	for _, b := range task {
		counter[b-'A']++
	}
	sort.Ints(counter)
	most := counter[len(counter)-1]
	minLength := (n+1)*(most-1) + 1
	for i := 0; i < len(counter)-1; i++ {
		if counter[i] == most {
			minLength++
		}
	}
	return max(len(task), minLength)
}

func five(task []byte, n int) int {
	counter := make([]int, 26)
	for i := 0; i < len(task); i++ {
		counter[task[i]-'A']++
	}
	sort.Ints(counter)
	most := counter[len(counter)-1]
	minLength := (n+1)*(most-1) + 1
	for i := 0; i < len(counter)-1; i++ {
		if counter[i] == most {
			minLength++
		}
	}
	return max(len(task), minLength)
}

func six(task []byte, n int) int {
	counter := make([]int, 26)
	for i := 0; i < len(task); i++ {
		counter[task[i]-'A']++
	}
	sort.Ints(counter)
	most := counter[len(counter)-1]
	minlength := (n+1)*(most-1) + 1
	for i := 0; i < len(counter)-1; i++ {
		if counter[i] == most {
			minlength++
		}
	}
	return max(len(task), minlength)
}

func seven(tasks []byte, n int) int {
	counter := make([]int, 26)
	for i := 0; i < len(tasks); i++ {
		counter[tasks[i]-'A']++
	}
	sort.Ints(counter)
	most := counter[len(counter)-1]
	minLength := (n+1)*(most-1) + 1
	for i := 0; i < len(counter)-1; i++ {
		if counter[i] == most {
			minLength++
		}
	}
	return max(len(tasks), minLength)
}

func eight(tasks []byte, n int) int {
	counter := make([]int, 26)
	for i := 0; i < len(tasks); i++ {
		counter[tasks[i]-'A']++
	}
	sort.Ints(counter)
	most := counter[len(counter)-1]
	minLength := (n+1)*(most-1) + 1
	for i := 0; i < len(counter)-1; i++ {
		if counter[i] == most {
			minLength++
		}
	}
	return max(len(tasks), minLength)
}

func nine(tasks []byte, n int) int {
	counter := make([]int, 26)
	for i := 0; i < len(tasks); i++ {
		counter[tasks[i]-'A']++
	}
	sort.Ints(counter)
	most := counter[len(counter)-1]
	minLength := (n+1)*(most-1) + 1
	for i := 0; i < len(counter)-1; i++ {
		if counter[i] == most {
			minLength++
		}
	}
	return max(len(tasks), minLength)
}

func ten(task []byte, n int) int {
	counter := make([]int, 26)
	for i := 0; i < len(task); i++ {
		counter[task[i]-'A']++
	}
	sort.Ints(counter)
	most := counter[len(counter)-1]
	minLength := (n+1)*(most-1) + 1
	for i := 0; i < len(counter)-1; i++ {
		if counter[i] == most {
			minLength++
		}
	}
	return max(len(task), minLength)
}

func eleven(task []byte, n int) int {
	counter := make([]int, 26)
	for i := 0; i < len(task); i++ {
		counter[task[i]-'A']++
	}
	sort.Ints(counter)
	most := counter[len(counter)-1]
	minLength := (n+1)*(most-1) + 1
	for i := 0; i < len(counter)-1; i++ {
		if counter[i] == most {
			minLength++
		}
	}
	return max(minLength, len(task))
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}

func main() {
	fmt.Println(leastInterval([]byte("AAAAABBBBB"), 2))

}
