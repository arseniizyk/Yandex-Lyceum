package main

func Join(nums1, nums2 []int) []int {
	joined := make([]int, 0, len(nums1)+len(nums2))
	joined = append(joined, nums1...)
	joined = append(joined, nums2...)

	return joined
}
