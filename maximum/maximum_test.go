package maximum

import (
	"testing"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"golang.org/x/exp/slices"
)

func TestMaximum(t *testing.T) {
	properties := gopter.NewProperties(&gopter.TestParameters{MinSize: 1})

	properties.Property("maximum is greater than or equal to all",
		prop.ForAll(
			func(nums []int) bool {
				m := Maximum(nums)
				for _, n := range nums {
					if m < n {
						return false
					}
				}
				return true
			},
			gen.SliceOf(gen.Int()),
		))

	properties.Property("maximum must be present in the slice",
		prop.ForAll(
			func(nums []int) bool {
				return slices.Contains(nums, Maximum(nums))
			},
			gen.SliceOf(gen.Int()),
		))

	properties.TestingRun(t)
}
