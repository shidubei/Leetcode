package stack

/*题目：给一个字符串表达式s，实现一个基本计算器来计算并返回它的值。*/
/*思路：很经典的辅助栈的题目，利用两个辅助栈分别记录数字和运算符号*/

// 错误1：以单个字符存储的时候，如果存在计算结果是大于2位数的时候，会出现错误
// 解决1：增加一个记位符，都是数字的时候++，那么大于1的时候更新栈顶元素
// 错误2：当存在1+(-2+1)的时候。首先存在1.负号在首位的问题，导致运算符和数字相等2.返回的负数会将式子变为1+-1，存在连续符号的问题
// 解决2.1：对于-号首位，将字符串前面添一个0.解决2：连续字符匹配，通过一个bool值记录前一个是数字还是符号，如果前一个是符号，遵循相同的+。相异得-来记录符号

//func Calculate(s string) int {
//	// 辅助栈,记录带符号的数字
//	stack := []byte{}
//	for i := 0; i < len(s); i++ {
//		if s[i] == ' ' {
//			continue
//		} else if s[i] == ')' {
//			// 先运算括号中答案
//			temp, index := "", len(stack)-1
//			for ; index >= 0; index-- {
//				// 找到左括号
//				if stack[index] == '(' {
//					break
//				}
//			}
//			temp = string(stack[index+1:]) // 记录括号中的字符串
//			stack = stack[:index]          //stack将括号之后的内容全部释放
//			// 记录结果
//			res := strconv.Itoa(calculateStr(temp))
//			for j := 0; j < len(res); j++ {
//				stack = append(stack, res[j])
//			}
//		} else {
//			stack = append(stack, s[i]) //放入字符
//		}
//	}
//	// 所有的字符都放完，开始计算且返回
//	return calculateStr(string(stack))
//}
//func calculateStr(s string) int {
//	// 辅助栈，sign记录符号,num记录数字
//	sign, nums := []byte{}, []int{}
//	cnt := 0 // 记位符
//	isDigit := false
//	if s[0] == '-' {
//		s = "0" + s
//	}
//	for i := 0; i < len(s); i++ {
//		if s[i] >= '0' && s[i] <= '9' {
//			num, _ := strconv.Atoi(string(s[i]))
//			cnt++        // 连续的数字，位数++
//			if cnt > 1 { //如果位数大于1
//				nums[len(nums)-1] = nums[len(nums)-1]*10 + num
//				continue // 不入栈，更新栈顶元素，继续循环
//			}
//			nums = append(nums, num) // 数字栈
//			isDigit = true
//		} else {
//			cnt = 0 // 符号为分割，每遇到符号，记位符清0
//			if len(sign) > 0 && sign[len(sign)-1] == '+' && s[i] == '+' && !isDigit {
//				continue
//			} else if len(sign) > 0 && sign[len(sign)-1] == '+' && s[i] == '-' && !isDigit {
//				sign[len(sign)-1] = '-'
//			} else if len(sign) > 0 && sign[len(sign)-1] == '-' && s[i] == '-' && !isDigit {
//				sign[len(sign)-1] = '+'
//			} else if len(sign) > 0 && sign[len(sign)-1] == '-' && s[i] == '+' && !isDigit {
//				continue
//			} else {
//				sign = append(sign, s[i])
//			}
//			isDigit = false
//		}
//	}
//	res := 0
//	if len(sign) == 0 {
//		return nums[len(nums)-1]
//	}
//	for i := 0; i < len(sign); i++ {
//		if sign[i] == '+' {
//			// 进行加法运算
//			res = nums[i] + nums[i+1]1
//			nums[i+1] = res
//		} else if sign[i] == '-' {
//			res = nums[i] - nums[i+1]
//			nums[i+1] = res
//		}
//	}
//	return res
//}
