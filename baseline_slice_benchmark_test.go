package gospeed

import "testing"

var s []int = []int{0}
var s10 []int = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
var s100 []int = make([]int, 100, 100)
var s1000 []int = make([]int, 1000, 1000)
var s10000 []int = make([]int, 10000, 10000)
var s100000 []int = make([]int, 100000, 100000)
var sd []int = []int{0}
var sd10 []int = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
var sd100 []int = make([]int, 100, 100)
var sd1000 []int = make([]int, 1000, 1000)
var sd10000 []int = make([]int, 10000, 10000)
var sd100000 []int = make([]int, 100000, 100000)

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

func BenchmarkBaselineNewSliceAppendElement100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = append(s, s100...)
	}
}

func BenchmarkBaselineNewSliceAppendElement1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = append(s, s1000...)
	}
}

func BenchmarkBaselineNewSliceAppendElement10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = append(s, s10000...)
	}
}

func BenchmarkBaselineNewSliceAppendElement100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = append(s, s100000...)
	}
}

func BenchmarkBaselineSliceGetAtOffset(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = s10[5]
	}
}

func BenchmarkBaselineSliceGetAtCalculatedOffset(b *testing.B) {
	base, offset := 2, 3
	for i := 0; i < b.N; i++ {
		_ = s10[base+offset]
	}
}

func BenchmarkBaselineSliceGetAtOffsetWithBase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = s10[2:][3]
	}
}

func BenchmarkBaselineSliceSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s10[5] = 1
	}
}

func BenchmarkBaselineSliceCopy1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		copy(sd, s)
	}
}

func BenchmarkBaselineSliceCopy10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		copy(sd10, s10)
	}
}

func BenchmarkBaselineSliceCopy100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		copy(sd100, s100)
	}
}

func BenchmarkBaselineSliceCopy1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		copy(sd1000, s1000)
	}
}

func BenchmarkBaselineSliceCopy10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		copy(sd10000, s10000)
	}
}

func BenchmarkBaselineSliceCopy100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		copy(sd100000, s100000)
	}
}
