package random

import "testing"

func TestGetRandomNum(t *testing.T) {
	t.Log(GetRandomNum(6))
	t.Log(GetRandomNum(5))
	t.Log(GetRandomNum(3))
}

func BenchmarkGetRandomNum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = GetRandomNum(6)
	}
}

func TestGetRandomNum1(t *testing.T) {
	println(GetRandomNum(6))
	println(GetRandomNum(66))
	println(GetRandomNum(67))
}
