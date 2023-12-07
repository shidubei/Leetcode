package array

import (
	"math"
	"sort"
)

// 题目描述
/*
题目:给定一个长度为n的整数数组nums和一个目标值target。
从nums数组中选出3个整数，使它们的和与target最接近

解法1: 利用排序
将数组从小到大进行排序，然后每次选择3个，少则进，多则退*/
func threeSumCloset(nums []int, target int) int {
	n, res, diff := len(nums), 0, math.MaxInt
	if n > 2 {
		sort.Ints(nums)
		for i := 0; i < n-2; i++ {
			if i > 0 && nums[i] == nums[i-1] {
				continue
			}
			for j, k := i+1, n-1; j < k; {
				sum := nums[i] + nums[j] + nums[k]
				if abs(sum-target) < diff {
					res, diff = sum, abs(sum-target)
				}
				if sum == target {
					return res
				} else if sum > target {
					k--
				} else {
					j++
				}
			}
		}
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}
