package map_cap

import "testing"

func BenchmarkFunc1(t *testing.B) {
	for i := 0; i < t.N; i++ {
		Func1()
	}
}

func BenchmarkFunc2(t *testing.B) {
	for i := 0; i < t.N; i++ {
		Func2()
	}
}
