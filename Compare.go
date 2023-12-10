package main

import (
	"fmt"
	"math/rand"
)

var first int
var last int

func findKthLargest(arr []int, k int) int {
	return randomChoice(arr, len(arr)-k)
}
func randomChoice(arr []int, i int) int {
	ans := 0
	for l, r := 0, len(arr)-1; l <= r; {
		partition2(arr, l, r, arr[l+rand.Intn(r-l+1)])
		//         i
		// ...first....last...
		if i < first {
			// 边界值在所需值右侧，更新右边界
			r = first - 1
		} else if i > last {
			//边界值在所需值左侧，更新左边界
			l = last + 1
		} else {
			ans = arr[i]
			break
		}
	}
	return ans
}

func partition2(arr []int, left, right, x int) {
	first, last = left, right
	fmt.Println(x)
	fmt.Println(first)
	fmt.Println(last)
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
	fmt.Println("after", first)
	fmt.Println("after", last)
}

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println(findKthLargest(arr, 2))
}
