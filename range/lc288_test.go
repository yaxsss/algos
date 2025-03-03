package ranged

import (
	"reflect"
	"testing"
)

func TestSummaryRange(t *testing.T) {
	cases := []struct {
		arr      []int
		expected []string
	}{
		{[]int{0, 1, 2, 4, 5, 7}, []string{"0->2", "4->5", "7"}},
		{[]int{0, 2, 3, 4, 6, 8, 9}, []string{"0", "2->4", "6", "8->9"}},
	}

	for _, c := range cases {
		res := summaryRange(c.arr)
		if !reflect.DeepEqual(res, c.expected) {
			t.Fatalf("%v not equal %v", res, c.expected)
		}
	}
}
