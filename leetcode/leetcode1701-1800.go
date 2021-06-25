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

//1732. 找到最高海拔
func largestAltitude(gain []int) int {
	mmax, now := 0, 0
	for _, v := range gain {
		now = now + v
		mmax = max(now ,mmax)
	}
	return mmax
}

//1733. 需要教语言的最少人数
func minimumTeachings(n int, languages [][]int, friendships [][]int) int {
	p := len(languages)
	mmin := p
	lMap := make(map[int]map[int]bool)
	fMap := make([][]int, 0, len(friendships))
	for i := 0; i < p; i++ {
		lMap[i + 1] = make(map[int]bool)
		for _, i3 := range languages[i] {
			lMap[i + 1][i3] = true
		}
	}
	for _, friendship := range friendships {
		a, b := friendship[0], friendship[1]
		flag := false
		for b2 := range lMap[a] {
			if lMap[b][b2] {
				flag = true
				break
			}
		}
		if !flag {
			fMap = append(fMap, []int{a, b})
		}
	}
	for i := 1; i <= n; i++ {
		sMap := make(map[int]bool)
		for _, friendship := range fMap {
			a, b := friendship[0], friendship[1]
			if v, ok := lMap[a][i]; !ok {
				sMap[a] = v
			}
			if v, ok := lMap[b][i]; !ok {
				sMap[b] = v
			}
		}
		mmin = min(mmin, len(sMap))
	}
	return mmin
}

//1734. 解码异或后的排列
func decode2(encoded []int) []int {
	n := len(encoded) + 1
	xor := 1
	for i := 2; i <= n; i++ {
		xor = xor ^ i
	}
	res := make([]int, n)
	res[0] = xor
	for i := 1; i < len(encoded); i += 2 {
		res[0] = res[0] ^ encoded[i]
	}
	for i := 1; i < n; i++ {
		res[i] = res[i - 1] ^ encoded[i - 1]
	}
	return res
}

//1736. 替换隐藏数字得到的最晚时间
func maximumTime(time string) string {
	l,r := time[:2],time[3:]
	res := ""
	if l[0] == '?'{
		if l[1] >= '4' && l[1] != '?'{
			res += "1"
		}else{
			res += "2"
		}

	}else{
		res += string(l[0])
	}
	if l[1] == '?'{
		if res[0] == '0' || res[0] == '1'{
			res += "9"
		}else{
			res += "3"
		}
	}else{
		res += string(l[1])
	}
	res+=":"
	if r[0] == '?'{
		res += "5"
	}else{
		res += string(r[0])
	}
	if r[1] == '?'{
		res += "9"
	}else{
		res += string(r[1])
	}
	return res
}

//1737. 满足三条件之一需改变的最少字符数
func minCharacters(a string, b string) int {
	aa, bb := make([]int, 26), make([]int, 26)
	for _, i2 := range a {
		aa[i2 - 'a']++
	}
	for _, i2 := range b {
		bb[i2 - 'a']++
	}
	asum, bsum := len(a), len(b)
	res, pa, pb := math.MaxInt32, 0, 0
	for i := 0; i < 25; i++ {
		pa += aa[i]
		pb += bb[i]
		res = min(min(res, asum + bsum - aa[i] - bb[i]), min(asum - pa + pb, bsum - pb + pa))
	}
	res = min(res, asum + bsum - aa[25] - bb[25])
	return res
}

//1738. 找出第 K 大的异或坐标值
func kthLargestValue(matrix [][]int, k int) int {
	m, n := len(matrix), len(matrix[0])
	kres := make([]int, 0)
	for _, ints := range matrix {
		for i, _ := range ints {
			if i > 0 {
				ints[i] = ints[i] ^ ints[i - 1]
			}
		}
	}
	for j := 0; j < n; j++ {
		for i := 1; i < m; i++ {
			matrix[i][j] = matrix[i][j] ^ matrix[i - 1][j]
		}
	}
	for _, ints := range matrix {
		for _, i3 := range ints {
			kres = append(kres, i3)
		}
	}
	sort.Ints(kres)
	return kres[m * n - k]
}

//1739. 放置盒子
func minimumBoxes(n int) int {
	height, bottom, total := 0, 0, 0
	for total < n {
		height++
		bottom = height * (height + 1) / 2
		total += bottom
	}
	if total > n{
		height--
		total -= bottom
		bottom = height * (height + 1) / 2
		for i := 1; i <= height + 1; i, bottom = i + 1, bottom + 1 {
			if total >= n {
				return bottom
			}
			total += i
		}
	}
	return bottom
}

//1742. 盒子中小球的最大数量
func countBalls(lowLimit int, highLimit int) int {
	hMap := make(map[int]int)
	for i := lowLimit; i <= highLimit; i++ {
		sum := 0
		p := i
		for p > 0 {
			sum += p % 10
			p /= 10
		}
		if _, exist := hMap[sum]; exist {
			hMap[sum]++
		} else {
			hMap[sum] = 1
		}
	}
	res := 0
	for _, i2 := range hMap {
		res = max(res, i2)
	}
	return res
}

//1743. 从相邻元素对还原数组
func restoreArray(adjacentPairs [][]int) []int {
	hMap := make(map[int]map[int]bool)
	for _, pair := range adjacentPairs {
		if _, ok := hMap[pair[0]]; !ok {
			hMap[pair[0]] = make(map[int]bool)
		}
		if _, ok := hMap[pair[1]]; !ok {
			hMap[pair[1]] = make(map[int]bool)
		}
		hMap[pair[0]][pair[1]] = true
		hMap[pair[1]][pair[0]] = true
	}
	res := make([]int, len(adjacentPairs) + 1)
	for i, ints := range hMap {
		if len(ints) == 1 {
			res[0] = i
		}
	}
	for i := 1; i < len(res); i++ {
		l := res[i - 1]
		for j, b := range hMap[l] {
			if b {
				res[i] = j
				b = false
				hMap[j][l] = false
			}
		}
	}
	return res
}

//1744. 你能在你最喜欢的那天吃到你最喜欢的糖果吗？
func canEat(candiesCount []int, queries [][]int) []bool {
	prefix := make([]int, len(candiesCount) + 1)
	for i, _ := range candiesCount {
		prefix[i + 1] = prefix[i] + candiesCount[i]
	}
	ans := make([]bool, len(queries))
	for i, query := range queries {
		cType, day, m := query[0], query[1], query[2]
		a, b := prefix[cType] + 1, prefix[cType + 1]
		l, r := day + 1, m * (day + 1)
		if (l <= b && l >= a) || (r <= b && r >= a) || (l <= a && r >= b){
			ans[i] = true
		}
	}
	return ans
}

//1745. 回文串分割 IV
func checkPartitioning(s string) bool {
	hMap := make(map[int]map[int]bool)
	for i := 0; i < len(s); i++ {
		hMap[i] = map[int]bool{}
		hMap[i][i] = true
	}
	dp := make([][]bool, len(s) + 1)
	for i, _ := range dp {
		dp[i] = make([]bool, 4)
	}
	dp[1][1] = true
	for i := 1; i < len(s); i++ {
		dp[i + 1][2] = dp[i][1]
		dp[i + 1][3] = dp[i][2]
		for j := 0; j < i; j++ {
			if s[i] == s[j] {
				if j == i - 1 || hMap[j + 1][i - 1] {
					hMap[j][i] = true
					if j == 0 {
						dp[i + 1][1] = true
					} else {
						dp[i + 1][2] = dp[i + 1][2] || dp[j][1]
						dp[i + 1][3] = dp[i + 1][3] || dp[j][2]
					}
				}
			}
		}
	}
	return dp[len(s)][3] || dp[len(s)][1]
}

//1748. 唯一元素的和
func sumOfUnique(nums []int) int {
	sum := 0
	hMap := make(map[int]int)
	for _, num := range nums {
		if _, ok := hMap[num]; ok {
			hMap[num]++
		} else {
			hMap[num] = 1
		}
	}
	for _, num := range nums {
		if hMap[num] == 1 {
			sum += num
		}
	}
	return sum
}

//1749. 任意子数组和的绝对值的最大值
func maxAbsoluteSum(nums []int) int {
	lsum, rsum := 0, 0
	lres, rres := 0, 0
	for i, j := 0, 0; j < len(nums); {
		for lsum < 0 && i < j {
			lsum -= nums[i]
			i++
			lres = max(lsum, lres)
		}
		lsum += nums[j]
		j++
		lres = max(lsum, lres)
		if j == len(nums) {
			for i < len(nums) {
				lsum -= nums[i]
				i++
				lres = max(lsum, lres)
			}
		}
	}
	for i, j := 0, 0; j < len(nums); {
		for rsum > 0 && i < j {
			rsum -= nums[i]
			i++
			rres = min(rsum, rres)
		}
		rsum += nums[j]
		j++
		rres = min(rsum, rres)
		if j == len(nums) {
			for i < len(nums) {
				rsum -= nums[i]
				i++
				rres = min(rsum, rres)
			}
		}
	}
	return max(lres, -rres)
}

//1750. 删除字符串两端相同字符后的最短长度
func minimumLength(s string) int {
	l, r := 0, len(s) - 1
	for l < r && s[l] == s[r] {
		for l < r && s[l] == s[l + 1] {
			l++
		}
		for l < r && s[r] == s[r - 1] {
			r--
		}
		l++
		r--
	}
	return max(r - l + 1, 0)
}

//1751. 最多可以参加的会议数目 II
func maxValue(events [][]int, k int) int {
	sort.Slice(events, func(i, j int) bool {
		return events[i][1] < events[j][1]
	})
	dp := make([][]int, len(events))
	for i := range dp {
		dp[i] = make([]int, k + 1)
		for j := range dp[i] {
			dp[i][j] = math.MinInt32
		}
		dp[i][0] = 0
	}
	dp[0][1] = events[0][2]
	for i := 1; i <= k; i++ {
		for j := 1; j < len(events); j++ {
			dp[j][i] = dp[j - 1][i]
			be := events[j][0]
			l, r := 0, j - 1
			for l < r {
				mid := (l + r + 1) >> 1
				if events[mid][1] < be {
					l = mid
				} else {
					r = mid - 1
				}
			}
			if events[l][1] >= be {
				dp[j][1] = max(dp[j][1], events[j][2])
			} else {
				dp[j][i] = max(dp[j][i], dp[l][i - 1] + events[j][2])
			}
		}
	}
	res := 0
	for p := 1; p <= k; p++ {
		res = max(res, dp[len(events) - 1][p])
	}
	return res
}

//1752. 检查数组是否经排序和轮转得到
func check(nums []int) bool {
	k := 0
	for i, _ := range nums {
		if nums[i] < nums[(i - 1 + len(nums)) % len(nums)] {
			k++
		}
	}
	return k <= 1
}

//1753. 移除石子的最大得分
func maximumScore(a int, b int, c int) int {
	round := (a + b + c)/2
	if a < b {
		a, b = b, a
	}
	if a < c && a + b < round {
		round = a + b
	} else if a >= c && b + c < round {
		round = b + c
	}

	return  round
}

//1754. 构造字典序最大的合并字符串
func largestMerge(word1 string, word2 string) string {
	var merge string = ""
	for i, j := 0, 0; i < len(word1) || j < len(word2); {
		if getFromString(word1[i:], word2[j:]) {
			merge += word1[i:i+1]
			i++
		} else {
			merge += word2[j:j+1]
			j++
		}
	}
	return merge
}

func getFromString(a, b string) bool {
	for i := range a {
		if i + 1 > len(b) || a[i] > b[i] {
			return true
		} else if b[i] > a[i] {
			return false
		}
	}
	return false
}

//1771. 由子序列构造的最长回文串的长度
func longestPalindrome(word1 string, word2 string) int {
	word := word1 + word2
	dp := make([][]int, len(word))
	for i := range dp {
		dp[i] = make([]int, len(word))
		dp[i][i] = 1
	}
	res := 0
	for j := 1; j < len(word); j++ {
		for i := 0; i + j < len(word); i++ {
			if word[i] == word[i + j] {
				dp[i][i + j] = dp[i + 1][i + j - 1] + 2
				if i < len(word1) && i + j >= len(word1) {
					res = max(res, dp[i][i + j])
				}
			} else {
				dp[i][i + j] = max(dp[i + 1][i + j], dp[i][i + j - 1])
			}
		}
	}
	return res
}





