package leetcode

// type dir struct {
// 	x, y int
// }

// // 方位上下左右
// var dirs = []dir{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

// func exist(board [][]byte, word string) bool {
// 	if len(board)*len(board[0]) < len(word) || len(board) == 0 {
// 		return false
// 	}
// 	m, n := len(board), len(board[0])
// 	//利用一个伴随的矩阵记录刺探信息，匹配相应位置则置为真
// 	vis := make([][]bool, len(board))
// 	for i := range vis {
// 		vis[i] = make([]bool, len(board[0]))
// 	}
// 	//递归
// 	var check func(i, j, k int) bool
// 	check = func(i, j, k int) bool {
// 		if board[i][j] != word[k] {
// 			return false
// 		} //不匹配返回
// 		if k == len(word)-1 {
// 			return true
// 		} //匹配完成，返回真
// 		vis[i][j] = true //当前匹配置为真

// 		defer func() { vis[i][j] = false }() //结束时回溯为假
// 		//尝试递归刺探
// 		for _, d := range dirs {
// 			//若上下左右有一个方位为假进行刺探
// 			if newI, newJ := i+d.x, j+d.y; 0 <= newI && newI < m && 0 <= newJ && newJ < n && !vis[newI][newJ] {
// 				if check(newI, newJ, k+1) {
// 					//刺探成功匹配返回真
// 					return true
// 				}
// 			}
// 		}
// 		//递归刺探失败返回假，调用defer
// 		return false
// 	}
// 	for i, raw := range board {
// 		for j := range raw {
// 			if check(i, j, 0) {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }
