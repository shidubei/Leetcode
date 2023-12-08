package DynamicProgram

import (
	"bufio"
	"os"
	"strconv"

	//"fmt"
	"strings"
)

/*
题目：64.最小路径和
  给定一个包含非负整数的mxn网格grid，
  请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
  要求只能往右或者往下走
思路：经典的动态规划问题
  dp[i][j] = min(dp[i-1][j],dp[i][j-1])+s[i][j]
时间：2023/12/8
*/

const M, N = 201, 201

var Grid = make([][]int, M)

func minPathSum(m, n int) int {
	pre := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			} else if j != 0 && i == 0 {
				pre = Grid[i][j-1]
			} else if i != 0 && j == 0 {
				pre = Grid[i-1][j]
			} else {
				pre = min(Grid[i-1][j], Grid[i][j-1])
			}
			Grid[i][j] = pre + Grid[i][j]
		}
	}
	return Grid[m-1][n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// 多行输入
	input := bufio.NewScanner(os.Stdin)
	output := bufio.NewWriter(os.Stdout)

	for l := 0; l < len(Grid); l++ {
		Grid[l] = make([]int, N)
	}
	m, n := 0, 0
	for input.Scan() { // 每调用一次Scan,从输入流中读取一行
		// 获取这一行的输入
		str := strings.Split(input.Text(), " ")
		n = len(str)
		for j := 0; j < len(str); j++ {
			Grid[m][j], _ = strconv.Atoi(str[j])
		}
		m++
	}
	output.WriteString(strconv.Itoa(minPathSum(m, n)) + " ")
	// 输出答案
	output.Flush()
}
