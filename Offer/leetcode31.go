package leetcode

/*双指针；条件可以再优化*/
// func exchange(nums []int) []int {
// 	if nums == nil {
// 		return nil
// 	}
// 	i, j := 0, len(nums)-1

// 	for i < j {
// 		if nums[i]%2 == 0 && nums[j]%2 == 1 {
// 			nums[i], nums[j] = nums[j], nums[i]
// 			i++
// 			j--
// 		} else if nums[i]%2 == 1 && nums[j]%2 == 1 {
// 			i++
// 		} else if nums[i]%2 == 0 && nums[j]%2 == 0 {
// 			j--
// 		} else if nums[i]%2 == 1 && nums[j]%2 == 0 {
// 			i++
// 			j--
// 		}
// 	}
// 	return nums
// }
