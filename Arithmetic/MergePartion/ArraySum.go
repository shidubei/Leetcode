package MergePartion

import (
	"bufio"
	"os"
	"strconv"

	//"fmt"
	"strings"
)

/*
题目：计算数组的小和-(归并分治的思想)
  给一个数组arr[1,3,5,2,4,6]
  小和即第i个数的左边所有小于其的和，数组的小和即数组所有小和之和
  如arr[0] = 0
    arr[1] = 1
    arr[2] = 1+3 = 4
    arr[3] = 1
    arr[4] = 1+3+2 = 6
    arr[5] = 1+2+3+4+5 = 15
  然后所有相加为27
思路：
1. 暴力破解
   2层for循环解决，时间复杂度为O(n)
时间复杂度：O(N^2)
空间复杂度：
2. 归并分治
  问题可以拆解为左问题，右问题和归并问题
  那么可以利用归并排序的思想来解题
  即小数之和等于左小数之和+右小数之和+归并时小数之和
时间：2023/12/9
时间复杂度：O(NlogN)
空间复杂度：
链接：https://www.nowcoder.com/practice/edfe05a1d45c4ea89101d936cac32469
*/

const Max = 100001

var arr = make([]int, Max)

var help = make([]int, Max)

// 暴力解法，但是时间复杂度为O(n^2)
func calArray1(length int) int64 {
	dp := make([]int, length)
	ans := 0
	dp[0] = 0
	for i := 1; i < length; i++ {
		for j := i - 1; j >= 0; j-- {
			if arr[j] <= arr[i] {
				dp[i] = dp[i] + arr[j]
			}
		}
	}
	for i := 0; i < len(dp); i++ {
		ans = ans + dp[i]
	}
	return int64(ans)
}

func calArray2(left, right int) int64 {
	if left == right {
		return 0
	}
	mid := (left + right) >> 1
	return calArray2(left, mid) + calArray2(mid+1, right) + merge(left, mid, right)
}

func merge(left, mid, right int) int64 {
	var ans int64
	// 默认右边基本大于左边，计算的公式如下
	// 在遍历右边的时候，依次遍历左边，若左侧小于等于右侧的数，则相加，否则不加
	// 如 1 3 5 | 2 4 6
	// 对2来说，只有1小于，sum=1,ans=1
	// 对4来说，右1,3小于，sum=4,ans=5
	// 对6来说，有1,3,5小于，sum=9,ans=14
	// 所以返回的merge为14
	for i, j, sum := left, mid+1, 0; j <= right; j++ {
		for i <= mid && arr[i] <= arr[j] {
			sum += arr[i]
			i++
		}
		ans += int64(sum)
	}
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

	return ans
}

func AcmMode() {
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
			arr[i], _ = strconv.Atoi(str[i])
		}
		// 将结果写入缓冲输出区
		output.WriteString(strconv.Itoa(int(calArray1(len(str)))) + " ")
		output.WriteString(strconv.FormatInt(calArray2(0, len(str)-1), 10))
		//TODO
	}
	// 输出答案
	output.Flush()
}
