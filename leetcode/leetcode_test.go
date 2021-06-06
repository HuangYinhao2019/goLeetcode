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