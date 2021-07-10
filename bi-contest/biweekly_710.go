package bi_contest

func countTriples(n int) int {
	sum := 0
	for a := 1; a <= n; a++ {
		for b := 1; b <= n; b++ {
			for c := 1; c <= n; c++ {
				if a * a + b * b == c * c {
					sum++
				}
			}
		}
	}
	return sum
}

func nearestExit(maze [][]byte, entrance []int) int {
	visited := make([][]bool, len(maze))
	for i := range visited {
		visited[i] = make([]bool, len(maze[0]))
	}

	queue := make([][]int, 0)
	queue = append(queue, []int{entrance[0], entrance[1], 0})
	return bfs(maze, queue, visited)
}

func bfs(maze [][]byte, queue [][]int, visited [][]bool) int {
	now, cnt := 0, 1
	for now < cnt {
		x, y, s := queue[now][0], queue[now][1], queue[now][2]
		if (x < 0 || y < 0 || x >= len(maze) || y >= len(maze[0]) || visited[x][y] || maze[x][y] == '+') {
			now++
			continue
		} else if s != 0 && (x == 0 || y == 0 || x == len(maze) - 1 || y == len(maze[0]) - 1) {
			return s
		} else {
			visited[x][y] = true
			now++
			cnt += 4
			queue = append(queue, []int{x - 1, y, s + 1})
			queue = append(queue, []int{x + 1, y, s + 1})
			queue = append(queue, []int{x, y - 1, s + 1})
			queue = append(queue, []int{x, y + 1, s + 1})
		}
	}
	return -1
}

func sumGame(num string) bool {
	lc, rc, lsum, rsum := 0, 0, 0, 0
	for i := range num {
		if i < len(num) / 2 {
			if num[i] == '?' {
				lc++
			} else {
				lsum += int(num[i] - '0')
			}
		} else {
			if num[i] == '?' {
				rc++
			} else {
				rsum += int(num[i] - '0')
			}
		}
	}
	if (lc > rc && lsum >= rsum) || (lc < rc && lsum <= rsum) {
		return true
	} else {
		if lc < rc {
			diff := rc - lc
			return lsum != rsum + (diff / 2) * 9
		} else {
			diff := lc - rc
			return rsum != lsum + (diff / 2) * 9
		}
	}
}
