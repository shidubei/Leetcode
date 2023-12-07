package Arithmetic

/*
面试题：判断一个数是否为step sum
step sum的定义 754 = 680+68+6 11=10+1

思路：二分*/

func findStepSum(num int) (bool, int) {
	low, mid, high := 1, (1+num)>>1, num
	for low < high {
		if stepSum(mid) > num {
			high = mid - 1
			mid = (low + high) >> 1
		} else if stepSum(mid) < num {
			low = mid + 1
			mid = (low + high) >> 1
		} else {
			return true, mid
		}
	}
	return false, 0
}
func stepSum(x int) int {
	sSum := x
	for x/10 != 0 {
		sSum = sSum + x/10
		x = x / 10
	}
	return sSum
}
