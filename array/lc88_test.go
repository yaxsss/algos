package array

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	nums := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	merge(nums, 3, nums2, 3)
	if !reflect.DeepEqual(nums, []int{1, 2, 2, 3, 5, 6}) {
		t.Fatalf("%v not equal expected [1,2,2,3,5,6]", nums)
	}

	nums = []int{1}
	nums2 = []int{}
	merge(nums, 1, nums2, 0)
	if !reflect.DeepEqual(nums, []int{1}) {
		t.Fatalf("%v not equal expected [1]", nums)
	}

	nums = []int{0}
	nums2 = []int{1}
	merge(nums, 0, nums2, 1)
	if !reflect.DeepEqual(nums, []int{1}) {
		t.Fatalf("%v not equal expected [1]", nums)
	}

}
