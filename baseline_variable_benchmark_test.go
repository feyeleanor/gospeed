package gospeed

import "testing"

func BenchmarkBaselineVariableGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = x
	}
}

func BenchmarkBaselineVariableSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x = 1
	}
}

func BenchmarkBaselineVariableGetInterface(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = in
	}
}

func BenchmarkBaselineVariableSetInterface(b *testing.B) {
	for i := 0; i < b.N; i++ {
		in = 1
	}
}

func BenchmarkBaselineVariableIncrement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x++
	}
}

func BenchmarkBaselineVariableDecrement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x--
	}
}

func BenchmarkBaselineVariableIntegerLeftShift(b *testing.B) {
	var i, y int
	x := uint(2)
	for ; i < b.N; i++ {
		y = i << x
	}
	i = y
}

func BenchmarkBaselineVariableIntegerRightShift(b *testing.B) {
	var i, y int
	x := uint(2)
	for ; i < b.N; i++ {
		y = i >> x
	}
	i = y
}

func BenchmarkBaselineVariableIntegerAdd(b *testing.B) {
	var i, x, y int
	for ; i < b.N; i++ {
		y = x + i
	}
	x = y
}

func BenchmarkBaselineVariableIntegerSubtract(b *testing.B) {
	var i, x, y int
	for ; i < b.N; i++ {
		y = x - i
	}
	x = y
}

func BenchmarkBaselineVariableIntegerMultiply(b *testing.B) {
	var i, y int
	x := 2
	for ; i < b.N; i++ {
		y = x * i
	}
	x = y
}

func BenchmarkBaselineVariableIntegerDivide(b *testing.B) {
	var i, y int
	x := 2
	for ; i < b.N; i++ {
		y = i / x
	}
	x = y
}

func BenchmarkBaselineVariableIntegerRemainder(b *testing.B) {
	var i, y int
	x := 2
	for ; i < b.N; i++ {
		y = i % x
	}
	x = y
}
