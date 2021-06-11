package leetcode

import (
	"math"
	"sort"
)

//1710. 卡车上的最大单元数
func maximumUnits(boxTypes [][]int, truckSize int) int {
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})
	sum, left := 0, truckSize
	for _, c := range boxTypes {
		if left == 0 {
			break
		}
		if c[0] <= left {
			left -= c[0]
			sum += c[0] * c[1]
		} else {
			sum += left * c[1]
			left = 0
		}
	}
	return sum
}

//1711. 大餐计数
func countPairs(deliciousness []int) int {
	hMap := make(map[int]int)
	de := make([]int, 22)
	de[0] = 1
	for i := 1; i < 22; i++ {
		de[i] = de[i-1] * 2
	}
	var sum int = 0
	for _, v := range deliciousness {
		for j := 21; j >= 0; j-- {
			if de[j]-v >= 0 {
				if _, exist := hMap[de[j]-v]; exist {
					sum = (sum + hMap[de[j]-v]) % 1000000007
				}
			} else {
				break
			}
		}
		if _, exist := hMap[v]; exist {
			hMap[v]++
		} else {
			hMap[v] = 1
		}
	}
	return sum
}

//1712. 将数组分成三个子数组的方案数
func waysToSplit(nums []int) int {
	prefix := make([]int, len(nums))
	prefix[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		prefix[i] = prefix[i-1] + nums[i]
	}
	sum := 0
	for i, j, k := 0, 1, 1; i < len(nums)-1; i++ {
		if j <= i {
			j = i + 1
		}
		if k <= i {
			k = i + 1
		}
		for k < len(nums)-1 && prefix[len(nums)-1]-prefix[k] >= prefix[k]-prefix[i] {
			k++
		}
		for j < len(nums)-1 && prefix[i] > prefix[j]-prefix[i] {
			j++
		}
		if j <= k {
			sum = (sum + k - j) % 1000000007
		}
	}
	return sum
}

//1713. 得到子序列的最少操作次数
func minOperations(target []int, arr []int) int {
	hMap := make(map[int]int)
	for i, i2 := range target {
		hMap[i2] = i
	}
	ar := make([]int, 0)
	for _, i2 := range arr {
		if num, exist := hMap[i2]; exist {
			ar = append(ar, num)
		}
	}
	if len(ar) == 0 {
		return len(target)
	}
	d := make([]int, 1, len(ar))
	d[0] = ar[0]
	for _, i2 := range ar {
		if i2 > d[len(d)-1] {
			d = append(d, i2)
		} else {
			l, r := 0, len(d)-1
			for l < r {
				mid := l + (r-l)/2
				if d[mid] < i2 {
					l = mid + 1
				} else {
					r = mid
				}
			}
			d[l] = i2
		}
	}
	return len(target) - len(d)
}

//1716. 计算力扣银行的钱
func totalMoney(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += (i-1)%7 + (i-1)/7 + 1
	}
	return sum
}

//1717. 删除子字符串的最大得分
func maximumGain(s string, x int, y int) int {
	sum := 0
	if x > y {
		a, b := 0, 0
		for i, _ := range s {
			if s[i:i+1] == "a" {
				a++
			} else if s[i:i+1] == "b" {
				if a > 0 {
					a--
					sum += x
				} else {
					b++
				}
			} else {
				if a > 0 && b > 0 {
					if a > b {
						sum += y * b
					} else {
						sum += y * a
					}
				}
				a, b = 0, 0
			}
		}
		if a > 0 && b > 0 {
			if a > b {
				sum += y * b
			} else {
				sum += y * a
			}
		}
		return sum
	} else {
		a, b := 0, 0
		for i, _ := range s {
			if s[i:i+1] == "b" {
				b++
			} else if s[i:i+1] == "a" {
				if b > 0 {
					b--
					sum += y
				} else {
					a++
				}
			} else {
				if a > 0 && b > 0 {
					if a > b {
						sum += x * b
					} else {
						sum += x * a
					}
				}
				a, b = 0, 0
			}
		}
		if a > 0 && b > 0 {
			if a > b {
				sum += x * b
			} else {
				sum += x * a
			}
		}
		return sum
	}
}

//1718. 构建字典序最大的可行序列
func constructDistancedSequence(n int) []int {
	res := make([]int, n*2-1)
	hMap := make(map[int]bool)
	for i := 1; i <= n; i++ {
		hMap[i] = true
	}
	traceBack1718(hMap, res, 0, n)
	return res
}

func traceBack1718(hMap map[int]bool, res []int, now int, n int) bool {
	if now == len(res) {
		return true
	} else {
		if res[now] != 0 {
			return traceBack1718(hMap, res, now+1, n)
		} else {
			for i := n; i >= 2; i-- {
				if hMap[i] && now+i < len(res) && res[now+i] == 0 {
					hMap[i] = false
					res[now] = i
					res[now+i] = i
					if traceBack1718(hMap, res, now+1, n) {
						return true
					} else {
						res[now] = 0
						res[now+i] = 0
						hMap[i] = true
					}
				}
			}
			if hMap[1] {
				hMap[1] = false
				res[now] = 1
				if traceBack1718(hMap, res, now+1, n) {
					return true
				} else {
					hMap[1] = true
					res[now] = 0
				}
			}
			return false
		}
	}
}

//1720. 解码异或后的数组
func decode(encoded []int, first int) []int {
	res := make([]int, len(encoded)+1)
	res[0] = first
	for in, num := range encoded {
		res[in+1] = res[in] ^ num
	}
	return res
}

//1721. 交换链表中的节点
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapNodes(head *ListNode, k int) *ListNode {
	p := head
	a, b, l := 0, 0, 1
	for p.Next != nil {
		p = p.Next
		l++
	}
	p = head
	for i := 0; i < k-1; i++ {
		p = p.Next
	}
	a = p.Val
	q := head
	for i := 0; i < l-k; i++ {
		q = q.Next
	}
	b = q.Val
	p.Val, q.Val = b, a
	return head
}

//1722. 执行交换操作后的最小汉明距离
func minimumHammingDistance(source []int, target []int, allowedSwaps [][]int) int {
	len := len(source)
	h := make([]int, len)
	for i := 0; i < len; i++ {
		h[i] = i
	}
	for _, swap := range allowedSwaps {
		r1, r2 := swap[0], swap[1]
		for h[r1] != r1 {
			r1 = h[r1]
		}
		for h[r2] != r2 {
			r2 = h[r2]
		}
		if h[r1] != h[r2] {
			h[r1] = h[r2]
		}
	}
	sum := len
	hMap := make(map[int]map[int]int)
	for i := 0; i < len; i++ {
		hMap[i] = make(map[int]int)
	}
	for i := 0; i < len; i++ {
		lis := make([]int, 0)
		r := i
		for h[r] != r {
			lis = append(lis, r)
			r = h[r]
		}
		for _, li := range lis {
			h[li] = r
		}
	}
	for i := 0; i < len; i++ {
		if _, exist := hMap[h[i]][source[i]]; exist {
			hMap[h[i]][source[i]]++
		} else {
			hMap[h[i]][source[i]] = 1
		}
	}
	for i := 0; i < len; i++ {
		if nu, exist := hMap[h[i]][target[i]]; exist && nu > 0 {
			hMap[h[i]][target[i]]--
			sum--
		}
	}
	return sum
}

//1723. 完成所有工作的最短时间
func minimumTimeRequired(jobs []int, k int) int {
	n := len(jobs)
	m := 1 << n
	sum := make([]int, m)
	for i := 1; i < m; i++ {
		p, c := i, 0
		for p > 0 {
			if p % 2 == 1 {
				sum[i] += jobs[c]
			}
			p, c = p / 2, c + 1
		}
	}

	dp := make([][]int, k)
	for i := range dp {
		dp[i] = make([]int, m)
	}
	for i, s := range sum {
		dp[0][i] = s
	}

	for i := 1; i < k; i++ {
		for j := 0; j < (1 << n); j++ {
			minn := math.MaxInt64
			for x := j; x > 0; x = (x - 1) & j {
				minn = min(minn, max(dp[i-1][j-x], sum[x]))
			}
			dp[i][j] = minn
		}
	}
	return dp[k-1][(1<<n)-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//1727. 重新排列后的最大子矩阵
func largestSubmatrix(matrix [][]int) int {
	preGrid := make([][]int, len(matrix))
	n, m := len(matrix), len(matrix[0])
	res := 0
	for i := 0; i < n; i++ {
		preGrid[i] = make([]int, m)
	}
	for i := 0; i < m; i++ {
		preGrid[0][i] = matrix[0][i]
		for j := 1; j < n; j++ {
			if matrix[j][i] == 1 {
				preGrid[j][i] = preGrid[j - 1][i] + 1
			}
		}
	}
	for p := 0; p < n; p++ {
		sort.Slice(preGrid[p], func(i, j int) bool {
			return preGrid[p][i] > preGrid[p][j]
		})
		for q := 0; q < m; q++ {
			height := preGrid[p][q]
			res = max(res, height * (q + 1))
		}
	}
	return res
}