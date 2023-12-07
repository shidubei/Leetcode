package stack

import (
	"bufio"
	"os"
	"fmt"
	"strings"
)

func stack() {
	// 多行输入
	input := bufio.NewScanner(os.Stdin)
	// 测试用例过长的时候，要先指定bufio的bufer，否则会读取失败
	// bufer := make([]byte,2000*1024)
	// input.Buffer(bufer,len(bufer))
	var str []string
	for input.Scan() { // 每调用一次Scan,从输入流中读取一行
		str = strings.Split(input.Text(), " ") // 获取这一行的输入
		//TODO
	}
	fmt.Println(str)
}
