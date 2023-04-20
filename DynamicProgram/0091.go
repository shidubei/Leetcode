package DynamicProgram

/*题目：每个字母对应一个数字，给一个字符串，算出它的加密方式一共有多少种*/

/*
思路：经典的dp问题，dp[0]=1.dp[1]在以0为开头的时候等于0，不为0的时候等于1
如果s[i-1:i]为1-9的字母，那么dp[i]+=dp[i-1]，s[i-2:i]为10-26的字母，dp[i]+=dp[i-2]
*/
//func numDecodings(s string) int {
//	dp := make([]int, len(s)+1)
//	dp[0] = 1
//	for i := 1; i <= len(s); i++ {
//		if s[i-1] != '0' {
//			dp[i] += dp[i-1]
//		}
//		if i > 1 && s[i-2] != '0' && (s[i-2]-'0')*10+(s[i-1]-'0') <= 26 {
//			dp[i] += dp[i-2]
//		}
//	}
//	return dp[len(s)]
//}
