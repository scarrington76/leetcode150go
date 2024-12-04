package merge_sorted_array

import "sort"

func merge(nums1 []int, m int, nums2 []int, n int) []int {

	if n == 0 {
		return nums1
	}

	nums1 = append(nums1[:m], nums2...)
	sort.Ints(nums1)
	return nums1
}
