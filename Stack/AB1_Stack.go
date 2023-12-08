package stack

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func AB1Stack() {
	// 多行输入
	input := bufio.NewScanner(os.Stdin)
	output := bufio.NewWriter(os.Stdout)
	// 测试用例过长的时候，要先指定bufio的bufer，否则会读取失败
	// bufer := make([]byte,2000*1024)zha
	// input.Buffer(bufer,len(bufer))
	var stack []string
	for input.Scan() { // 每调用一次Scan,从输入流中读取一行
		str := strings.Split(input.Text(), " ")
		switch str[0] {
		case "push":
			stack = append(stack, str[1])
		case "pop":
			if len(stack) == 0 {
				output.WriteString(fmt.Sprintln("error"))
			} else {
				output.WriteString(fmt.Sprintln(stack[:len(stack)-1]))
				stack = stack[:len(stack)-1]
			}
		case "top":
			if len(stack) == 0 {
				output.WriteString(fmt.Sprintln("error"))
			} else {
				output.WriteString(fmt.Sprintln(stack[:len(stack)-1]))
			}
		default:
			continue
		}
	}
	output.Flush()
}
