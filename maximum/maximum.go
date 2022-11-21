package maximum

import "math"

func Maximum(nums []int) int {
	if len(nums) == 0 {
		panic("cannot find maximum of an empty slice")
	}

	res := math.MinInt
	for _, n := range nums {
		if n > res {
			res = n
		}
	}

	return res
}
