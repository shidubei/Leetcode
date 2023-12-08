package Leetcode_String

/*题目：用一个函数来查找一个字符串数组的最长公共前缀*/

func longestCommonPrefix(strs []string) string {
	prefix := strs[0] // 以首字符作为标准
	for i := 0; i < len(strs); i++ {
		for j := 0; j < len(prefix); j++ {
			if len(strs[i]) <= j || strs[i][j] != prefix[j] { // 当当前字符长度小于匹配字符或不匹配的时候更新
				prefix = prefix[:j]
				break
			}
		}
	}
	return prefix
}
