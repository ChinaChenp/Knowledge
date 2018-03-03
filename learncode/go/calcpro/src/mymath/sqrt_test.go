package mymath

import "testing"

func TestSqrt1(t *testing.T) {
	v := Sqrt(10)
	if v != 3 {
		t.Errorf("Sqrt(16) failed. Got %v, expected 4.", v)
	}
}
