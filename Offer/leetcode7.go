package leetcode

/*剑指Offer7：左旋转字符串
字符串的左旋转操作是把字符串前面的若干个字符转移到字符串的尾部。请定义一个函数实现字符串左旋转操作的功能。比如，输入字符串"abcdefg"和数字2，该函数将返回左旋转两位得到的结果"cdefgab"。
解题思路：既然只是左旋，左旋部分的顺序没有改变
那么直接将原字符串切片再合并就好了*/
// func reverseLeftWords(s string, n int) string {
// 	return s[n:] + s[:n]
// }
