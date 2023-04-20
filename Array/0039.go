package array

//
//import "sort"
//
///*题目描述：给一个无重复数的数组和一个目标数，找到一个和为目标数的序列，数字可重复利用*/
//
///*
//思路：
//1.递归迭代，即每次传入的target数是上一次的target-nums[i],即在原数组中一直去遍历所有去寻找相加为diff的数字
//*/
//func combinationSum(candidates []int, target int) [][]int {
//	if len(candidates) == 0 {
//		return [][]int{}
//	}
//	var res [][]int
//	var ans []int
//	sort.Ints(candidates) //  先进行一轮排序
//	findCombinationSum(candidates, target, 0, ans, &res)
//	return res
//}
//
//func findCombinationSum(nums []int, target, index int, ans []int, res *[][]int) {
//	if target <= 0 { //回溯条件，target为0或小于0了
//		if target == 0 {
//			ansF := make([]int, len(ans))
//			copy(ansF, ans)
//			*res = append(*res, ansF)
//		}
//		return
//	}
//	// 迭代,利用递归的思想，每次去寻找新一轮的目标数（原目标数减去已匹配的数字）
//	for i := index; i < len(nums); i++ {
//		if nums[i] > target {
//			break
//		} // 当前的数大于target了则退出，否则迭代
//		ans = append(ans, nums[i])
//		findCombinationSum(nums, target-nums[i], i, ans, res) // 迭代，传入新的目标数
//		ans = ans[:len(ans)-1]                                // 退回说明不行，则推出
//	}
//}
