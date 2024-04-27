package convert

import (
	"bytes"
)

func convert(s string, numRows int) string {
	if numRows < 2 || numRows >= len(s) {
		return s
	}
	mat := make([][]byte, numRows)
	i, flag := 0, -1
	for _, ch := range s {
		mat[i] = append(mat[i], byte(ch))
		if i == 0 || i == numRows-1 {
			flag = -flag
		}
		i += flag
	}
	return string(bytes.Join(mat, nil))
}

func two(s string, numRow int) string {
	if numRow < 2 || numRow >= len(s) {
		return s
	}
	snake := make([][]byte, numRow)
	i, flag := 0, -1
	for _, ch := range s {
		snake[i] = append(snake[i], byte(ch))
		if i == 0 || i == numRow-1 {
			flag = -flag
		}
		i += flag
	}
	return string(bytes.Join(snake, nil))
}

func three(s string, numRow int) string {
	if numRow < 2 || numRow >= len(s) {
		return s
	}
	snake := make([][]byte, numRow)
	i, flag := 0, -1
	for _, v := range s {
		snake[i] = append(snake[i], byte(v))
		if i == 0 || i == numRow-1 {
			flag = -flag
		}
		i += flag
	}
	return string(bytes.Join(snake, nil))
}

func four(s string, numRow int) string {
	if numRow < 2 || numRow > len(s) {
		return s
	}
	snake := make([][]byte, numRow)
	i, flag := 0, -1
	for _, v := range s {
		snake[i] = append(snake[i], byte(v))
		if i == 0 || i == numRow-1 {
			flag = -flag
		}
		i += flag
	}
	return string(bytes.Join(snake, nil))
}

func five(s string, numRow int) string {
	if numRow < 2 || numRow > len(s) {
		return s
	}
	snake := make([][]byte, numRow)
	i, flag := 0, -1
	for _, v := range s {
		snake[i] = append(snake[i], byte(v))
		if i == 0 || i == numRow-1 {
			flag = -flag
		}
		i += flag
	}
	return string(bytes.Join(snake, nil))
}
