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

func BenchmarkBaselineFunctionCallArg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		farg(1)
	}
}

func BenchmarkBaselineFunctionCallInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fint(1)
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
