package main

func SliceCopy(nums []int) []int {
	slice := make([]int, 0, len(nums))
	slice = append(slice, nums...)
	return slice
}
