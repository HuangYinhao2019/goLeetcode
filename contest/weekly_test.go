package contest

import "testing"

func TestRotateGrid(t *testing.T) {
	rotateGrid([][]int{{40,10},{30,20}}, 1)
}

func TestWonderfulSubstrings(t *testing.T) {
	wonderfulSubstrings("aba")
}

func TestColorTheGrid(t *testing.T) {
	colorTheGrid(5, 5)
}

func TestMaxCompatibilitySum(t *testing.T) {
	maxCompatibilitySum([][]int{{0,0,1,1,1,0,1},{0,1,1,0,0,0,0},{0,0,1,1,1,1,1},{0,1,0,0,1,0,1},{1,0,1,1,1,1,1}},
		[][]int{{0,1,1,0,0,0,0},{0,1,0,0,0,0,1},{0,1,0,1,0,0,1},{1,0,0,0,1,0,1},{1,1,1,1,1,0,0}})
}