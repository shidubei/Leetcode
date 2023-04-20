package array

/*题目：传入一个n*n的数组，要求在原数组的基础上直接返回它顺时针旋转90度的结果*/
//
//func rotate(matrix [][]int) {
//	if len(matrix) == 1 {
//		return
//	}
//	for i := 0; i < len(matrix); i++ {
//		for j := i + 1; j < len(matrix); j++ {
//			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
//		}
//	}
//	for i := 0; i < len(matrix); i++ {
//		for j := 0; j < len(matrix)/2; j++ {
//			matrix[i][j], matrix[i][len(matrix)-j-1] = matrix[i][len(matrix)-j-1], matrix[i][j]
//		}
//	}
//}
