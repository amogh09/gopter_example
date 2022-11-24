package bst

import (
	"math"
	"reflect"
	"testing"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"golang.org/x/exp/slices"
)

func toBST(nums []int) gopter.Gen {
	if len(nums) == 0 {
		return func(gp *gopter.GenParameters) *gopter.GenResult {
			result := gopter.NewEmptyResult(reflect.TypeOf((*TreeNode)(nil)))
			result.Sieve = func(i interface{}) bool { return true }
			return result
		}
	}

	return gen.IntRange(0, len(nums)-1).FlatMap(func(v interface{}) gopter.Gen {
		i := v.(int)
		lefts := nums[:i]
		rights := nums[i+1:]
		return toBST(lefts).FlatMap(func(v interface{}) gopter.Gen {
			var left *TreeNode
			if v == nil {
				left = nil
			} else {
				left = v.(*TreeNode)
			}
			return toBST(rights).Map(func(right *TreeNode) *TreeNode {
				return &TreeNode{
					val:   nums[i],
					left:  left,
					right: right,
				}
			})
		}, reflect.TypeOf((*TreeNode)(nil)))
	}, reflect.TypeOf(int(0)))
}

func treeGen() gopter.Gen {
	return gen.SliceOf(gen.IntRange(0, math.MaxInt32)).FlatMap(func(v interface{}) gopter.Gen {
		nums := v.([]int)
		slices.Sort(nums)
		return toBST(nums)
	}, reflect.TypeOf(([]int)(nil)))
}

func TestBst(t *testing.T) {
	params := gopter.DefaultTestParameters()
	params.MinSize = 1
	params.MinSuccessfulTests = 10
	properties := gopter.NewProperties(params)
	properties.Property("test", prop.ForAll(func(tree *TreeNode) bool {
		t.Logf("tree: %v", tree)
		return true
	}, treeGen()))
	properties.TestingRun(t)
}
