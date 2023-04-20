package stack

/*题目：给出一个数组，数组表示x轴从0开始的每宽度为单位长度的柱子高度，求这个数组能接多少雨水*/

/*思路：
1. 首先明确雨水的计算标准，可以将雨水看作最高高度（随时更新）-当前高度的值，即max-height[i]，仅在当前小于最大值的情况下，否则更新
2. 利用双指针，左右同时进行，左边小于右边时左移，否则右移，知道二者相遇，减少循环*/

//func trap(height []int) int {
//	res, l, r, maxL, maxR := 0, 0, len(height)-1, 0, 0
//	// 双指针遍历
//	for l <= r {
//		// 移动条件
//		if height[l] <= height[r] {
//			if height[l] > maxL {
//				maxL = height[l]
//			} else {
//				// 增加雨水体积
//				res += maxL - height[l]
//			}
//			l++
//		} else {
//			if height[r] >= maxR {
//				maxR = height[r]
//			} else {
//				// 增加雨水体积
//				res += maxR - height[r]
//			}
//			r--
//		}
//	}
//	return res
//}
