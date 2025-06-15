package main

func SortAndMerge(left, right []int) []int {
	sortedLeft := divide(left)
	sortedRight := divide(right)
	return merge(sortedLeft, sortedRight)
}

func divide(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	mid := len(nums) / 2
	left := divide(nums[:mid])
	right := divide(nums[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	sorted := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			sorted = append(sorted, left[i])
			i++

		} else {
			sorted = append(sorted, right[j])
			j++
		}
	}

	sorted = append(sorted, left[i:]...)
	sorted = append(sorted, right[j:]...)

	return sorted
}
