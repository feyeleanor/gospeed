package gospeed

import "testing"

func BenchmarkValueMethodCall(b *testing.B) {
	for i := 0; i < b.N; i++ { dummy.m1() }
}

func BenchmarkValueMethodCall1Arg(b *testing.B) {
	for i := 0; i < b.N; i++ { dummy.m1arg(1) }
}

func BenchmarkValueMethodCall1Int(b *testing.B) {
	for i := 0; i < b.N; i++ { dummy.m1int(1) }
}

func BenchmarkValueMethodCall5Args(b *testing.B) {
	for i := 0; i < b.N; i++ { dummy.m1varargs(1, 2, 3, 4, 5) }
}

func BenchmarkValueMethodCall5Ints(b *testing.B) {
	for i := 0; i < b.N; i++ { dummy.m1varints(1, 2, 3, 4, 5) }
}

func BenchmarkPointerMethodCall(b *testing.B) {
	for i := 0; i < b.N; i++ { dummy.m2() }
}

func BenchmarkPointerMethodCall1Arg(b *testing.B) {
	for i := 0; i < b.N; i++ { dummy.m2arg(1) }
}

func BenchmarkPointerMethodCall1Int(b *testing.B) {
	for i := 0; i < b.N; i++ { dummy.m2int(1) }
}

func BenchmarkPointerMethodCall5Args(b *testing.B) {
	for i := 0; i < b.N; i++ { dummy.m2varargs(1, 2, 3, 4, 5) }
}

func BenchmarkPointerMethodCall5Ints(b *testing.B) {
	for i := 0; i < b.N; i++ { dummy.m2varints(1, 2, 3, 4, 5) }
}

func BenchmarkValueMethodGetSlice(b *testing.B) {
	for i := 0; i < b.N; i++ { _ = dummyAccessor.get(0) }
}

func BenchmarkValueMethodSetSlice(b *testing.B) {
	for i := 0; i < b.N; i++ { dummyAccessor.set(0, 1) }
}

func BenchmarkPointerMethodGetSlice(b *testing.B) {
	for i := 0; i < b.N; i++ { _ = (&dummyAccessor).get(0) }
}

func BenchmarkPointerMethodSetSlice(b *testing.B) {
	for i := 0; i < b.N; i++ { (&dummyAccessor).set(0, 1) }
}

func BenchmarkInterfaceMethodGetSlice(b *testing.B) {
	for i := 0; i < b.N; i++ { _ = dai.get(0) }
}

func BenchmarkInterfaceMethodSetSlice(b *testing.B) {
	for i := 0; i < b.N; i++ { dai.set(0, 1) }
}