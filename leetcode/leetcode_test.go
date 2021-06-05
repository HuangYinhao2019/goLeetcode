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
