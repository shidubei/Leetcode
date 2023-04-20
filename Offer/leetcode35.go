package leetcode

// func movingCount(m int, n int, k int) int {
// 	res := 0
// 	visit := make([][]bool, m)
// 	for i, _ := range visit {
// 		visit[i] = make([]bool, n)
// 	}
// 	visit[0][0] = true
// 	for i := 0; i < m; i++ {
// 		for j := 0; j < n; j++ {
// 			if i == 0 && j == 0 || digitSum(i, j) > k {
// 				continue
// 			}
// 			if i >= 1 {
// 				visit[i][j] = visit[i][j] || visit[i-1][j]
// 			}
// 			if j >= 1 {
// 				visit[i][j] = visit[i][j] || visit[i][j-1]
// 			}
// 			if visit[i][j] {
// 				res++
// 			}
// 		}
// 	}
// 	return res + 1
// }
// func digitSum(i, j int) int {
// 	res := 0
// 	for i != 0 {
// 		res += i % 10
// 		i /= 10
// 	}
// 	for j != 0 {
// 		res += j % 10
// 		j /= 10
// 	}
// 	return res
// }
