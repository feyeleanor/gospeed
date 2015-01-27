package gospeed

import "testing"

func BenchmarkPanicRecover(b *testing.B) {
	for i := 0; i < b.N; i++ {
		func() {
			defer func() {
				recover()
			}()
			panic(nil)
		}()
	}
}

func BenchmarkPanicRecoverAssign(b *testing.B) {
	for i := 0; i < b.N; i++ {
		func() {
			defer func() {
				_ = recover()
			}()
			panic(nil)
		}()
	}
}