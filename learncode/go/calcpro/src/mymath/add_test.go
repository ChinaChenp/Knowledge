package mymath

import "testing"

func TestAdd1(t *testing.T) {
	re := Add(1, 2)
	if re != 3 {
		t.Errorf("Add(1, 2) error, get %d, expected 3.", re)
	}
}
