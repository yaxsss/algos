package array

import "testing"

func TestMajorityElement(t *testing.T) {
	e := majorityElement([]int{3, 2, 3})
	if e != 3 {
		t.Fatalf("%d not equal 3", e)
	}
	e = majorityElement([]int{2, 2, 1, 1, 1, 2, 2})
	if e != 2 {
		t.Fatalf("%d not equal 2", e)
	}
}
