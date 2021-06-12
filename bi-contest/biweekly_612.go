package bi_contest

func isCovered(ranges [][]int, left int, right int) bool {
	hMap := make(map[int]bool)
	for _, ints := range ranges {
		for i := ints[0]; i <= ints[1]; i++ {
			hMap[i] = true
		}
	}
	for i := left; i <= right; i++ {
		if !hMap[i] {
			return false
		}
	}
	return true
}

func chalkReplacer(chalk []int, k int) int {
	var sum int64
	sum = 0
	for _, i2 := range chalk {
		sum += int64(i2)
	}
	k = int(int64(k) % sum)
	for i, i2 := range chalk {
		if k < i2 {
			return i
		}
		k -= i2
	}
	return 0
}

func largestMagicSquare(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	pre := make([][][]int, n + 1)
	for i := 0; i < n + 1; i++ {
		pre[i] = make([][]int, m + 1)
		for j := 0; j <= m; j++ {
			pre[i][j] = make([]int, 4)
		}
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			pre[i][j][0] = pre[i - 1][j][0] + grid[i - 1][j - 1]
		}
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			pre[i][j][1] = pre[i][j - 1][1] + grid[i - 1][j - 1]
		}
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			pre[i][j][2] = pre[i - 1][j - 1][2] + grid[i - 1][j - 1]
		}
	}
	for i := 1; i <= n; i++ {
		for j := m - 1; j >= 0; j-- {
			pre[i][j][3] = pre[i - 1][j + 1][3] + grid[i - 1][j]
		}
	}
	for p := min(n ,m); p >= 2; p-- {
		for i := 0; i + p <= n; i++ {
			for j := 0; j + p <= m; j++ {
				flag := true
				s := -1
				if pre[i + p][j + p][2] - pre[i][j][2] == pre[i + p][j][3] - pre[i][j + p][3] {
					s = pre[i + p][j + p][2] - pre[i][j][2]
				} else {
					continue
				}
				for q := 1; q <= p; q++ {
					if pre[i + p][j + q][0] - pre[i][j + q][0] != s || pre[i + q][j + p][1] - pre[i + q][j][1] != s{
						flag = false
						break
					}
				}
				if flag {
					return p
				}
			}
		}
	}
	return 1
}

func min (a, b int) int {
	if a < b {
		return a
	}
	return b
}