package leetcode

/*hashmap*/
// func twoSum(nums []int, target int) []int {
// 	if nums == nil {
// 		return nil
// 	}
// 	m := make(map[int]int)
// 	for i, _ := range nums {
// 		diff := target - nums[i]
// 		if j, ok := m[diff]; ok {
// 			return []int{nums[i], nums[j]}
// 		} else {
// 			m[nums[i]] = i
// 		}
// 	}
// 	return []int{}
// }
/*双指针*/
// func twoSum(nums []int, target int) []int {
// 	if nums == nil {
// 		return nil
// 	}
// 	i, j := 0, len(nums)-1
// 	for nums[i]+nums[j]!=target {
// 		if nums[i]+nums[j] < target {
// 			i++
// 		} else {
// 			j--
// 		}
// 	}
// 	return []int{nums[i],nums[j]}
// }
