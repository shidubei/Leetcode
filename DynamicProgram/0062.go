package DynamicProgram

/*题目：给一个m*n的数组，记左上角为start,右下角为end,记录从左上角一直到左下角一共有多少种走法*/

/*思路：经典动态规划题，右下角只有左边和上边可达，所以dp[end]=dp[end_left]+dp[end_below]*/
//func uniquePaths(m int, n int) int {
//	dp := make([][]int, m)
//	for i := 0; i < m; i++ {
//		dp[i] = make([]int, n)
//	}
//	dp[0][0] = 1
//	i, j := 0, 0
//	for i = 0; i < m; i++ {
//		for j = 0; j < n; j++ {
//			if i == 0 || j == 0 {
//				dp[i][j] = 1
//			} else {
//				dp[i][j] = dp[i-1][j] + dp[i][j-1]
//			}
//		}
//	}
//	return dp[i][j]
//}

// 法2：数学公式法
