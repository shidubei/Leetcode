package Leetcode_String

/*题目：给一个只包含'('和')'的字符串，要求找出最长有效的括号子串的长度*/

/*解法1：辅助栈法-记录隔断
1.此题是在匹配括号的基础上的加强版，与单纯的匹配不同，确定长度在此基础上是记录隔断的位置，
隔断即未匹配成功的右括号*/

func longestValidParentheses(s string) int {
	stack, res := []int{}, 0
	stack = append(stack, -1)
	// 开始遍历
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			// 左括号入栈
			stack = append(stack, i)
		} else { //右括号出栈
			stack = stack[:len(stack)-1]
			// 检测之前的是否匹配完
			if len(stack) == 0 {
				// 隔断入栈
				stack = append(stack, i)
			} else {
				// 之前存在隔断，更新最大量
				res = max(res, i-stack[len(stack)-1])
			}
		}
	}
	return res
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
