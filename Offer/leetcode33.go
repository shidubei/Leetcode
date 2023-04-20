package leetcode

// import "strings"

/*利用栈来解决*/
// func reverseWords(s string) string {
// 	if s == "" {
// 		return ""
// 	} //空串
// 	var res string
// 	stuck := make([]string, 0)
// 	str := strings.Split(s, " ")
// 	for i, _ := range str {
// 		if str[i] != "" {
// 			stuck = append(stuck, str[i])
// 		}
// 	}
// 	//如果全部为空字符
// 	if len(stuck) == 0 {
// 		return ""
// 	}
// 	for i := len(stuck) - 1; i >= 0; i-- {
// 		res += stuck[i] + " "
// 	}
// 	return res[:len(res)-1]
// }
