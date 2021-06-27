package contest

import (
	"sort"
)

func maxProductDifference(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums) - 1] * nums[len(nums) - 2] - (nums[0] * nums[1])
}

func rotateGrid(grid [][]int, k int) [][]int {
	n, m := len(grid), len(grid[0])
	p := n
	if n > m {
		p = m
	}
	for i := 0; i < p / 2; i++ {
		rotate(grid, i, k)
	}
	return grid
}

func rotate(grid [][]int, layer, k int) {
	n, m := len(grid), len(grid[0])
	s := (n - 2 * layer) * (m - 2 * layer) - ((n - 2 * layer - 2) * (m - 2 * layer - 2))
	k = k % s
	arr, c := make([]int, s), 0
	for i := layer; i < m - layer - 1; i++ {
		arr[c] = grid[layer][i]
		c++
	}
	for i := layer; i < n - layer - 1; i++ {
		arr[c] = grid[i][m - layer - 1]
		c++
	}
	for i := m - layer - 1; i > layer ; i-- {
		arr[c] = grid[n - layer - 1][i]
		c++
	}
	for i := n - layer - 1; i > layer ; i-- {
		arr[c] = grid[i][layer]
		c++
	}
	c = 0
	for i := layer; i < m - layer - 1; i++ {
		grid[layer][i] = arr[(c + k) % s]
		c++
	}
	for i := layer; i < n - layer - 1; i++ {
		grid[i][m - layer - 1] = arr[(c + k) % s]
		c++
	}
	for i := m - layer - 1; i > layer ; i-- {
		grid[n - layer - 1][i] = arr[(c + k) % s]
		c++
	}
	for i := n - layer - 1; i > layer ; i-- {
		grid[i][layer] = arr[(c + k) % s]
		c++
	}
}

func wonderfulSubstrings(word string) int64 {
	dp := make([][]int64, 2)
	h := make(map[int]bool)
	for i := 1; i < 2048; i = i * 2 {
		h[i] = true
	}
	h[0] = true
	var res int64 = 1
	for i := range dp {
		dp[i] = make([]int64, 1024)
	}
	dp[0][1 << (word[0] - 'a')] = 1
	for i := 1; i < len(word); i++ {
		p := word[i] - 'a'
		for j := 0; j < 1024; j++ {
			dp[i % 2][j] = dp[(i + 1) % 2][(1 << p) ^ j]
			if j == 1 << p {
				dp[i % 2][1 << p] += 1
			}
			if h[j] {
				res += dp[i % 2][j]
			}
		}
	}
	return res
}

