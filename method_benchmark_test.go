package gospeed

import (
	"fmt"
	"testing"
)

func BenchmarkValueMethodCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dummy.m1()
	}
}

func BenchmarkValueMethodCallArg1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dummy.m1arg1(1)
	}
}

func BenchmarkValueMethodCallArg2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dummy.m1arg2(1, 2)
	}
}

func BenchmarkValueMethodCallArg3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dummy.m1arg3(1, 2, 3)
	}
}

func BenchmarkValueMethodCallArg4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dummy.m1arg4(1, 2, 3, 4)
	}
}

func BenchmarkValueMethodCallArg5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dummy.m1arg5(1, 2, 3, 4, 5)
	}
}

func BenchmarkValueMethodCallArg6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dummy.m1arg6(1, 2, 3, 4, 5, 6)
	}
}

func BenchmarkValueMethodCallArg7(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dummy.m1arg7(1, 2, 3, 4, 5, 6, 7)
	}
}

func BenchmarkValueMethodCallArg8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dummy.m1arg8(1, 2, 3, 4, 5, 6, 7, 8)
	}
}

func BenchmarkValueMethodCallArg9(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dummy.m1arg9(1, 2, 3, 4, 5, 6, 7, 8, 9)
	}
}

func BenchmarkValueMethodCallArg10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dummy.m1arg10(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	}
}

func BenchmarkValueMethodCallInt1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dummy.m1int1(1)
	}
}

func BenchmarkValueMethodCallInt2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dummy.m1int2(1, 2)
	}
}

func BenchmarkValueMethodCallInt3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dummy.m1int3(1, 2, 3)
	}
}

func BenchmarkValueMethodCallInt4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dummy.m1int4(1, 2, 3, 4)
	}
}

func BenchmarkValueMethodCallInt5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dummy.m1int5(1, 2, 3, 4, 5)
	}
}

func BenchmarkValueMethodCallInt6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dummy.m1int6(1, 2, 3, 4, 5, 6)
	}
}

func BenchmarkValueMethodCallInt7(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dummy.m1int7(1, 2, 3, 4, 5, 6, 7)
	}
}

func BenchmarkValueMethodCallInt8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dummy.m1int8(1, 2, 3, 4, 5, 6, 7, 8)
	}
}

func BenchmarkValueMethodCallInt9(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dummy.m1int9(1, 2, 3, 4, 5, 6, 7, 8, 9)
	}
}

func BenchmarkValueMethodCallInt10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dummy.m1int10(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	}
}

func BenchmarkPointerMethodCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dummy.m2()
	}
}

func BenchmarkPointerMethodCallArg1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dummy.m2arg1(1)
	}
}

func BenchmarkPointerMethodCallArg2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dummy.m2arg2(1, 2)
	}
}

func BenchmarkPointerMethodCallArg3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dummy.m2arg3(1, 2, 3)
	}
}

func BenchmarkPointerMethodCallArg4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dummy.m2arg4(1, 2, 3, 4)
	}
}

func BenchmarkPointerMethodCallArg5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dummy.m2arg5(1, 2, 3, 4, 5)
	}
}

func BenchmarkPointerMethodCallArg6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dummy.m2arg6(1, 2, 3, 4, 5, 6)
	}
}

func BenchmarkPointerMethodCallArg7(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dummy.m2arg7(1, 2, 3, 4, 5, 6, 7)
	}
}

func BenchmarkPointerMethodCallArg8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dummy.m2arg8(1, 2, 3, 4, 5, 6, 7, 8)
	}
}

func BenchmarkPointerMethodCallArg9(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dummy.m2arg9(1, 2, 3, 4, 5, 6, 7, 8, 9)
	}
}

func BenchmarkPointerMethodCallArg10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dummy.m2arg10(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	}
}

func BenchmarkPointerMethodCallInt1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dummy.m2int1(1)
	}
}

func BenchmarkPointerMethodCallInt2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dummy.m2int2(1, 2)
	}
}

func BenchmarkPointerMethodCallInt3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dummy.m2int3(1, 2, 3)
	}
}

func BenchmarkPointerMethodCallInt4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dummy.m2int4(1, 2, 3, 4)
	}
}

func BenchmarkPointerMethodCallInt5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dummy.m2int5(1, 2, 3, 4, 5)
	}
}

func BenchmarkPointerMethodCallInt6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dummy.m2int6(1, 2, 3, 4, 5, 6)
	}
}

func BenchmarkPointerMethodCallInt7(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dummy.m2int7(1, 2, 3, 4, 5, 6, 7)
	}
}

func BenchmarkPointerMethodCallInt8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dummy.m2int8(1, 2, 3, 4, 5, 6, 7, 8)
	}
}

func BenchmarkPointerMethodCallInt9(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dummy.m2int9(1, 2, 3, 4, 5, 6, 7, 8, 9)
	}
}

func BenchmarkPointerMethodCallInt10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dummy.m2int10(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	}
}

func BenchmarkValueMethodCallVarArgs(b *testing.B) {
	for _, v := range scaling_slices {
		vv := make([]any, len(v))
		for i, w := range v {
			vv[i] = w
		}
		b.Run(fmt.Sprintf("count_%v", len(vv)), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				dummy.m1varargs(vv...)
			}
		})
	}
}

func BenchmarkValueMethodCallVarInts(b *testing.B) {
	for _, v := range scaling_slices {
		vv := make([]int, len(v))
		for i, w := range v {
			vv[i] = w
		}
		b.Run(fmt.Sprintf("count_%v", len(vv)), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				dummy.m1varints(vv...)
			}
		})
	}
}

func BenchmarkPointerMethodCallVarArgs(b *testing.B) {
	for _, v := range scaling_slices {
		vv := make([]any, len(v))
		for i, w := range v {
			vv[i] = w
		}
		b.Run(fmt.Sprintf("count_%v", len(vv)), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				dummy.m2varargs(vv...)
			}
		})
	}
}

func BenchmarkPointerMethodCallVarInts(b *testing.B) {
	for _, v := range scaling_slices {
		vv := make([]int, len(v))
		for i, w := range v {
			vv[i] = w
		}
		b.Run(fmt.Sprintf("count_%v", len(vv)), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				dummy.m2varints(vv...)
			}
		})
	}
}

func BenchmarkValueMethodGetSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = dummyAccessor.get(0)
	}
}

func BenchmarkValueMethodSetSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dummyAccessor.set(0, 1)
	}
}

func BenchmarkPointerMethodGetSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = (&dummyAccessor).get(0)
	}
}

func BenchmarkPointerMethodSetSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		(&dummyAccessor).set(0, 1)
	}
}

func BenchmarkInterfaceMethodGetSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = dai.get(0)
	}
}

func BenchmarkInterfaceMethodSetSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dai.set(0, 1)
	}
}
