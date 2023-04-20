package leetcode

// func maxValue(grid [][]int) int {
// 	m, n := len(grid), len(grid[0])
// 	for i := 0; i < m; i++ {
// 		for j := 0; j < n; j++ {
// 			if i == 0 && j == 0 {
// 				continue
// 			} else if i == 0 && j != 0 {
// 				grid[i][j] += grid[i][j-1]
// 			} else if i != 0 && j == 0 {
// 				grid[i][j] += grid[i-1][j]
// 			} else if i != 0 && j != 0 {
// 				grid[i][j] += max(grid[i][j-1], grid[i-1][j])
// 			}
// 		}
// 	}
// 	return grid[m-1][n-1]
// }
// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	} else {
// 		return b
// 	}
// }
