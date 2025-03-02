package dp

import "testing"

func TestUniquePaths(t *testing.T) {
	m := uniquePaths(7, 3)
	if m != 28 {
		t.Fatalf("m = 7, n = 3 res %d not equal %d", m, 28)
	}
	m = uniquePaths(3, 3)
	if m != 6 {
		t.Fatalf("m = 7, n = 3 res %d not equal %d", m, 6)
	}
}
