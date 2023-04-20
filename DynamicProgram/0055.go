package DynamicProgram

/*题目：给定一个数组，每个数组上的数字表示从该位置能跳的最远距离，要求判断能否跳到数组的最远位置*/

/*思路：
一道简单的规划题，首先需要维持一个maxReach即每个位置能到达的最大长度数，在遍历的过程中更新
如果当前位置大于之前的maxReach,说明该位置不可达，返回false*/

//func canJump(nums []int) bool {
//	if len(nums) == 0 {
//		return false
//	}
//	if len(nums) == 1 {
//		return true
//	}
//	maxReach := 0
//	for i, x := range nums {
//		// 判断当前位置是否可达
//		if i > maxReach {
//			return false
//		}
//		if i+x > maxReach {
//			maxReach = i + x
//		}
//	}
//	return true
//}
//
//func canJump2(nums []int) bool {
//	if len(nums) == 1 {
//		return true
//	}
//	a := 1
//	for i := len(nums) - 2; i >= 0; i-- {
//		if nums[i] < a {
//			a++
//		} else {
//			a = 1
//		}
//	}
//	return nums[0] >= a
//}
