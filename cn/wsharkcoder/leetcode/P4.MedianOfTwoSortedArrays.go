package leetcode

import "math"

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if nums1 == nil || nums2 == nil || len(nums1)+len(nums2) == 0 {
		return 0
	}
	// l 左中位数下标 r 右中位数下标
	m, n, i, j, l, r := len(nums1), len(nums2), 0, 0, int((len(nums1)+len(nums2)-1)/2), int(math.Floor(float64((len(nums1)+len(nums2))/2)))

	sum := 0

	for i+j <= r {
		if i >= m || j >= n {
			break
		}
		if nums1[i] < nums2[j] {
			if i+j == l {
				sum += nums1[i]
			}
			if i+j == r {
				sum += nums1[i]
			}
			i++
		} else {

			if i+j == l {
				sum += nums2[j]
			}
			if i+j == r {
				sum += nums2[j]
			}
			j++

		}
	}

	for j < n {
		if i+j == l {
			sum += nums2[j]
		}
		if i+j == r {
			sum += nums2[j]

		}
		j++
	}

	for i < m {
		if i+j == l {
			sum += nums1[i]
		}
		if i+j == r {
			sum += nums1[i]
		}
		i++
	}
	return float64(sum) / 2

}
