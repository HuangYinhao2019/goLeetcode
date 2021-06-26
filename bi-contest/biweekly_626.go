package bi_contest

import "strings"

func canBeIncreasing(nums []int) bool {
	if Incre(nums) {
		return true
	}
	for i := range nums {
		arr := make([]int, 0)
		for j := range nums {
			if i != j {
				arr = append(arr, nums[j])
			}
		}
		if Incre(arr) {
			return true
		}
	}
	return false
}

func Incre(nums []int) bool {
	for i := 0; i < len(nums) - 1; i++ {
		if nums[i] >= nums[i + 1] {
			return false
		}
	}
	return true
}

func removeOccurrences(s string, part string) string {
	for l := strings.Index(s, part); l != -1; {
		s = s[:l] + s[l + len(part):]
		l = strings.Index(s, part)
	}
	return s
}

func maxAlternatingSum(nums []int) int64 {
	var l, r, sum int64 = -1, -1, 0
	for i := range nums {
		if l == -1 {
			l = int64(nums[i])
		} else {
			if r == -1 {
				if int64(nums[i]) < l {
					r = int64(nums[i])
				} else {
					l = int64(nums[i])
				}
			} else {
				if int64(nums[i]) <= r {
					r = int64(nums[i])
				} else {
					sum += l - r
					l = int64(nums[i])
					r = -1
				}
			}
		}
	}
	return sum + l
}