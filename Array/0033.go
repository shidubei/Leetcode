package array

func search33(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > nums[low] {
			if nums[low] <= target && target < nums[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else if nums[mid] < nums[high] {
			if nums[high] >= target && nums[mid] < target {
				low = mid + 1
			} else {
				high = mid - 1
			}
		} else {
			if nums[low] == nums[mid] {
				low++
			} else if nums[high] == nums[mid] {
				high--
			}
		}
	}
	return -1
}
