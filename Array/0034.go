package array

//func FindFirstAndLastPosition(nums []int, target int) []int {
//	return []int{findFirst(nums, target), findLast(nums, target)}
//}

//func findFirst(nums []int, target int) int {
//	low, high := 0, len(nums)-1
//	for low <= high {
//		mid := low + (high-low)/2
//		if nums[mid] < target {
//			low = mid + 1
//		} else if nums[mid] > target {
//			high = mid - 1
//		} else {
//			if mid == 0 || nums[mid-1] != target {
//				return mid
//			}
//			high = mid - 1
//		}
//	}
//	return -1
//}
//
//func findLast(nums []int, target int) int {
//	low, high := 0, len(nums)-1
//	for low <= high {
//		mid := low + (high-low)/2
//		if nums[mid] < target {
//			low = mid + 1
//		} else if nums[mid] > target {
//			high = mid - 1
//		} else {
//			if mid == len(nums)-1 || nums[mid+1] != target {
//				return mid
//			}
//			low = mid + 1
//		}
//	}
//	return -1
//}
