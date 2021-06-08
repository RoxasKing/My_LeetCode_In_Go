package main

/*
  There is a tree (i.e., a connected, undirected graph that has no cycles) consisting of n nodes numbered from 0 to n - 1 and exactly n - 1 edges. Each node has a value associated with it, and the root of the tree is node 0.

  To represent this tree, you are given an integer array nums and a 2D array edges. Each nums[i] represents the ith node's value, and each edges[j] = [uj, vj] represents an edge between nodes uj and vj in the tree.

  Two values x and y are coprime if gcd(x, y) == 1 where gcd(x, y) is the greatest common divisor of x and y.

  An ancestor of a node i is any other node on the shortest path from node i to the root. A node is not considered an ancestor of itself.

  Return an array ans of size n, where ans[i] is the closest ancestor to node i such that nums[i] and nums[ans[i]] are coprime, or -1 if there is no such ancestor.

  Example 1:
    Input: nums = [2,3,3,2], edges = [[0,1],[1,2],[1,3]]
    Output: [-1,0,0,1]
    Explanation: In the above figure, each node's value is in parentheses.
      - Node 0 has no coprime ancestors.
      - Node 1 has only one ancestor, node 0. Their values are coprime (gcd(2,3) == 1).
      - Node 2 has two ancestors, nodes 1 and 0. Node 1's value is not coprime (gcd(3,3) == 3), but node 0's
        value is (gcd(2,3) == 1), so node 0 is the closest valid ancestor.
      - Node 3 has two ancestors, nodes 1 and 0. It is coprime with node 1 (gcd(3,2) == 1), so node 1 is its
        closest valid ancestor.

  Example 2:
    Input: nums = [5,6,10,2,3,6,15], edges = [[0,1],[0,2],[1,3],[1,4],[2,5],[2,6]]
    Output: [-1,0,-1,0,0,0,-1]

  Constraints:
    1. nums.length == n
    2. 1 <= nums[i] <= 50
    3. 1 <= n <= 10^5
    4. edges.length == n - 1
    5. edges[j].length == 2
    6. 0 <= uj, vj < n
    7. uj != vj

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/tree-of-coprimes
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// BFS
// Math

func getCoprimes(nums []int, edges [][]int) []int {
	n := len(nums)
	es := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		es[u] = append(es[u], v)
		es[v] = append(es[v], u)
	}

	out := make([]int, n)
	dep := make([]int, n)
	for i := range out {
		out[i] = -1
		dep[i] = -1
	}

	visited := make([]bool, n)
	parents := make([][51][2]int, n)
	for i := range parents {
		parents[i] = [51][2]int{}
		for j := 1; j <= 50; j++ {
			parents[i][j] = [2]int{-1, -1}
		}
	}

	visited[0] = true
	q := [][2]int{{0, 0}}
	for len(q) > 0 {
		u, depth := q[0][0], q[0][1]
		q = q[1:]
		for i := 1; i <= 50; i++ {
			if parents[u][i][0] != -1 && Gcd(nums[u], i) == 1 && (parents[u][i][1] > dep[u]) {
				out[u] = parents[u][i][0]
				dep[u] = parents[u][i][1]
			}
		}
		parents[u][nums[u]] = [2]int{u, depth}
		for _, v := range es[u] {
			if visited[v] {
				continue
			}
			visited[v] = true
			parents[v] = parents[u]
			q = append(q, [2]int{v, depth + 1})
		}
	}

	return out
}

func Gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}