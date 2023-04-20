package DynamicProgram

/*题目：和0062基本一致，但是在62的基础上路径中多出了障碍物*/
/*思路：一样的dp，在检测到障碍物的时候dp置为0*/
//func uniquePathsWithObstacles(obstacleGrid [][]int) int {
//	if len(obstacleGrid) == 0 || obstacleGrid[0][0] == 1 {
//		return 0
//	}
//	row := len(obstacleGrid)
//	col := len(obstacleGrid[0])
//	dp := make([][]int, row)
//	for i := 0; i < row; i++ {
//		dp[i] = make([]int, col)
//	}
//	dp[0][0] = 1
//	for i := 1; i < row; i++ {
//		if obstacleGrid[i][0] != 1 && dp[i-1][0] != 0 {
//			dp[i][0] = 1
//		}
//	}
//	for i := 1; i < col; i++ {
//		if obstacleGrid[0][i] != 1 && dp[0][i-1] != 0 {
//			dp[0][i] = 1
//		}
//	}
//	for i := 1; i < row; i++ {
//		for j := 1; j < col; j++ {
//			if obstacleGrid[i][j] != 1 {
//				dp[i][j] = dp[i-1][j] + dp[i][j-1]
//			}
//		}
//	}
//	return dp[row-1][col-1]
//}
