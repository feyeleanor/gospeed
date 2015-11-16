package gospeed

import "testing"

func BenchmarkBaselineForLoopIteration1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < 1; j++ {
		}
	}
}

func BenchmarkBaselineForLoopIteration10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10; j++ {
		}
	}
}

func BenchmarkBaselineForLoopIteration100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < 100; j++ {
		}
	}
}

func BenchmarkBaselineForLoopIteration1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < 1000; j++ {
		}
	}
}

func BenchmarkBaselineForLoopIteration10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
		}
	}
}

func BenchmarkBaselineForLoopIteration100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < 100000; j++ {
		}
	}
}

func BenchmarkBaselineForReverseLoopIteration1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 1; j > 0; j-- {
		}
	}
}

func BenchmarkBaselineForReverseLoopIteration10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 10; j > 0; j-- {
		}
	}
}

func BenchmarkBaselineForReverseLoopIteration100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 100; j > 0; j-- {
		}
	}
}

func BenchmarkBaselineForReverseLoopIteration1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 1000; j > 0; j-- {
		}
	}
}

func BenchmarkBaselineForReverseLoopIteration10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 10000; j > 0; j-- {
		}
	}
}

func BenchmarkBaselineForReverseLoopIteration100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 100000; j > 0; j-- {
		}
	}
}

func BenchmarkBaselineForSliceRange1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, _ = range s {
		}
	}
}

func BenchmarkBaselineForSliceRange10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, _ = range s10 {
		}
	}
}

func BenchmarkBaselineForSliceRange100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, _ = range s100 {
		}
	}
}

func BenchmarkBaselineForSliceRange1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, _ = range s1000 {
		}
	}
}

func BenchmarkBaselineForSliceRange10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, _ = range s10000 {
		}
	}
}

func BenchmarkBaselineForSliceRange100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, _ = range s100000 {
		}
	}
}

func BenchmarkBaselineForSliceLength1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(s); j++ {
		}
	}
}

func BenchmarkBaselineForSliceLength10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(s10); j++ {
		}
	}
}

func BenchmarkBaselineForSliceLength100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(s100); j++ {
		}
	}
}

func BenchmarkBaselineForSliceLength1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(s1000); j++ {
		}
	}
}

func BenchmarkBaselineForSliceLength10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(s10000); j++ {
		}
	}
}

func BenchmarkBaselineForSliceLength100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(s100000); j++ {
		}
	}
}

func BenchmarkBaselineForReverseSliceLength1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := len(s); j > 0; j-- {
		}
	}
}

func BenchmarkBaselineForReverseSliceLength10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := len(s10); j > 0; j-- {
		}
	}
}

func BenchmarkBaselineForReverseSliceLength100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := len(s100); j > 0; j-- {
		}
	}
}

func BenchmarkBaselineForReverseSliceLength1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := len(s1000); j > 0; j-- {
		}
	}
}

func BenchmarkBaselineForReverseSliceLength10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := len(s10000); j > 0; j-- {
		}
	}
}

func BenchmarkBaselineForReverseSliceLength100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := len(s100000); j > 0; j-- {
		}
	}
}
