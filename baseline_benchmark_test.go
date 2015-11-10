package gospeed

import "testing"

func BenchmarkBaselineCastInt32ToInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = int(x32)
	}
}

func BenchmarkBaselineCastIntToInt32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = int32(x)
	}
}

func BenchmarkBaselineCastInt32ToUInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = uint(x32)
	}
}

func BenchmarkBaselineCastUIntToInt32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = int32(u)
	}
}

func BenchmarkBaselineCastInt64ToUint64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = uint64(x64)
	}
}

func BenchmarkBaselineCastUint64ToInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = int64(u64)
	}
}

func BenchmarkBaselineIf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if i >= 0 {
		}
	}
}

func BenchmarkBaselineIfElse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if i < 0 {
		} else {
		}
	}
}

func BenchmarkBaselineSwitchDefault(b *testing.B) {
	for i := 0; i < b.N; i++ {
		switch i {
		default:
			{
			}
		}
	}
}

func BenchmarkBaselineSwitchOneCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		switch i {
		case 0:
			{
			}
		}
	}
}

func BenchmarkBaselineSwitchTwoCases(b *testing.B) {
	for i := 0; i < b.N; i++ {
		switch i {
		case 0:
			{
			}
		case 1:
			{
			}
		}
	}
}

func BenchmarkBaselineSwitchTwoCasesFallthrough(b *testing.B) {
	for i := 0; i < b.N; i++ {
		switch i {
		case 0:
			fallthrough
		case 1:
			{
			}
		}
	}
}
