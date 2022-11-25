package bst

import (
	"fmt"
	"math"
	"reflect"
	"testing"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"golang.org/x/exp/slices"
)

// A generator that always generates nil
func nilGen() gopter.Gen {
	// generate a nil value. Couldn't find a better way to do this.
	return func(gp *gopter.GenParameters) *gopter.GenResult {
		result := gopter.NewEmptyResult(reflect.TypeOf((*TreeNode)(nil)))
		result.Sieve = func(i interface{}) bool { return true } // all values accepted
		return result
	}
}

// Generates a random Binary Search Tree with the given keys.
func toBST(keys []int) gopter.Gen {
	if len(keys) == 0 {
		return nilGen()
	}

	// choose a random index as pivot
	return gen.
		IntRange(0, len(keys)-1).
		FlatMap(func(v interface{}) gopter.Gen {
			i := v.(int)
			lefts := keys[:i]    // values for left subtrree
			rights := keys[i+1:] // values for right subtree
			return toBST(lefts).FlatMap(func(v interface{}) gopter.Gen {
				var left *TreeNode
				if v == nil {
					left = nil
				} else {
					left = v.(*TreeNode)
				}
				return toBST(rights).Map(func(right *TreeNode) *TreeNode {
					return &TreeNode{
						val:   keys[i],
						left:  left,
						right: right,
					}
				})
			}, reflect.TypeOf((*TreeNode)(nil)))
		}, reflect.TypeOf(int(0)))
}

// Contains a Binary Search Tree and all its keys in in-order
type bstWithKeys struct {
	bst  *TreeNode
	keys []int
}

// Generates a Binary Search Tree with keys
func bstWithKeysGen() gopter.Gen {
	return gen.SliceOf(gen.IntRange(0, 30)).FlatMap(func(v interface{}) gopter.Gen {
		nums := v.([]int)
		slices.Sort(nums)
		return toBST(nums).Map(func(bst *TreeNode) *bstWithKeys {
			return &bstWithKeys{bst: bst, keys: nums}
		})
	}, reflect.TypeOf(([]int)(nil)))
}

// Helper function for comparing two pointer values.
// Returns an empty string to indicate that the pointer values are equal or both pointers are nil.
// Returns a string describing the mismatch otherwise.
func pointerEqual[T comparable](x *T, y *T) string {
	if x == nil && y == nil {
		return ""
	} else if x == nil {
		return fmt.Sprintf("<nil> != %v", *y)
	} else if y == nil {
		return fmt.Sprintf("%v != <nil>", *x)
	} else if *x != *y {
		return fmt.Sprintf("%v != %v", *x, *y)
	} else {
		return ""
	}
}

func TestFirstKeyGreaterThan(t *testing.T) {
	params := gopter.DefaultTestParameters()
	params.MinSize = 1
	properties := gopter.NewProperties(params)
	type bstWithKeysWithInt struct {
		bstWithKeys bstWithKeys
		n           int
	}
	properties.Property("test", prop.ForAll(func(bstKeysInt bstWithKeysWithInt) string {
		bstKeys := bstKeysInt.bstWithKeys
		n := bstKeysInt.n
		// t.Logf("tree: %v", bstKeys.bst)
		var expected *int = nil
		for _, key := range bstKeys.keys {
			if key > n {
				expected = &key
				break
			}
		}
		actual := bstKeys.bst.FirstKeyGreaterThan(n)
		return pointerEqual(expected, actual)
	}, bstWithKeysGen().FlatMap(func(v interface{}) gopter.Gen {
		bstKeys := v.(*bstWithKeys)
		lows := gen.IntRange(math.MinInt, bstKeys.keys[0]-1)
		mids := gen.IntRange(bstKeys.keys[0], bstKeys.keys[len(bstKeys.keys)-1])
		highs := gen.IntRange(bstKeys.keys[len(bstKeys.keys)-1], math.MaxInt)
		return gen.Weighted([]gen.WeightedGen{
			{Weight: 1, Gen: lows},
			{Weight: 1, Gen: mids},
			{Weight: 1, Gen: highs},
		}).Map(func(n int) bstWithKeysWithInt { return bstWithKeysWithInt{bstWithKeys: *bstKeys, n: n} })
	}, reflect.TypeOf((*bstWithKeys)(nil)))))
	properties.TestingRun(t)
}
