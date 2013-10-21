package gospeed

import "testing"

func BenchmarkPanic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		func() {
			defer func() {
				b.StopTimer()
				recover()
			}()
			b.StartTimer()
			panic(nil)
		}()
	}
}

func BenchmarkPanicRecover(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		func() {
			defer func() {
				b.StartTimer()
				recover()
			}()
			panic(nil)
		}()
	}
}

func BenchmarkPanicFunctionCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		func() {
			defer func() {
				_ = recover()
			}()
			panic(nil)
		}()
	}
}