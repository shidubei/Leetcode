package array
/*思路：对撞指针*/
// func maxArea(height []int) int {
// 	max, start, end := 0, 0, len(height)-1

// 	for start < end {
// 		width := end - start
// 		high := 0
// 		if height[start] < height[end] {
// 			high = height[start]
// 			start++
// 		} else {
// 			high = height[end]
// 			end--
// 		}
// 		s := high * width
// 		if s > max {
// 			max = s
// 		}
// 	}
// 	return max
// }
