package main

import "fmt"

func revertStr(str string) string {
	ch := []byte(str)
	swap(0, len(ch)-1, ch)
	start := 0

	for i := 0; i < len(ch); i++ {
		if ch[i] == ' ' {
			swap(start, i-1, ch)
			start = i + 1
		}
	}

	swap(start, len(ch)-1, ch)
	return string(ch)
}

func swap(x, y int, ch []byte) {
	for x < y {
		ch[x], ch[y] = ch[y], ch[x]
		x++
		y--
	}
}

func main() {
	test := "student"
	test1 := "i am a student"
	test2 := ""
	test3 := " "
	fmt.Println(revertStr(test))
	fmt.Println(revertStr(test1))
	fmt.Println(fmt.Printf("final:= %stag", revertStr(test2)))
	fmt.Println(fmt.Printf("final:= %sTag", revertStr(test3)))

}
