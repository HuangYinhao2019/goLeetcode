package bi_contest

import (
	"container/heap"
	"fmt"
	"sort"
	"strings"
)

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

type Node struct {
	movie int
	shop int
	price int
	isRent bool
}

type H []*Node

func (h H) Len() int {
	return len(h)
}
func (h H) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h H) Less(i, j int) bool {
	if h[i].price > h[j].price {
		return false
	}
	if h[i].price < h[j].price {
		return true
	}
	if h[i].shop > h[j].shop {
		return false
	}
	if h[i].shop < h[j].shop {
		return true
	}
	return h[i].movie < h[j].movie
}

func (h *H) Push(v interface{}) {
	*h = append(*h, v.(*Node))
}

func (h *H) Pop() interface{} {
	res := (*h)[len(*h) - 1]
	*h = (*h)[:len(*h) - 1]
	return res
}

type MovieRentingSystem struct {
	Map  map[string]*Node
	H    H
	rent map[int][]*Node
}


func Constructor(n int, entries [][]int) MovieRentingSystem {
	Map := make(map[string]*Node)
	H := make([]*Node, 0)
	rent := make(map[int][]*Node)

	for i := range entries {
		node := Node{
			shop: entries[i][0],
			movie: entries[i][1],
			price: entries[i][2],
			isRent: false,
		}
		key := fmt.Sprint("%v_%v", entries[i][0], entries[i][1])
		Map[key] = &node
		rent[entries[i][1]] = append(rent[entries[i][1]], &node)
	}

	for i := range rent {
		sort.Slice(rent[i], func(a, b int) bool {
			if rent[i][a].price < rent[i][b].price {
				return true
			}
			if rent[i][a].price > rent[i][b].price {
				return false
			}
			return rent[i][a].shop < rent[i][b].shop
		})
	}

	return MovieRentingSystem{Map, H, rent}
}


func (this *MovieRentingSystem) Search(movie int) []int {
	nodes := this.rent[movie]
	res := make([]int, 0)
	for i := range nodes {
		if nodes[i].isRent {
			continue
		}
		res = append(res, nodes[i].shop)
		if len(res) >= 5 {
			break
		}
	}
	return res
}


func (this *MovieRentingSystem) Rent(shop int, movie int)  {
	key := fmt.Sprint("%v_%v", shop, movie)
	this.Map[key].isRent = true
	heap.Push(&this.H, this.Map[key])
}

func (this *MovieRentingSystem) Drop(shop int, movie int)  {
	key := fmt.Sprint("%v_%v", shop, movie)
	this.Map[key].isRent = false
}

func (this *MovieRentingSystem) Report() [][]int {
	hMap := make(map[string]bool, 0)
	res := make([][]int, 0)
	tmp := make([]*Node, 0)

	for this.H.Len() > 0 {
		node := heap.Pop(&this.H).(*Node)
		key := fmt.Sprint("%v_%v", node.shop, node.movie)
		if !node.isRent {
			continue
		}
		if hMap[key] {
			continue
		}
		hMap[key] = true
		res = append(res, []int{node.shop, node.movie})
		tmp = append(tmp, node)
		if len(res) >= 5 {
			break
		}
	}
	for i := range tmp {
		heap.Push(&this.H, tmp[i])
	}
	return res
}


/**
 * Your MovieRentingSystem object will be instantiated and called as such:
 * obj := Constructor(n, entries);
 * param_1 := obj.Search(movie);
 * obj.Rent(shop,movie);
 * obj.Drop(shop,movie);
 * param_4 := obj.Report();
 */