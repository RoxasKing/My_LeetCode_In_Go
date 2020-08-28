package main

/*
  在一个 8 x 8 的棋盘上，有一个白色车（rook）。也可能有空方块，白色的象（bishop）和黑色的卒（pawn）。它们分别以字符 “R”，“.”，“B” 和 “p” 给出。大写字符表示白棋，小写字符表示黑棋。
  车按国际象棋中的规则移动：它选择四个基本方向中的一个（北，东，西和南），然后朝那个方向移动，直到它选择停止、到达棋盘的边缘或移动到同一方格来捕获该方格上颜色相反的卒。另外，车不能与其他友方（白色）象进入同一个方格。
  返回车能够在一次移动中捕获到的卒的数量。
*/

func numRookCaptures(board [][]byte) int {
	var count int
	var row, col int
	forwards := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for i := range board {
		for j := range board[0] {
			if board[i][j] == 'R' {
				row, col = i, j
				goto THERE
			}
		}
	}
THERE:
	for _, forward := range forwards {
		for step := 0; ; step++ {
			r, c := row+step*forward[0], col+step*forward[1]
			if r < 0 || r > len(board)-1 || c < 0 || c > len(board[0])-1 || board[r][c] == 'B' {
				break
			} else if board[r][c] == 'p' {
				count++
				break
			}
		}
	}
	return count
}
