package gospeed

import "testing"

func BenchmarkBaselineNewMapLiteralIntToInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = map[int]int{}
	}
}

func BenchmarkBaselineNewMapLiteralIntToInterface(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = map[int]any{}
	}
}

func BenchmarkBaselineNewMapLiteralStringToInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = map[string]int{}
	}
}

func BenchmarkBaselineNewMapLiteralStringToInterface(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = map[string]any{}
	}
}

func BenchmarkBaselineNewMapLiteralIntToInt2Item(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = map[int]int{i: 0, i: 1}
	}
}

func BenchmarkBaselineNewMapLiteralIntToInterface2Item(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = map[int]any{i: 0, i: 1}
	}
}

func BenchmarkBaselineNewMapLiteralStringToInt2Item(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = map[string]int{"0": i, "1": i}
	}
}

func BenchmarkBaselineNewMapLiteralStringToInterface2Item(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = map[string]any{"0": i, "1": i}
	}
}

func BenchmarkBaselineNewMapIntToInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = make(map[int]int)
	}
}

func BenchmarkBaselineNewMapIntToInterface(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = make(map[int]any)
	}
}

func BenchmarkBaselineNewMapStringToInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = make(map[string]int)
	}
}

func BenchmarkBaselineNewMapStringToInterface(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = make(map[string]any)
	}
}

func BenchmarkBaselineMapIntGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = h[0]
	}
}

func BenchmarkBaselineMapIntSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		h[0] = 1
	}
}

func BenchmarkBaselineMapStringGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = hs["0"]
	}
}

func BenchmarkBaselineMapStringSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hs["0"] = 1
	}
}
