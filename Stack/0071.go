package stack

// import "strings"

// func simplifyPath(path string) string {
// 	arr := strings.Split(path, "/")
// 	stack := make([]string, 0)
// 	for i := 0; i < len(arr); i++ {
// 		cur := arr[i]
// 		if cur == ".." {
// 			if len(stack) > 0 {
// 				stack = stack[:len(stack)-1]
// 			}
// 		} else if cur != "." && len(cur) > 0 {
// 			stack = append(stack, cur)
// 		}
// 	}
// 	if len(stack) == 0 {
// 		return "/"
// 	}
// 	res := strings.Join(stack, "/")
// 	return "/" + res

// }
