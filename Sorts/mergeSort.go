package Sort

import (
	"bufio"
	"os"
	"strconv"

	//"fmt"
	"strings"
)

/*
题目：归并排序-递归和非递归解法
思路：左部分排好序，右部分排好序，再归并
  eg: [6 4 2 3 9 4]
  左侧 [6] merge [4] -> [4 6] merge [2] -> [2 4 6]
  右侧 [3] merge [9] -> [3 9] merge [4] -> [3 4 9]
  接着merge [2 3 4 4 6 9]
时间：2023/12/8
*/

// 设计静态变量

const Max = 501

var help [Max]string

var arr []string

func mergeSortRecurrent(left, right int) {
	if left == right {
		return
	}
	mid := (left + right) >> 1
	mergeSortRecurrent(left, mid)
	mergeSortRecurrent(mid+1, right)
	merge(left, mid, right)
}

func merge(left, mid, right int) {
	//	要在left~right范围内有序
	index, lindex, rindex := left, left, mid+1
	for lindex <= mid && rindex <= right {
		if arr[lindex] <= arr[rindex] {
			help[index] = arr[lindex]
			index++
			lindex++
		} else {
			help[index] = arr[rindex]
			index++
			rindex++
		}
	}

	for lindex <= mid {
		help[index] = arr[lindex]
		index++
		lindex++
	}
	for rindex <= right {
		help[index] = arr[rindex]
		index++
		rindex++
	}
	for i := left; i <= right; i++ {
		arr[i] = help[i]
	}
}
func AcmMode() {
	// 多行输入
	input := bufio.NewScanner(os.Stdin)
	output := bufio.NewWriter(os.Stdout)
	// 测试用例过长的时候，要先指定bufio的bufer，否则会读取失败
	// bufer := make([]byte,2000*1024)
	// input.Buffer(bufer,len(bufer))
	cnt := 0
	for input.Scan() { // 每调用一次Scan,从输入流中读取一行
		cnt++
		// 获取这一行的输入
		str := strings.Split(input.Text(), " ")
		if cnt == 1 {
			n, _ := strconv.Atoi(str[0])
			arr = make([]string, n)
		} else {
			for i := 0; i < len(arr); i++ {
				arr[i] = str[i]
			}
			mergeSortRecurrent(0, len(arr)-1)
		}
		// 将结果写入缓冲输出区
		for i := 0; i < len(arr); i++ {
			output.WriteString(arr[i] + " ")
		}
		//TODO
	}
	// 输出答案
	output.Flush()
}
