package Arithmetic

import (
	"fmt"
	"sort"
)

func Haffman(array []int) {
	sort.Ints(array)
	sum := 0
	for len(array) > 1 {
		cur := array[0] + array[1]
		array = array[2:]
		sum += cur
		array = append(array, cur)
		sort.Ints(array)
	}
	fmt.Println(sum)
}
