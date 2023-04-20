package leetcode

func getLeastNumbers(arr []int, k int) []int {
	QuickSort(arr, 0, len(arr)-1)
	var ans []int
	for i := 0; i < k; i++ {
		ans = append(ans, arr[i])
	}
	return ans
}
func QuickSort(array []int, begin, end int) {
	if begin < end {
		loc := partition(array, begin, end)
		QuickSort(array, begin, loc-1)
		QuickSort(array, loc+1, end)
	}
}
func partition(array []int, begin, end int) int {
	i := begin + 1 //array[begin]为基准数
	j := end       //最后一Printf

	for i < j {
		if array[i] > array[begin] {
			array[i], array[j] = array[j], array[i] //大于基准数交换
			j--                                     //左移
		} else {
			i++ //小于基准数不变，右移
		}
	}
	if array[i] >= array[begin] {
		i--
	}
	array[i], array[begin] = array[begin], array[i]
	return i
}
