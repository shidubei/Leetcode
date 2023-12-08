package main

import "fmt"

func main() {
	s := "dahsfghjkkjhgfqwrffagoagnkviadfdangvjandlkfgnadvfnbfagioasfnfa"
	fmt.Println(longestPalindrome3(s))
}

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
