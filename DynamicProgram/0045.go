package DynamicProgram

/*题目：给一个数组，数组中的数字表示从该位置可以跳的最大距离，请找出一个能到达数组结尾跳跃次数最少的次数
如：{2,3,0,1,4}最短为2，即下标0到下标1，到下标4*/

/*
思路：
1.肯定需要遍历数组，数组的下标和值的结合就是当前位置能跳跃的最远位置，如果这个最远位置大于等于了数组长度，那么直接当前步数加1
2.实时的去检测每个位置能到的最远位置，且记录当前位置能到达的最远位置，在最远位置没有出界的情况下，一直遍历到最远位置还为出界时需要再进行跳跃，step++
*/
func jump(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	step, canReach, longest := 0, 0, 0 // 初始化步数，可以到的最远位置和当前位置需要判定的最远位置
	// 开始遍历
	for i, x := range nums {
		if i+x > canReach {
			// 更新canReach
			canReach = i + x
			if canReach >= len(nums)-1 {
				return step + 1
			}
		}
		if i == longest {
			// 如果遍历到最远位置，当前最远位置更新，步数需要加1
			longest = canReach
			step++
		}
	}
	return step
}
