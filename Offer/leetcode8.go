package leetcode

/*剑指offer：数组中重复的数字*/
/*解题思路：
思路1：利用一个辅助数组记录每个数字与其出现的次数，最好查询大于1*/
// func findRepeatNumber(nums []int) int {
// 	rep := make([]int, len(nums))

// 	for i := 0; i < len(nums); i++ {
// 		rep[nums[i]]++
// 	}
// 	for i := 0; i < len(rep); i++ {
// 		if rep[i] > 1 {
// 			return i
// 		}
// 	}
// 	return 0
// }
