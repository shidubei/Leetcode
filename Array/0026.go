package array

// func removeDuplicates(nums []int) int {
// 	if len(nums) == 0 {
// 		return 0
// 	}
// 	last, find := 0, 0
// 	for last < len(nums)-1 {
// 		for nums[last] == nums[find] {
// 			find++
// 			if find == len(nums) {
// 				return last + 1
// 			}
// 		}
// 		nums[last+1] = nums[find]
// 		last++
// 	}
// 	return last + 1
// }
