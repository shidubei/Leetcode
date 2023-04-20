package leetcode

/*剑指offer6：替换字符串的空格字符
即将字符串中的空格全部替换为%20
解题思路：考察的是对字符串的操作，第一个想法是对字符串用正则表达式匹配来替换，但失败了
然后是常规的想法，将字符串转化为字符串数组，用一个字节数组copy该字符串，遇到空格直接替换就是了*/

// func replaceSpace(s string) string {
// 	res := []byte(s)
// 	re := make([]byte, 0)

// 	for _, v := range res {
// 		if v == 32 {
// 			re = append(re, '%')
// 			re = append(re, '2')
// 			re = append(re, '0')
// 		}
// 		re = append(re, v)
// 	}
// 	return string(re)
// }
