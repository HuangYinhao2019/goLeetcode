package leetcode

import (
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
