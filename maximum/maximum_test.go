package maximum

import (
	"testing"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/slices"
)

func TestMaximum(t *testing.T) {
	properties := gopter.NewProperties(&gopter.TestParameters{MinSize: 1})

	properties.Property("maximum is smaller than no element", prop.ForAll(
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

	properties.Property("maximum must be present in the slice", prop.ForAll(
		func(nums []int) bool {
			return slices.Contains(nums, Maximum(nums))
		},
		gen.SliceOf(gen.Int()),
	))

	properties.TestingRun(t)
}

// Typical example based test for maximum function
func TestMaximumExampleBased(t *testing.T) {
	t.Run("maximum should be 3", func(t *testing.T) {
		assert.Equal(t, 3, Maximum([]int{1, 2, 3}))
	})
	t.Run("maximum should be 5", func(t *testing.T) {
		assert.Equal(t, 5, Maximum([]int{5, 1, 2, 3}))
	})
}
