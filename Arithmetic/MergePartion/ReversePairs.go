package MergePartion

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	//"fmt"
	"strings"
)

/*
题目：翻转对
  给定一个数组nums,如果i<j且nums[i]>2*nums[j],则将(i,j)称为一个重要翻转对
  返回数组中重要翻转对的数量
思路：
 1.暴力算法
   2层循环直接遍历
 2.归并分治
   分为左部分和右部分以及归并部分
   总翻转对个数=左翻转对个数+右翻转对个数+归并时翻转对个数
   且左右各自有序时，对归并时的处理有帮助
时间：2023/12/10
链接：https://leetcode.cn/problems/reverse-pairs/
*/

const Max2 = 50001

var arr1 = make([]int, Max2)

var help1 = make([]int, Max2)

func reversePairs(left, right int) int {
	if left == right {
		return 0
	}
	mid := (left + right) / 2
	return reversePairs(left, mid) + reversePairs(mid+1, right) + merge2(left, mid, right)
}

func merge2(left, mid, right int) int {
	ans := 0
	for i, j, sum := left, mid+1, 0; j <= right; j++ {
		for i <= mid && arr1[i] > 2*arr1[j] {
			sum++
			i++
		}
		ans += sum
	}
	// 归并时按从大到小排序
	index, lindex, rindex := left, left, mid+1
	for lindex <= mid && rindex <= right {
		if arr1[lindex] >= arr1[rindex] {
			help1[index] = arr1[lindex]
			index++
			lindex++
		} else {
			help1[index] = arr1[rindex]
			index++
			rindex++
		}
	}
	for lindex <= mid {
		help1[index] = arr1[lindex]
		index++
		lindex++
	}
	for rindex <= right {
		help1[index] = arr1[rindex]
		index++
		rindex++
	}
	for i := left; i <= right; i++ {
		arr1[i] = help1[i]
	}
	return ans
}

func main() {
	// 多行输入
	input := bufio.NewScanner(os.Stdin)
	output := bufio.NewWriter(os.Stdout)
	// 测试用例过长的时候，要先指定bufio的bufer，否则会读取失败
	// bufer := make([]byte,2000*1024)
	// input.Buffer(bufer,len(bufer))
	for input.Scan() { // 每调用一次Scan,从输入流中读取一行
		// 获取这一行的输入
		str := strings.Split(input.Text(), " ")
		for i := 0; i < len(str); i++ {
			arr1[i], _ = strconv.Atoi(str[i])
		}
		fmt.Println(arr1[:len(str)])
		//TODO
		// 将结果写入缓冲输出区
		output.WriteString(strconv.Itoa(reversePairs(0, len(str)-1)))
		//TODO
	}
	// 输出答案
	output.Flush()
}
