package leetcode

/*第一次只出现一次的字符*/
// func firstUniqChar(s string) byte {
// 	if s == "" {
// 		return ' '
// 	}
// 	var charArray [26]int

// 	for i := 0; i < len(s); i++ {
// 		charArray[s[i]-'a']++
// 	}
// 	for i := 0; i < len(s); i++ {
// 		if charArray[s[i]-'a'] == 1 {
// 			return s[i]
// 		}
// 	}
// 	return ' '
// }
