package string

// 回溯的思想
/*题目：给定n对括号，要求输出所有匹配正确的可能形式，如n=3时
"((()))""(()())""()()()""(())()""()(())"*/

/*思路：
1.n型适配的题目，可能有某种规律，可能和递归相关，求解基本问题
2.将问题看作一个给定数组的插入问题，左括号的index肯定始终小于对应的右括号的index
3.n对括号分为左n个括号和右n个括号，左括号先插入，右括号只在已经存在左括号的时候插入，且位置在左括号之后
*/

//func generateParenthesis(n int) []string {
//	if n == 0 {
//		return []string{}
//	}
//	var res []string
//	findGenerateParenthesis(n, n, "", &res)
//	return res
//}
//
//func findGenerateParenthesis(lindex, rindex int, str string, res *[]string) {
//	if lindex == 0 && rindex == 0 {
//		// 左括号和右括号都插入
//		*res = append(*res, str)
//	}
//	// 插入左括号
//	if lindex > 0 {
//		findGenerateParenthesis(lindex-1, rindex, str+"(", res)
//	}
//	// 插入右括号，右括号要在左括号存在时插入
//	if rindex > 0 && lindex < rindex {
//		findGenerateParenthesis(lindex, rindex-1, str+")", res)
//	}
//}
