package contest

import "sort"

func leastMinutes(n int) int {
	t := 0
	res := n
	for k := 1; k < n; k = k * 2 {
		res = min(res, n / k + y(n, k) + t)
		t++
	}
	return res
}

func y(n, k int) int {
	if n % k != 0 {
		return 1
	}
	return 0
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func halfQuestions(questions []int) int {
	N := len(questions) / 2
	hmap := make(map[int]int)
	for _, question := range questions {
		if _, ok := hmap[question]; ok {
			hmap[question]++
		} else {
			hmap[question] += 1
		}
	}
	ks := make([]int, len(hmap))
	i := 0
	for _, i2 := range hmap {
		ks[i] = i2
		i++
	}
	sort.Ints(ks)
	i = len(ks) - 1
	for N > 0 {
		N -= ks[i]
		i--
	}
	return len(ks) - 1 - i
}

func largestArea(grid []string) int {
	n, m := len(grid), len(grid[0])
	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m)
	}
	res := 0
	for i := 1; i < n - 1; i++ {
		for j := 1; j < m - 1; j++ {
			if grid[i][j:j+1] == "0" {
				continue
			}
			if nu, ok := dfs(i, j, n, m, visited, grid, grid[i][j:j+1]); ok {
				res = max(res, nu)
			}
		}
	}
	return res
}

func dfs (x, y, n, m int, visited [][]bool, grid []string, st string) (int, bool) {
	if x < 0 || y < 0 || x >= n || y >= m || grid[x][y:y+1] == "0" {
		return 0, false
	}
	if visited[x][y] || grid[x][y:y+1] != st {
		return 0, true
	}
	visited[x][y] = true
	a1, b1 := dfs(x + 1, y, n, m, visited, grid, st)
	a2, b2 := dfs(x, y + 1, n, m, visited, grid, st)
	a3, b3 := dfs(x - 1, y, n, m, visited, grid, st)
	a4, b4 := dfs(x, y - 1, n, m, visited, grid, st)
	return (1 + a1 + a2 + a3 + a4), b1 && b2 && b3 && b4
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}