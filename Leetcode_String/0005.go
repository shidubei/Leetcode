package Leetcode_String

import (
	"bufio"
	"os"
	//"fmt"
	"strings"
)

/*
题目：给你一个字符串 s，找到 s 中最长的回文子串。
思路：
  1.回文字符串分为过中心和不过中心两种和都存在三种情况
  2.首先确定字符串的中心，向两边扩散，若两边不相同，则说明回文字符串在两边，步骤3
  3.将字符串分为两部分，然后继续查找字符串(递归)
  4.若字符串在中心，继续扩散，直到不相等，此时记录长度，将剩余的两部分字符串继续查找，直到字符串边界
*/
// 解法1：动态规划
func longestPalindrome1(s string) string {
	ans, dp := "", make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}
	// 动态规划的公式为dp[i][j]=dp[i+1][j-1]&&s[i]==s[j]
	// 又因为dp[i+1][j-1]需要i和j之间相差至少3个，所以最终的公式为
	// dp[i][j]=(s[i] == s[j]) && ((j-i < 3) || dp[i+1][j-1]) 注意的是 (j-i<3)要在||左边来跳过后面的dp[i+1][j-1],否则会报错
	for i := len(s) - 1; i >= 0; i-- {
		// i要依次往前，j从i位置依次往后
		//fmt.Println()
		for j := i; j < len(s); j++ {
			dp[i][j] = (s[i] == s[j]) && ((j-i < 3) || dp[i+1][j-1])
			//fmt.Print(s[i:j+1] + " ")
			//fmt.Println(dp[i][j], i, j)
			// 又因为到最后每次单个的时候一定为true，所以，要么此时ans为空，要么此时的遍历长度大于当前的ans，才更新ans
			if dp[i][j] && (ans == "" || j-i+1 > len(ans)) {
				ans = s[i : j+1]
			}
		}
	}
	return ans
}

// 解法2：中心扩散
// 相对于动态规划，将目标选取为从中心向两边扩散
// eg:  ..a..b..c..d..c..b..f..
// ..为虚中心(起始为两边的字符)，字符为实中心(起始为实际字符)
func longestPalindrome2(s string) string {
	ans := ""
	for i := 0; i < len(s); i++ {
		// 长度为奇数的时候，中心实际且一致，扩展的起点都是i
		ans = maxExpandCenter(s, i, i, ans)
		// 长度为偶数的时候，中心虚拟，扩展的起点为i和i+1
		ans = maxExpandCenter(s, i, i+1, ans)
	}
	return ans
}

func maxExpandCenter(s string, i, j int, ans string) string {
	res := ""
	for i >= 0 && j < len(s) && s[i] == s[j] { //扩展条件
		res = s[i : j+1]
		i--
		j++
	}
	// 返回目前的最大子串
	if len(res) > len(ans) {
		return ans
	}
	return ans
}

// 马拉车算法 manacher算法
// 将字符串进行预处理，得到长度始终为奇数的字符串(2n+1)
// 马拉车算法相对于中心扩散算法的优化在于，中心扩散算法会遍历每个扩散中心，而马拉车算法重点在于计算出下一个扩散中心
// 1.需要遍历的为i,考虑最大回文距离maxRight,扩散中心center，i关于center的对称点mirror，记录当前最大扩散半径的数组p
// 2.i>=maxRight时，maxRight更新，center更新
// 3.i<maxRight时，分类讨论.......mirror....center....i....maxRight....
//
//	3.1 以i为扩散中心时，到不了当前center的maxRight，由于对称性，此时的p[i]=p[mirror]
//	3.2 以i为扩散中心时，刚好到达maxRight,此时的p[i]=maxRight-i
//	3.3 以i为扩散中心时，大于maxRight,此时的p[i]>maxRight-i
//
// 因此得到p[i]的计算公式为min(maxRight-i,p[mirror]),依次更新即可
func longestPalindrome3(s string) string {
	ss := preManacher(s)
	p := make([]int, len(ss))
	maxRight, center, begin, maxLen := 0, 0, 0, 1
	for i := 0; i < len(ss); i++ {
		// 获取初始的p[i]
		if i < maxRight {
			mirror := center<<1 - i
			p[i] = min(maxRight-i, p[mirror])
		}
		// 此时在p[i]的基础上进行中心扩散，若能继续扩散，则更新p[i]
		left, right := i-(1+p[i]), i+(1+p[i])
		for left >= 0 && right < len(ss) && ss[left] == ss[right] {
			p[i]++
			left--
			right++
		}
		//	更新下一个中心
		if i+p[i] > maxRight {
			maxRight = i + p[i]
			center = i
		}
		//	维持最大回文子串
		if p[i] > maxLen {
			maxLen = p[i]
			begin = (i - maxLen) >> 1
		}
	}
	return s[begin : begin+maxLen]
}

func preManacher(s string) string {
	res := ""
	for i := 0; i < 2*len(s)+1; i++ {
		if i%2 == 0 {
			res += "#"
		} else {
			res += string(s[(i-1)/2])
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func AcmMode() {
	// 多行输入
	input := bufio.NewScanner(os.Stdin)
	output := bufio.NewWriter(os.Stdout)
	// 测试用例过长的时候，要先指定bufio的bufer，否则会读取失败
	// bufer := make([]byte,2000*1024)
	// input.Buffer(bufer,len(bufer))
	for input.Scan() { // 每调用一次Scan,从输入流中读取一行
		// 获取这一行的输入
		str := strings.Split(input.Text(), " ")
		//TODO
		// 将结果写入缓冲输出区
		output.WriteString(str[0])
		//TODO
	}
	// 输出答案
	output.Flush()
}
