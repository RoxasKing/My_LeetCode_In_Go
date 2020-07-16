package leetcode

/*
  给定一个无向图graph，当这个图为二分图时返回true。

  如果我们能将一个图的节点集合分割成两个独立的子集A和B，并使图中的每一条边的两个节点一个来自A集合，一个来自B集合，我们就将这个图称为二分图。

  graph将会以邻接表方式给出，graph[i]表示图中与节点i相连的所有节点。每个节点都是一个在0到graph.length-1之间的整数。这图中没有自环和平行边： graph[i] 中不存在i，并且graph[i]中没有重复的值。

  注意:
    graph 的长度范围为 [1, 100]。
    graph[i] 中的元素的范围为 [0, graph.length - 1]。
    graph[i] 不会包含 i 或者有重复的值。
    图是无向的: 如果j 在 graph[i]里边, 那么 i 也会在 graph[j]里边。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/is-graph-bipartite
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Graph + DFS
func isBipartite(graph [][]int) bool {
	const UNCOLORED, RED, GREEN = 0, 1, 2
	colors := make([]int, len(graph))
	isValid := true
	var dfs func(int, int)
	dfs = func(node, color int) {
		colors[node] = color
		nextColor := RED
		if color == RED {
			nextColor = GREEN
		}
		for _, neighbor := range graph[node] {
			if colors[neighbor] == UNCOLORED {
				dfs(neighbor, nextColor)
				if !isValid {
					return
				}
			} else if colors[neighbor] != nextColor {
				isValid = false
				return
			}
		}
	}
	for i := 0; i < len(graph) && isValid; i++ {
		if colors[i] == UNCOLORED {
			dfs(i, RED)
		}
	}
	return isValid
}

// Graph + BFS
func isBipartite2(graph [][]int) bool {
	const UNCOLORED, RED, GREEN = 0, 1, 2
	colors := make([]int, len(graph))
	for i := range graph {
		if colors[i] == UNCOLORED {
			queue := []int{i}
			colors[i] = RED
			for len(queue) != 0 {
				node := queue[0]
				queue = queue[1:]
				nextColor := RED
				if colors[node] == RED {
					nextColor = GREEN
				}
				for _, neighbor := range graph[node] {
					if colors[neighbor] == UNCOLORED {
						queue = append(queue, neighbor)
						colors[neighbor] = nextColor
					} else if colors[neighbor] != nextColor {
						return false
					}
				}
			}
		}
	}
	return true
}
