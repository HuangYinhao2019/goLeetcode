package leetcode

import "testing"

func TestWaysToSplit(t *testing.T) {
	if waysToSplit([]int{1, 2, 2, 2, 5, 0}) != 3 {
		t.Error("waysToSplit([]int{1, 2, 2, 2, 5, 0}) != 3")
	}
	if waysToSplit([]int{7, 0, 5}) != 0 {
		t.Error("waysToSplit([]int{7, 0, 5}) != 0")
	}
}

func TestConstructDistancedSequence(t *testing.T) {
	constructDistancedSequence(3)
}

func TestMinOperations(t *testing.T) {
	minOperations([]int{6, 4, 8, 1, 3, 2}, []int{4, 7, 6, 2, 3, 8, 6, 1})
}

func TestMinimumHammingDistance(t *testing.T) {
	minimumHammingDistance([]int{5, 1, 2, 4, 3}, []int{1, 5, 4, 2, 3}, [][]int{{0, 4}, {4, 2}, {1, 3}, {1, 4}})
}

func TestLargestSubmatrix(t *testing.T) {
	largestSubmatrix([][]int{{0,0,1},{1,1,1},{1,0,1}})
}

func TestKthLargestValue(t *testing.T) {
	kthLargestValue([][]int{{5,2},{1,6}}, 1)
}

func TestCheckPartitioning(t *testing.T) {
	checkPartitioning("abcbdd")
}

func TestMaxAbsoluteSum(t *testing.T) {
	maxAbsoluteSum([]int{1,-3,2,3,-4})
}

func TestMaxValue(t *testing.T) {
	maxValue([][]int{{30,40,34},{6,11,6},{60,81,36}}, 1)
}