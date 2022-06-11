package gospeed

import (
	"fmt"
	"testing"
)

func BenchmarkBaselineFunctionCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f()
	}
}

func BenchmarkBaselineFunctionCallArg1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		farg1(0)
	}
}

func BenchmarkBaselineFunctionCallArg2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		farg2(0, 1)
	}
}

func BenchmarkBaselineFunctionCallArg3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		farg3(0, 1, 2)
	}
}

func BenchmarkBaselineFunctionCallArg4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		farg4(0, 1, 2, 3)
	}
}

func BenchmarkBaselineFunctionCallArg5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		farg5(0, 1, 2, 3, 4)
	}
}

func BenchmarkBaselineFunctionCallArg6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		farg6(0, 1, 2, 3, 4, 5)
	}
}

func BenchmarkBaselineFunctionCallArg7(b *testing.B) {
	for i := 0; i < b.N; i++ {
		farg7(0, 1, 2, 3, 4, 5, 6)
	}
}

func BenchmarkBaselineFunctionCallArg8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		farg8(0, 1, 2, 3, 4, 5, 6, 7)
	}
}

func BenchmarkBaselineFunctionCallArg9(b *testing.B) {
	for i := 0; i < b.N; i++ {
		farg9(0, 1, 2, 3, 4, 5, 6, 7, 8)
	}
}

func BenchmarkBaselineFunctionCallArg10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		farg10(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	}
}

func BenchmarkBaselineFunctionCallInt1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fint1(0)
	}
}

func BenchmarkBaselineFunctionCallInt2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fint2(0, 1)
	}
}

func BenchmarkBaselineFunctionCallInt3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fint3(0, 1, 2)
	}
}

func BenchmarkBaselineFunctionCallInt4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fint4(0, 1, 2, 3)
	}
}

func BenchmarkBaselineFunctionCallInt5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fint5(0, 1, 2, 3, 4)
	}
}

func BenchmarkBaselineFunctionCallInt6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fint6(0, 1, 2, 3, 4, 5)
	}
}

func BenchmarkBaselineFunctionCallInt7(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fint7(0, 1, 2, 3, 4, 5, 6)
	}
}

func BenchmarkBaselineFunctionCallInt8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fint8(0, 1, 2, 3, 4, 5, 6, 7)
	}
}

func BenchmarkBaselineFunctionCallInt9(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fint9(0, 1, 2, 3, 4, 5, 6, 7, 8)
	}
}

func BenchmarkBaselineFunctionCallInt10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fint10(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	}
}

func BenchmarkBaselineFunctionCallVarArgs(b *testing.B) {
	for _, v := range scaling_slices {
		vv := make([]any, len(v))
		for i, w := range v {
			vv[i] = w
		}
		b.Run(fmt.Sprintf("count_%v", len(vv)), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fvarargs(vv...)
			}
		})
	}
}

func BenchmarkBaselineFunctionCallVarInts(b *testing.B) {
	for _, v := range scaling_slices {
		b.Run(fmt.Sprintf("count_%v", len(v)), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fvarints(v...)
			}
		})
	}
}

func BenchmarkBaselineFunctionInlineCallWithDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		func() {
			defer func() {}()
		}()
	}
}

func BenchmarkBaselineFunctionVariableCallWithDefer(b *testing.B) {
	b.StopTimer()
	f := func() {
		defer func() {}()
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		f()
	}
}

func BenchmarkBaselineDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		defer func() {}()
	}
}
