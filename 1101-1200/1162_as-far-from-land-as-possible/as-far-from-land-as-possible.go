package main

/*
  Given an n x n grid containing only values 0 and 1, where 0 represents water and 1 represents land, find a water cell such that its distance to the nearest land cell is maximized, and return the distance. If no land or water exists in the grid, return -1.

  The distance used in this problem is the Manhattan distance: the distance between two cells (x0, y0) and (x1, y1) is |x0 - x1| + |y0 - y1|.

  Example 1:
    Input: grid = [[1,0,1],[0,0,0],[1,0,1]]
    Output: 2
    Explanation: The cell (1, 1) is as far as possible from all the land with distance 2.

  Example 2:
    Input: grid = [[1,0,0],[0,0,0],[0,0,0]]
    Output: 4
    Explanation: The cell (2, 2) is as far as possible from all the land with distance 4.

  Constraints:
    1. n == grid.length
    2. n == grid[i].length
    3. 1 <= n <= 100
    4. grid[i][j] is 0 or 1

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/as-far-from-land-as-possible
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// BFS
func maxDistance(grid [][]int) int {
	forwards := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	n := len(grid)
	dist := make([][]int, n)
	q := [][]int{}
	for i := 0; i < n; i++ {
		dist[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				q = append(q, []int{i, j})
			} else {
				dist[i][j] = 1<<31 - 1
			}
		}
	}

	if len(q) == n*n || len(q) == 0 {
		return -1
	}

	for len(q) > 0 {
		x, y := q[0][0], q[0][1]
		q = q[1:]
		for _, f := range forwards {
			nx, ny := x+f[0], y+f[1]
			if nx < 0 || n-1 < nx || ny < 0 || n-1 < ny || grid[nx][ny] == 1 || dist[x][y]+1 >= dist[nx][ny] {
				continue
			}
			dist[nx][ny] = dist[x][y] + 1
			q = append(q, []int{nx, ny})
		}
	}

	out := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			out = Max(out, dist[i][j])
		}
	}
	return out
}

// Dynamic Programming
func maxDistance2(grid [][]int) int {
	n := len(grid)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				dp[i][j] = 1<<31 - 1
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				continue
			}
			if i-1 >= 0 && dp[i-1][j]+1 < dp[i][j] {
				dp[i][j] = dp[i-1][j] + 1
			}
			if j-1 >= 0 && dp[i][j-1]+1 < dp[i][j] {
				dp[i][j] = dp[i][j-1] + 1
			}
		}
	}

	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if grid[i][j] == 1 {
				continue
			}
			if i+1 < n && dp[i+1][j]+1 < dp[i][j] {
				dp[i][j] = dp[i+1][j] + 1
			}
			if j+1 < n && dp[i][j+1]+1 < dp[i][j] {
				dp[i][j] = dp[i][j+1] + 1
			}
		}
	}

	out := -1
	for i := range grid {
		for j := range grid {
			if grid[i][j] == 1 || dp[i][j] == 1<<31-1 {
				continue
			}
			out = Max(out, dp[i][j])
		}
	}
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
