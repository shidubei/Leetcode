package leetcode

/*解题思路：二维数组的查找
解题思路：
思路1：因为从上到下和从左到右都是非递减的顺序，所以最右上的数字
很特殊，其左边比其小，下边比其大，因此用来判断
如果目标小于它，就往左边走，大于它，往下边走
直到左边或下边无法再走*/
// func findNumberIn2DArray(matrix [][]int, target int) bool {
// 	for i := 0; i < len(matrix); i++ {
// 		for j := len(matrix[0]) - 1; j >= 0; j-- {
// 			if target == matrix[i][j] {
// 				return true
// 			} else if target > matrix[i][j] {
// 				break
// 			} else if target < matrix[i][j] {
// 				continue
// 			}
// 		}
// 	}
// 	return false
// }
/*思路2:二维二分查找？*/
/*考虑用例的特殊情况[],[[]],[[1,1]],[[1],[1]]等边界用例*/
// func findNumberIn2DArray(matrix [][]int, target int) bool {
// 	for i := 0; i < len(matrix); i++ {
// 		if findNumberInRow(matrix[i], target) {
// 			return true
// 		}
// 	}
// 	return false
// }
// func findNumberInRow(num []int, target int) bool {
// 	start, end := 0, len(num)-1
// 	mid := (start + end) / 2

// 	for start < end {
// 		if target == num[mid] {
// 			return true
// 		}
// 		if target > num[mid] {
// 			start = mid
// 			mid = (start + end) / 2
// 			continue
// 		}
// 		if target < num[mid] {
// 			end = mid
// 			mid = (start + end) / 2
// 			continue
// 		}
// 	}
// 	return false
// }
