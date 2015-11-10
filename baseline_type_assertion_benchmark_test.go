package gospeed

import "testing"

func BenchmarkBaselineTypeAssertion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = in.(int)
	}
}

func BenchmarkBaselineTypeAssertionEmptyInterface(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = in.(interface{})
	}
}

func BenchmarkBaselineTypeAssertionInterface1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = di.(dummyInterface1)
	}
}

func BenchmarkBaselineTypeAssertionInterface2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = di.(dummyInterface2)
	}
}

func BenchmarkBaselineTypeCheck(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, ok := in.(int); ok {
		}
	}
}

func BenchmarkBaselineTypeCheckEmptyInterface(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, ok := in.(interface{}); ok {
		}
	}
}

func BenchmarkBaselineTypeCheckInterface1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, ok := di.(dummyInterface1); ok {
		}
	}
}

func BenchmarkBaselineTypeCheckInterface2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, ok := di.(dummyInterface2); ok {
		}
	}
}

func BenchmarkBaselineTypeSwitchOneCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		switch in.(type) {
		case int:
		}
	}
}

func BenchmarkBaselineTypeSwitchBasicTypesCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		switch in.(type) {
		case uint8:
		case uint16:
		case uint32:
		case uint64:
		case uint:
		case float32:
		case float64:
		case complex64:
		case complex128:
		case int8:
		case int16:
		case int32:
		case int64:
		case int:
		}
	}
}

func BenchmarkBaselineTypeSwitchEmptyInterface(b *testing.B) {
	for i := 0; i < b.N; i++ {
		switch in.(type) {
		case interface{}:
		}
	}
}

func BenchmarkBaselineTypeSwitchInterface1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		switch di.(type) {
		case dummyInterface1:
		}
	}
}

func BenchmarkBaselineTypeSwitchInterface2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		switch di.(type) {
		case dummyInterface2:
		}
	}
}
