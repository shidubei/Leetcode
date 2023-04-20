package array

// func nextPermutation(nums []int) {
// 	i, j := 0, 0
// 	//从后向前，记录第一个没有降序排列的位置
// 	for i = len(nums) - 2; i >= 0; i-- {
// 		if nums[i] < nums[i+1] {
// 			break
// 		}
// 	}
// 	if i >= 0 {
// 		for j = len(nums) - 1; j > i; j-- {
// 			if nums[j] > nums[i] {
// 				break
// 			}
// 		}
// 		swap(&nums, i, j)
// 	}
// 	reverse(&nums, i+1, len(nums)-1)
// }

// func reverse(nums *[]int, n, m int) {
// 	for n < m {
// 		swap(nums, n, m)
// 		n++
// 		m--
// 	}
// }
// func swap(nums *[]int, i, j int) {
// 	(*nums)[i], (*nums)[j] = (*nums)[j], (*nums)[i]
// }
