package Sort

import "math/rand"

func QuickSort(array []int, begin, end int) []int {
	if begin < end {
		index := partition(array, begin, end)
		QuickSort(array, begin, index-1)
		QuickSort(array, index+1, end)
	}
	return array
}

// 更新返回基准值的index
func partition(array []int, begin, end int) int {
	// 以array[begin]作为基准数
	i := begin + 1 // i记录基准值左边
	j := end       // j记录基准值右边

	for i < j {
		if array[i] < array[begin] {
			// 大于基准值，交换
			array[i], array[j] = array[j], array[i]
			j--
		} else {
			// 小于基准值，不交换
			i++
		}
	}
	// 循环结束，检查i是否满足
	if array[i] <= array[begin] {
		i--
	}
	array[i], array[begin] = array[begin], array[i]
	return i
}

// 随机快排算法，将时间复杂度降到了O(NlogN)
var first int
var last int

func quickSort2(arr []int, left, right int) {
	if left > right {
		return
	}
	x := arr[left+rand.Intn(right-left+1)]
	partition2(arr, left, right, x)
	l, r := first, last
	quickSort2(arr, left, l-1)
	quickSort2(arr, r+1, right)
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
