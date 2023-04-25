package main

import (
	"fmt"
	"leetcode/Sorts"
	"math/rand"
	"sort"
	"time"
)

// this is a method which can check your sort func is right or error

// base on the zuochengyun's thought, the steps are as follows:

/* 1.the func that want to check */
func sortTest(arr []int) []int {
	/* please put your sort function in there */
	res := make([]int, len(arr))
	Sort.QuickSort(arr, 0, len(arr)-1)
	/* slice is a reference type,so use copy func to copy its value */
	/* then when you change the source,the new one will not change */
	copy(res, arr)
	return res
}

/* 2.the absolutely correct func but the complexity is not good */
func sortCorrect(arr []int) []int {
	res := make([]int, len(arr))
	sort.Ints(arr)
	copy(res, arr)
	return res
}

/* 3.a func which can provide random sample */
/* for Sort check,we need provide a lot of array*/
func sampleRandom(size int) []int {
	rand.Seed(time.Now().UnixNano())
	// create a array which use l as the length of array based on size
	l := rand.Intn(size)
	sam := make([]int, l)
	// fill the array with random
	for i := 0; i < l; i++ {
		si := rand.Intn(100)
		sam[i] = si
	}
	return sam
}

/* 4.compare the func you write and the absolutely correct */
func Compare(round, size int) string {
	i := 0
	for round > 0 {
		arrSam := sampleRandom(size)
		arr1 := sortTest(arrSam)
		arr2 := sortCorrect(arrSam)
		for i = 0; i < len(arrSam); i++ {
			if arr1[i] != arr2[i] {
				break
			}
			time.Sleep(time.Nanosecond)
		}
		if i != len(arrSam) {
			return "no! they are different"
		}
		round--
	}
	return "yes, success!"
}

func main() {
	round := 5
	size := 20
	fmt.Println(Compare(round, size))
}
