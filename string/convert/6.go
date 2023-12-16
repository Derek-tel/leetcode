package convert

import "bytes"

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
