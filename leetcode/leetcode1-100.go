package leetcode

//90. 子集 II
func SubsetsWithDup(nums []int) [][]int {
	h := make(map[int]int, 30)
	for _, num := range nums {
		if _, v := h[num]; v {
			h[num]++
		} else {
			h[num] = 1
		}
	}
	var s []int
	sum := 1
	for i, n := range h {
		s = append(s, i)
		sum *= (n + 1)
	}
	res := make([][]int, 0)
	arr := make([]int, 0)
	dfs90(h, s, 0, &res, &arr)
	return res
}

func dfs90(m map[int]int, s []int, now int, res *[][]int, arr *[]int) {
	if now >= len(s) {
		_arr := make([]int, len(*arr))
		copy(_arr, *arr)
		*res = append(*res, _arr)
	} else {
		for i := 0; i <= m[s[now]]; i++ {
			_arr := make([]int, len(*arr))
			copy(_arr, *arr)
			for j := 0; j < i; j++ {
				_arr = append(_arr, s[now])
			}
			dfs90(m, s, now+1, res, &_arr)
		}
	}
}
