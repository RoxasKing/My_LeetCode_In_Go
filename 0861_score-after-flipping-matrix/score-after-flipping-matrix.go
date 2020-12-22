package main

/*
  有一个二维矩阵 A 其中每个元素的值为 0 或 1 。
  移动是指选择任一行或列，并转换该行或列中的每一个值：将所有 0 都更改为 1，将所有 1 都更改为 0。
  在做出任意次数的移动后，将该矩阵的每一行都按照二进制数来解释，矩阵的得分就是这些数字的总和。
  返回尽可能高的分数。

  示例：
    输入：[[0,0,1,1],[1,0,1,0],[1,1,0,0]]
    输出：39
    解释：
    转换为 [[1,1,1,1],[1,0,0,1],[1,1,1,1]]
    0b1111 + 0b1001 + 0b1111 = 15 + 9 + 15 = 39

  提示：
    1 <= A.length <= 20
    1 <= A[0].length <= 20
    A[i][j] 是 0 或 1

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/score-after-flipping-matrix
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Dynamic Programming
func matrixScore(A [][]int) int {
	m, n := len(A), len(A[0])
	count := make([]int, n)
	for i := range A {
		for j := range A[i] {
			if A[i][j] == 1 {
				count[j]++
			}
		}
	}
	for i := range A {
		if A[i][0] == 1 {
			continue
		}
		for j := range A[i] {
			if A[i][j] == 0 {
				A[i][j] = 1
				count[j]++
			} else {
				A[i][j] = 0
				count[j]--
			}
		}
	}
	for j := range count {
		if count[j] >= (m+1)>>1 {
			continue
		}
		for i := range A {
			if A[i][j] == 0 {
				A[i][j] = 1
			} else {
				A[i][j] = 0
			}
		}
		count[j] = m - count[j]
	}
	out := 0
	for j := range count {
		out += count[j] * (1 << (n - 1 - j))
	}
	return out
}
