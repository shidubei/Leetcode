package stack

// func isVaild(s string) bool {
// 	if s == "" || len(s)/2 != 0 {
// 		return false
// 	}
// 	stack := make([]rune, 0)
// 	for _, v := range s {
// 		if (v == '(') || (v == '{') || (v == '[') {
// 			stack = append(stack, v)
// 		} else if ((v == ')' )&& len(stack) > 0 && stack[len(stack)-1] == '(') ||
// 			((v == ']' )&& len(stack) > 0&& stack[len(stack)-1] == '[') ||
// 			((v == '}' )&& len(stack) > 0&& stack[len(stack)-1] == '{') {
// 			stack = stack[:len(stack)-1]
// 		} else {
// 			return false
// 		}
// 	}
// 	return len(stack) == 0
// }
