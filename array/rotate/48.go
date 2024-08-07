package main

func rotate(matrix [][]int) [][]int {
	length := len(matrix)

	for i := 0; i < length/2; i++ {
		matrix[i], matrix[length-i-1] = matrix[length-i-1], matrix[i]
	}

	for i := 0; i < length; i++ {
		for j := 0; j <= i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	return matrix
}

func test(matrix [][]int) [][]int {
	length := len(matrix)

	for i := 0; i < length/2; i++ {
		matrix[i], matrix[length-i-1] = matrix[length-i-1], matrix[i]
	}
	for i := 0; i < length; i++ {
		for j := 0; j <= i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	return matrix
}

func three(matrix [][]int) [][]int {
	length := len(matrix)

	for i := 0; i < length/2; i++ {
		matrix[i], matrix[length-i-1] = matrix[length-i-1], matrix[i]
	}
	for i := 0; i < length; i++ {
		for j := 0; j <= i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	return matrix
}

func four(matrix [][]int) [][]int {
	length := len(matrix)
	for i := 0; i < length/2; i++ {
		matrix[i], matrix[length-i-1] = matrix[length-i-1], matrix[i]
	}
	for i := 0; i < length; i++ {
		for j := 0; j <= i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	return matrix
}

func five(matrix [][]int) [][]int {
	length := len(matrix)
	for i := 0; i < length/2; i++ {
		matrix[i], matrix[length-1-i] = matrix[length-1-i], matrix[i]
	}
	for i := 0; i < length; i++ {
		for j := 0; j <= i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	return matrix
}

func six(matrix [][]int) [][]int {
	length := len(matrix)
	for i := 0; i < length/2; i++ {
		matrix[i], matrix[length-1-i] = matrix[length-1-i], matrix[i]
	}
	for i := 0; i < length; i++ {
		for j := 0; j <= i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	return matrix
}

func seven(matrix [][]int) [][]int {
	length := len(matrix)
	for i := 0; i < length/2; i++ {
		matrix[i], matrix[length-1-i] = matrix[length-1-i], matrix[i]
	}
	for i := 0; i < length; i++ {
		for j := 0; j <= i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	return matrix
}

func eight(matrix [][]int) [][]int {
	length := len(matrix)
	for i := 0; i < length/2; i++ {
		matrix[i], matrix[length-1-i] = matrix[length-1-i], matrix[i]
	}
	for i := 0; i < length; i++ {
		for j := 0; j <= i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	return matrix
}

func nine(matrix [][]int) [][]int {
	length := len(matrix)
	for i := 0; i < length/2; i++ {
		matrix[i], matrix[length-i-1] = matrix[length-i-1], matrix[i]
	}
	for i := 0; i < length; i++ {
		for j := 0; j <= i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	return matrix
}

func ten(matrix [][]int) [][]int {
	length := len(matrix)
	for i := 0; i < length/2; i++ {
		matrix[i], matrix[length-i-1] = matrix[length-i-1], matrix[i]
	}
	for i := 0; i < length; i++ {
		for j := 0; j <= i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	return matrix
}

func eleven(matrix [][]int) [][]int {
	length := len(matrix)
	for i := 0; i < length/2; i++ {
		matrix[i], matrix[length-1-i] = matrix[length-1-i], matrix[i]
	}
	for i := 0; i < length; i++ {
		for j := 0; j <= i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	return matrix
}

func twelve(matrix [][]int) {
	length := len(matrix)
	for i := 0; i < len(matrix)/2; i++ {
		matrix[i], matrix[length-i-1] = matrix[length-1-i], matrix[i]
	}
	for i := 0; i < length; i++ {
		for j := 0; j <= i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

func thirteen(matrix [][]int) {
	length := len(matrix)
	for i := 0; i < len(matrix)/2; i++ {
		matrix[i], matrix[length-i-1] = matrix[length-i-1], matrix[i]
	}
	for i := 0; i < length; i++ {
		for j := 0; j <= i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

func fourteen(matrix [][]int) {
	length := len(matrix)
	for i := 0; i < len(matrix)/2; i++ {
		matrix[i], matrix[length-i-1] = matrix[length-i-1], matrix[i]
	}
	for i := 0; i < length; i++ {
		for j := 0; j <= i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

func fifteen(matrix [][]int) {
	length := len(matrix)
	for i := 0; i < len(matrix)/2; i++ {
		matrix[i], matrix[length-i-1] = matrix[length-i-1], matrix[i]
	}
	for i := 0; i < length; i++ {
		for j := 0; j <= i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
