package main

import (
	"fmt"
)

func UnderLimit(nums []int, limit, n int) ([]int, error) {
	if nums == nil {
		return nil, fmt.Errorf("nums is nil")
	}
	if n < 0 {
		return nil, fmt.Errorf("incorrect n")
	}
	underLim := make([]int, 0)
	// sort.Ints(nums)

	for _, num := range nums {
		if len(underLim) >= n {
			break
		}
		if num < limit {
			underLim = append(underLim, num)
		}
	}

	return underLim, nil
}
