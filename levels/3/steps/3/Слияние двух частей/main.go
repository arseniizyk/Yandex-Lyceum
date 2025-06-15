package main

func Mix(nums []int) []int {
	mixed := make([]int, len(nums))

	for i := range nums {
		if i%2 == 0 {
			mixed[i] = nums[i/2]
		} else {
			mixed[i] = nums[len(nums)/2+i/2]
		}
	}

	return mixed
}
