package gospeed

import "testing"

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

func BenchmarkBaselineFunctionCall5VarArgs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fvarargs(1, 2, 3, 4, 5)
	}
}

func BenchmarkBaselineFunctionCallInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fint(1)
	}
}

func BenchmarkBaselineFunctionCall5VarInts(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fvarints(1, 2, 3, 4, 5)
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
