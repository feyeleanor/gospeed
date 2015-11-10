package gospeed

import "testing"

func BenchmarkBaselineNewSliceLiteral(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = []int{0, 1}
	}
}

func BenchmarkBaselineNewSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = make([]int, 2)
	}
}

func BenchmarkBaselineSliceGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = s[0]
	}
}

func BenchmarkBaselineSliceSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s[0] = 1
	}
}

func BenchmarkBaselineSliceCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		copy(d10, s10)
	}
}

func BenchmarkBaselineNewSliceAppendElement1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = append(s, 1)
	}
}

func BenchmarkBaselineNewSliceAppendElement10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = append(s, s10...)
	}
}
