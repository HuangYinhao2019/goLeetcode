package contest

import "strconv"

func getConcatenation(nums []int) []int {
	l := len(nums)
	for i := 0; i < l; i++ {
		nums = append(nums, nums[i])
	}
	return nums
}

func countPalindromicSubsequence(s string) int {
	res := 0
	arr := make([][]int, 26)
	b := make([][]int, 26)
	for i := range b {
		b[i] = make([]int, 26)
	}
	for i := range arr {
		arr[i] = make([]int, 2)
		arr[i][0] = -1
		arr[i][1] = -1
	}
	for i := range s {
		if arr[s[i] - 'a'][0] == -1 {
			arr[s[i] - 'a'][0] = i
			arr[s[i] - 'a'][1] = i
		} else {
			if arr[s[i] - 'a'][0] != arr[s[i] - 'a'][1] {
				b[s[i] - 'a'][s[i] - 'a'] = 1
			}
			arr[s[i] - 'a'][1] = i
		}
		if arr[s[i] - 'a'][1] > arr[s[i] - 'a'][0] + 1 {
			for j := 0; j < 26; j++ {
				if (int)(s[i] - 'a') != j && arr[j][1] != -1 && arr[j][1] > arr[s[i] - 'a'][0] {
					b[s[i] - 'a'][j] = 1
				}
			}
		}
	}
	for i := range b {
		for j := range b[i] {
			res += b[i][j]
		}
	}
	return res
}

func colorTheGrid(m int, n int) int {
	qMap := make(map[string]bool)
	hMap := make(map[string][]string)
	pMap := make(map[string][]int64)
	buildQMap(m, qMap, "")
	for b := range qMap{
		pMap[b] = make([]int64, 2)
		pMap[b][0] = 1
		for b2 := range qMap {
			flag := true
			for i := 0; i < m; i++ {
				if b[i] == b2[i] {
					flag = false
					break
				}
			}
			if flag {
				if _, exist := hMap[b]; exist {
					hMap[b] = append(hMap[b], b2)
				} else {
					hMap[b] = []string{b2}
				}
			}
		}
	}
	for i := 1; i < n; i++ {
		for s := range qMap {
			cnt := pMap[s][(i + 1) % 2]
			for _, j := range hMap[s] {
				pMap[j][i % 2] = (pMap[j][i % 2] + cnt) % 1000000007
			}
			pMap[s][(i + 1) % 2] = 0
		}
	}
	var res int64 = 0
	for _, int64s := range pMap {
		res += int64s[0]
		res += int64s[1]
	}
	res = res % 1000000007
	return int(res)
}

func buildQMap(m int, q map[string]bool, st string) {
	if len(st) == m {
		q[st] = true
	} else {
		for i := 0; i < 3; i++ {
			if len(st) == 0 || strconv.Itoa(i) != st[len(st) - 1:] {
				str := st
				str += strconv.Itoa(i)
				buildQMap(m, q, str)
			}
		}
	}
}