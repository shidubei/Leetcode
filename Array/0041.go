package array

// 解法1：通过hash表来存储对比求解
//func firstMissingPositive(nums []int) int {
//	m := make(map[int]int, len(nums))
//	for _, v := range nums {
//		m[v] = v
//	}
//	for index := 1; index < len(nums)+1; index++ {
//		if _, ok := m[index]; !ok {
//			return index
//		}
//	}
//	return len(nums) + 1
//}

// 解法2：将所有的负数全都变为len(nums)+1,然后排序对比下标和数字，不相等的第一个正整数既是答案
// 思路错误，会出现重复
//func firstMissingPositive(nums []int) int {
//	for i := 0; i < len(nums); i++ {
//		if nums[i] < 0 {
//			nums[i] = len(nums) + 1
//		}
//	}
//	sort.Ints(nums)
//	if nums[0] == 0 {
//		nums = nums[1:]
//	}
//	for i := 0; i < len(nums); i++ {
//		if i+1 != nums[i] {
//			return i + 1
//		}
//	}
//	return len(nums) + 1
//}

// 解法3：1.将所有的负数都转化为最大值
//2.将所有数字绝对值对应的下标置为负数
//3.第一个不为负数的下标+1为所求值

//func firstMissingPositive(nums []int) int {
//	for i := 0; i < len(nums); i++ {
//		if nums[i] < 0 {
//			nums[i] = len(nums) + 1
//		}
//	}
//	for i := 0; i < len(nums); i++ {
//		if abs(nums[i]) <= len(nums) && nums[abs(nums[i])-1] > 0 {
//			nums[abs(nums[i])-1] *= -1
//		}
//	}
//	for i := 0; i < len(nums); i++ {
//		if nums[i] > 0 {
//			return i + 1
//		}
//	}
//	return len(nums) + 1
//}
//
//func abs(i int) int {
//	if i < 0 {
//		return -i
//	}
//	return i
//}
