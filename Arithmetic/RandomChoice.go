package Arithmetic

import (
	"bufio"
	"math/rand"
	"os"
	"strconv"

	//"fmt"
	"strings"
)

/*
题目：随机选择算法之返回无序数组中第K大或第K小的数字
思路：
  1.改良随机快排算法，选择一个随机数进行快排，如果这个数的边界值刚好包含K，则返回这个数
  2.否则继续在左边或右边执行
时间：2023/12/10
链接：
*/

var first int
var last int

var arr []int

func findKthLargest(k int) int {
	return randomChoice(len(arr) - k)
}
func randomChoice(i int) int {
	ans := 0
	for l, r := 0, len(arr)-1; l <= r; {
		partition2(arr, l, r, arr[l+rand.Intn(r-l+1)])
		//         i
		// ...first....last...
		if i < first {
			// 边界值在所需值右侧，更新右边界
			r = first + 1
		} else if i > last {
			//边界值在所需值左侧，更新左边界
			l = last + 1
		} else {
			ans = arr[i]
			break
		}
	}
	return ans
}

func partition2(arr []int, left, right, x int) {
	first, last = left, right
	i := left
	for i <= last {
		if arr[i] == x {
			i++
		} else if arr[i] < x {
			arr[first], arr[i] = arr[i], arr[first]
			first++
			i++
		} else {
			arr[i], arr[last] = arr[last], arr[i]
			last--
		}
	}
}

func AcmMode() {
	// 多行输入
	input := bufio.NewScanner(os.Stdin)
	output := bufio.NewWriter(os.Stdout)
	// 测试用例过长的时候，要先指定bufio的bufer，否则会读取失败
	// bufer := make([]byte,2000*1024)
	// input.Buffer(bufer,len(bufer))
	i := 0
	ans := 0
	for input.Scan() { // 每调用一次Scan,从输入流中读取一行
		// 获取这一行的输入
		str := strings.Split(input.Text(), " ")
		if i == 0 {
			arr = make([]int, len(str))
			for i := 0; i < len(str); i++ {
				arr[i], _ = strconv.Atoi(str[i])
			}
		} else {
			k, _ := strconv.Atoi(str[0])
			ans = findKthLargest(k)
		}

		//TODO
		// 将结果写入缓冲输出区
		output.WriteString(strconv.Itoa(ans))
		i++
		//TODO
	}
	// 输出答案
	output.Flush()
}
