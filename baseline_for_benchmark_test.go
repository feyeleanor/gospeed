package gospeed

import "testing"

func BenchmarkBaselineForLoopIteration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < 1; j++ {
		}
	}
}

func BenchmarkBaselineForReverseLoopIteration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 1; j > 0; j-- {
		}
	}
}

func BenchmarkBaselineForRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, _ = range s {
		}
	}
}

func BenchmarkBaselineForSliceLength(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(s); j++ {
		}
	}
}

func BenchmarkBaselineForReverseSliceLength(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := len(s); j > 0; j-- {
		}
	}
}

func BenchmarkBaselineForLoopIteration10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10; j++ {
		}
	}
}

func BenchmarkBaselineForReverseLoopIteration10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 10; j > 0; j-- {
		}
	}
}

func BenchmarkBaselineForRange10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, _ = range s10 {
		}
	}
}

func BenchmarkBaselineForSliceLength10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(s10); j++ {
		}
	}
}

func BenchmarkBaselineForReverseSliceLength10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := len(s10); j > 0; j-- {
		}
	}
}

func BenchmarkBaselineForLoopIteration100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < 100; j++ {
		}
	}
}

func BenchmarkBaselineForReverseLoopIteration100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 100; j > 0; j-- {
		}
	}
}

func BenchmarkBaselineForRange100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, _ = range s100 {
		}
	}
}

func BenchmarkBaselineForSliceLength100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(s100); j++ {
		}
	}
}

func BenchmarkBaselineForReverseSliceLength100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := len(s100); j > 0; j-- {
		}
	}
}

func BenchmarkBaselineForLoopIteration10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
		}
	}
}

func BenchmarkBaselineForReverseLoopIteration10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 10000; j > 0; j-- {
		}
	}
}

func BenchmarkBaselineForRange10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, _ = range s10000 {
		}
	}
}

func BenchmarkBaselineForSliceLength10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(s10000); j++ {
		}
	}
}

func BenchmarkBaselineForReverseSliceLength10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := len(s10000); j > 0; j-- {
		}
	}
}
