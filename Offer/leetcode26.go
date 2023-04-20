package leetcode

// func lengthOfLongestSubstring(s string) int {
// 	m := make(map[byte]int)
// 	res := make([]byte, 0)
// 	var Max int
// 	for i:=0;i<len(s);i++{
// 		res = append(res, s[i])
// 		if _,ok := m[s[i]];!ok {
// 			m[s[i]]++
// 			Max=max(Max,len(res))
// 		} else {
// 			res=res[1:]
// 			i--
// 		}
// 	}
// 	return Max
// }



// func lengthOfLongestSubstring(s string) int{
// 	 m:=map[byte]int{}
// 	rk,ans:=-1,0
// 	n:=len(s)

// 	for i:=0;i<n;i++{
// 		if i!=0{
// 			delete(m,s[i-1])
// 		}
// 		for rk+1<n&&m[s[rk+1]]==0{
// 			m[s[rk+1]]++
// 			rk++
// 		}
// 		ans=max(ans,rk-i+1)
// 	}
// 	return ans
// }
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
