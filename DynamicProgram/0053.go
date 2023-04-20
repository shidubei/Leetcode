package DynamicProgram

/*题目：给一个int数组，要求返回其和值最大的子数组的和值*/

/*思路：该题简单在只返回值，不返回数组，返回数组的话会麻烦一点
1. 从下标0开始扩容，遇到负数不扩容，遇到正数更新
*/

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	res, ap := nums[0], make([]int, len(nums))
	ap[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if ap[i-1] > 0 {
			// 是正数
			ap[i] = ap[i-1] + nums[i]
		} else {
			ap[i] = ap[i-1]
		}
		res = max(res, ap[i])
	}
	return 0
}
func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
