package array

/*题目：给定一个数组，要求返回这个数组的全排列*/
/*思路:
DFS:深度优先搜索*/
func permute(nums []int) [][]int {
	var res [][]int

	if len(nums) == 0 {
		return [][]int{}
	}
	ifUse, t := make([]bool, len(nums)), []int{}
	listPermute(nums, 0, t, &res, &ifUse)
	return res
}

// DFS方法，传入原数组，当前遍历的下标，暂时结果t,最终结果res和是否使用ifUse
func listPermute(nums []int, index int, t []int, res *[][]int, ifUse *[]bool) {
	// 返回条件，遍历完，index==len(nums)
	if index == len(nums) {
		temp := make([]int, len(t))
		copy(temp, t)
		*res = append(*res, temp) // 将当前结果加入答案
		return
	}
	for i := 0; i < len(nums); i++ {
		// 先检查当前数字是否使用
		if !(*ifUse)[i] {
			(*ifUse)[i] = true                        // 置为使用
			t = append(t, nums[i])                    // 加入暂存
			listPermute(nums, index+1, t, res, ifUse) // DFS
			// 回溯
			t = t[:len(t)-1]
			(*ifUse)[i] = false
		}
	}
	return

}
