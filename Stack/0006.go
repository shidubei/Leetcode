package stack

/*按照顺序转为一个二维数组（有点像列输入，行输出）*/
// func convert(s string, numRows int) string {
// 	if len(s) <= numRows {
// 		return s
// 	}
// 	matrix := make([][]byte, numRows, numRows)
// 	down, up := 0, numRows-2
// 	for i := 0; i != len(s); {
// 		if down != numRows {
// 			matrix[down] = append(matrix[down], byte(s[i]))
// 			i++
// 			down++
// 		} else if up > 0 {
// 			matrix[up] = append(matrix[up], byte(s[i]))
// 			i++
// 			up--
// 		} else {
// 			down = 0
// 			up = numRows - 2
// 		}
// 	}
// 	res := make([]byte, 0, len(s))
// 	for _, row := range matrix {
// 		for _, item := range row {
// 			res = append(res, item)
// 		}
// 	}
// 	return string(res)
// }
/*思考，有没有规律*/
