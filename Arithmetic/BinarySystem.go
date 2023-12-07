package Arithmetic

import "fmt"

/*利用与运算来输出一个数的二进制形式i*/
func printBinary(x int) {
	for i := 31; i >= 0; i-- {
		if x&(1<<i) == 0 {
			fmt.Print(0)
		} else {
			fmt.Print(1)
		}
	}
}

// 位运算加法
func binaryAdd(a, b int) int {
	ans := a
	for b != 0 {
		ans = a ^ b
		b = (a & b) << 1
		a = ans
	}
	return ans
}

// 位运算减法，直接利用位运算加法
func binarySub(a, b int) int {
	return binaryAdd(a, binaryAdd(^b, 1))
}

// 位运算乘法
// 需要注意的是，go中没有无符号右移的操作，先将int类型转化为uint类型再进行右移操作
func binaryMultiply(a, b int) int {
	ans := 0
	ub := uint(b)
	for ub != 0 {
		if (ub & 1) != 0 {
			ans = binaryAdd(ans, a)
		}
		a = a << 1
		ub = ub >> 1
	}
	return ans
}

// 位运算除法
// 若a.b非整数最小值
func binaryDiv(a, b int) int {
	x, y, ans := 0, 0, 0
	if a < 0 {
		x = binaryAdd(^a, 1)
	} else {
		x = a
	}
	if b < 0 {
		y = binaryAdd(^b, 1)
	} else {
		y = b
	}
	for i := 30; i >= 0; i = binarySub(i, 1) {
		if x>>i >= y {
			ans = ans | 1<<i
			x = binarySub(x, y<<i)
		}
	}
	if binaryMultiply(a, b) < 0 {
		return binaryAdd(^ans, 1)
	} else {
		return ans
	}
}
