package array

// import "sort"

/*思路1：哈希记录*/
// func threeSum(nums []int) [][]int {
// 	res := [][]int{}
// 	cnt := map[int]int{}
// 	//map表记录每个数字出现的次数
// 	for _, value := range nums {
// 		cnt[value]++
// 	}

// 	num := []int{}
// 	for key := range cnt {
// 		num = append(num, key)
// 	}
// 	sort.Ints(num)

// 	for i := 0; i < len(num); i++ {
// 		if num[i]*3 == 0 && cnt[num[i]] >= 3 {
// 			res = append(res, []int{num[i], num[i], num[i]})
// 		}
// 		for j := i + 1; j < len(num); j++ {
// 			if num[i]*2+num[j] == 0 && cnt[num[i]] > 1 {
// 				res = append(res, []int{num[i], num[i], num[j]})
// 			}
// 			if num[j]*2+num[i] == 0 && cnt[num[j]] > 1 {
// 				res = append(res, []int{num[i], num[j], num[j]})
// 			}
// 			c := 0 - num[i] - num[j]
// 			if c > num[j] && cnt[c] > 0 {
// 				res = append(res, []int{num[i], num[j], c})
// 			}
// 		}
// 	}
// 	return res
// }
/*思路2：双指针+排序*/
// func threeSum(nums []int) [][]int {
// 	res := [][]int{}
// 	length := len(nums)
// 	sort.Ints(nums)

// 	for first := 0; first < length; first++ {
// 		if first > 0 && nums[first] == nums[first-1] {
// 			continue
// 		}
// 		third := length - 1
// 		target := 0 - nums[first]
// 		for second := first + 1; second < length; second++ {
// 			if second > first+1 && nums[second] == nums[second-1] {
// 				continue
// 			}
// 			for second < third && nums[second]+nums[third] > target {
// 				third--
// 			}
// 			if second == third {
// 				break
// 			}
// 			if nums[second]+nums[third] == target {
// 				res = append(res, []int{nums[first], nums[second], nums[third]})
// 			}
// 		}
// 	}
// 	return res
// }
