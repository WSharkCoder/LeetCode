package leetcode

// 解法一、双循环暴力解 O(n^2)
//func twoSum(nums []int, target int) []int {
//	var result []int = make([]int, 2)
//	for i := 0; i < len(nums)-1; i++ {
//		for j := i + 1; j < len(nums); j++ {
//			if nums[i]+nums[j] == target {
//				result[0] = i
//				result[1] = j
//				break
//			}
//		}
//	}
//	return result
//}

// 解法二、一遍遍历数组，一遍往map中放，然后判断另外一个目标值在不在map中 O(n)
func twoSum(nums []int, target int) []int {

	m := make(map[int]int, len(nums))

	for j := 0; j < len(nums); j++ {
		another := target - nums[j]
		// Go语言map特性,同时获取值和是否有值
		i, ok := m[another]
		if ok {
			return []int{i, j}

		}
		m[nums[j]] = j
	}
	return nil
}
