package isMatch

func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	isMatch := func(i, j int) bool {
		if i == 0 {
			return false
		}

		if p[j-1] == '.' {
			return true
		}
		return s[i-1] == p[j-1]
	}

	f := make([][]bool, m+1)
	for i, _ := range f {
		f[i] = make([]bool, n+1)
	}

	f[0][0] = true
	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				f[i][j] = f[i][j] || f[i][j-2]
				if isMatch(i, j-1) {
					f[i][j] = f[i][j] || f[i-1][j]
				}
			} else if isMatch(i, j) {
				f[i][j] = f[i][j] || f[i-1][j-1]
			}
		}
	}
	return f[m][n]
}
