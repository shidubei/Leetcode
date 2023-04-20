package leetcode

// import "sort"

/*扑克牌中的顺子*/
/*常规思路，无排序，只考虑成功情况*/
// func isStraight(nums []int) bool {
// 	maxDiff := 4
// 	min := 13
// 	max := 1
// 	m := make([]int, 14)
// 	for i := range nums {
// 		m[nums[i]]++
// 	}
// 	for i := range m {
//     if i != 0 && m[i] >= 2 {
// 			return false
// 		}
// 	}
// 	for i := range nums {
// 		if nums[i] >= max {
// 			max = nums[i]
// 		}
// 		if nums[i] <= min && nums[i] != 0 {
// 			min = nums[i]
// 		}
// 		if max-min <= maxDiff {
// 			return true
// 		}
// 	}
// 	return false
// }
/*借助排序*/
// func isStraight(nums []int) bool {
// 	sort.Ints(nums)
// 	joker := 0
// 	for i, v := range nums {
// 		if v == 0 {
// 			joker++
// 			continue
// 		}
// 		if i == 4 {
// 			break
// 		}
// 		if nums[i] == nums[i+1] {
// 			return false
// 		}
// 	}
// 	return nums[4]-nums[joker] < 5
// }
