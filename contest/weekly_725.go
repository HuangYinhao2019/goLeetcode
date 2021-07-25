package contest

func getLucky(s string, k int) int {
	sum := 0
	for i := range s {
		k := int(s[i] - 'a' + 1)
		ss := k / 10 + k % 10
		sum += ss
	}
	for i := 1; i < k; i++ {
		p := 0
		for sum > 0 {
			p += sum % 10
			sum /= 10
		}
		sum = p
	}
	return sum
}

func maximumNumber(num string, change []int) string {
	flag := false
	nums := []byte(num)
	for i := range nums {
		be := int(nums[i] - '0')
		af := change[be]
		if af < be && flag {
			return string(nums)
		} else if af > be {
			nums[i] = byte(af + '0')
			flag = true
		}
	}
	return string(nums)
}

func maxCompatibilitySum(students [][]int, mentors [][]int) int {
	n, m := len(students[0]), len(students)
	grid := make([][]int, m)
	for i := range grid {
		grid[i] = make([]int, m)
	}
	for i := range students {
		for j := range mentors {
			s := 0
			for k := 0; k < n; k++ {
				if students[i][k] == mentors[j][k] {
					s++
				}
			}
			grid[i][j] = s
		}
	}
	return dfs4maxCompatibilitySum(grid, make(map[int]int, n), 0)
}

func dfs4maxCompatibilitySum(grid [][]int, s2m map[int]int, n int) int {
	if n >= len(grid) {
		sum := 0
		for i, j := range s2m {
			sum += grid[i][j]
		}
		return sum
	} else {
		sum := 0
		for i := 0; i < len(grid); i++ {
			if _, exist := s2m[i]; !exist {
				s2m[i] = n
				sum = max(sum, dfs4maxCompatibilitySum(grid, s2m, n + 1))
				delete(s2m, i)
			}
		}
		return sum
	}
}

