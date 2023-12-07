package array

/*
峰值元素是指其值严格大于左右相邻值的元素。
给你一个整数数组 nums，找到峰值元素并返回其索引。数组可能包含多个峰值，在这种情况下，返回任何一个峰值所在位置即可。
*/
func findPeakElement(nums []int) int {
	low, mid, high := 1, 0, len(nums)-2
	ans := 0
	if len(nums) == 1 {
		return 0
	}
	if nums[0] > nums[1] {
		return 0
	}
	if nums[len(nums)-1] > nums[len(nums)-2] {
		return len(nums) - 1
	}
	for low <= high {
		mid = (low + high) / 2
		if nums[mid-1] > nums[mid] {
			high = mid - 1
		} else if nums[mid] < nums[mid+1] {
			low = mid + 1
		} else {
			ans = mid
			break
		}
	}
	return ans
}
