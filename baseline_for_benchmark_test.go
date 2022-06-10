package gospeed

import (
	"fmt"
	"testing"
)

var scaling_points []int = []int{0, 1, 10, 100, 1000, 10000, 100000, 1000000}
var scaling_slices [][]int = [][]int{s0, s, s10, s100, s1000, s10000, s100000, s1000000}

func BenchmarkBaselineForLoopIteration(b *testing.B) {
	for _, v := range scaling_points {
		b.Run(fmt.Sprintf("count_%v", v), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for j := 0; j < v; j++ {
				}
			}
		})
	}
}

func BenchmarkBaselineForReverseLoopIteration(b *testing.B) {
	for _, v := range scaling_points {
		b.Run(fmt.Sprintf("count_%v", v), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for j := v; j > 0; j-- {
				}
			}
		})
	}
}

func BenchmarkBaselineForSliceRange(b *testing.B) {
	for _, v := range scaling_slices {
		b.Run(fmt.Sprintf("count_%v", len(v)), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for _, _ = range v {
				}
			}
		})
	}
}

func BenchmarkBaselineForSliceLength(b *testing.B) {
	for _, v := range scaling_slices {
		b.Run(fmt.Sprintf("count_%v", len(v)), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for j := 0; j < len(s); j++ {
				}
			}
		})
	}
}

func BenchmarkBaselineForReverseSliceLength(b *testing.B) {
	for _, v := range scaling_slices {
		b.Run(fmt.Sprintf("count_%v", len(v)), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for j := len(s); j > 0; j-- {
				}
			}
		})
	}
}
