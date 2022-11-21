package largest_number

import (
	"sort"
	"strconv"
	"strings"
)

func LargestNumber(nums []int) string {
	slc := make([]string, len(nums))
	for i, n := range nums {
		slc[i] = strconv.Itoa(n)
	}
	sort.Slice(slc, func(i, j int) bool {
		return greaterThan(slc[i], slc[j])
	})
	return strings.Join(slc, "")
}

func greaterThan(x, y string) bool {
	i, j := 0, 0
	for i < len(x) || j < len(y) {
		if x[i] > y[j] {
			return true
		} else if x[i] < y[j] {
			return false
		} else {
			i = min(i+1, len(x)-1)
			j = min(j+1, len(y)-1)
		}
	}
	return true
}

func min(x, y int) int {
	if x <= y {
		return x
	} else {
		return y
	}
}
