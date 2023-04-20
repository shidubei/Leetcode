package array

/*leetcode第四题：寻找两个正序数组的中位数
要求：算法时间复杂度为O(log(m+n)*/

/*思路：
1.看到时间复杂度的要求为log，第一反应为使用二分查找
2.对于二分查找来说，重点是确定二分的位置，由两个数组的长度可知最终合并的数组的二分位置在哪里
3.思路类似于插入排序，即寻找到插入二分位置的数字*/

// func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
// 	if len(nums1) > len(nums2) {
// 		return findMedianSortedArrays(nums2, nums1)
// 	}
// 	//两个正序数组的二分
// 	low, high, k, mid1, mid2 := 0, len(nums1), (len(nums1)+len(nums2)+1)/2, 0, 0
// 	for low <= high {
// 		mid1 = low + (high-low)/2
// 		mid2 = k - mid1
// 		if mid1 > 0 && nums1[mid1-1] > nums2[mid2] {
// 			high = mid1 - 1
// 		} else if mid1 != len(nums1) && nums1[mid1] < nums2[mid2-1] {
// 			low = mid1 + 1
// 		} else {
// 			break
// 		}
// 	}
// 	midleft, midright := 0, 0
// 	if mid1 == 0 {
// 		midleft = nums2[mid2-1]
// 	} else if mid2 == 0 {
// 		midleft = nums1[mid1-1]
// 	} else {
// 		midleft = max(nums1[mid1-1], nums2[mid2-1])
// 	}
// 	if (len(nums1)+len(nums2))%2 == 1 {
// 		return float64(midleft)
// 	}
// 	if mid1 == len(nums1) {
// 		midright = nums2[mid2]
// 	} else if mid2 == len(nums2) {
// 		midright = nums1[mid1]
// 	} else {
// 		midright = min(nums1[mid1], nums2[mid2])
// 	}
// 	return float64(midleft+midright) / 2
// }
// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	} else {
// 		return b
// 	}
// }
// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	} else {
// 		return b
// 	}
// }
