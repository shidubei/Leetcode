package array

func solveSudoku(board [][]byte) {
	//行列的已出现数字
	var line, column [9][9]bool
	//块的已出现数字
	var block [3][3][9]bool
	//空格的位置
	var spaces [][2]int

	for i, row := range board {
		for j, b := range row {
			if b == '.' {
				//记录空格位
				spaces = append(spaces, [2]int{i, j})
			} else {
				index := b - '1'
				//记录每行每列每块已出现数字
				line[i][index] = true
				column[j][index] = true
				block[i/3][j/3][index] = true
			}
		}
	}
	//重点理解
	var dfs func(int) bool
	dfs = func(pos int) bool {
		if pos == len(spaces) {
			//找完了空格
			return true
		}
		//开始遍历所有的空格
		i, j := spaces[pos][0], spaces[pos][1]
		for index := byte(0); index < 9; index++ {
			if !line[i][index] && !column[j][index] && !block[i/3][j/3][index] {
				//更新出现的数字
				line[i][index] = true
				column[j][index] = true
				block[i/3][j/3][index] = true
				//填入数字
				board[i][j] = index + '1'
				//递归下个空格
				if dfs(pos + 1) {
					return true
				}
				//如果之后的失败,回溯
				line[i][index] = false
				column[j][index] = false
				block[i/3][j/3][index] = false
			}
		}
		return false
	}
	//开始递归
	dfs(0)
}
