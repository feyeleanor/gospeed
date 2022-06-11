package gospeed

import (
	"reflect"
	"testing"
)

func BenchmarkReflectPrimitiveToValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reflect.ValueOf(in)
	}
}

func BenchmarkReflectPrimitiveToValueIsValid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reflect.ValueOf(in).IsValid()
	}
}

func BenchmarkReflectSliceToValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reflect.ValueOf(s)
	}
}

func BenchmarkReflectSliceToValueIsValid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reflect.ValueOf(s).IsValid()
	}
}

func BenchmarkReflectStructToValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reflect.ValueOf(dummy)
	}
}

func BenchmarkReflectStructToValueIsValid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reflect.ValueOf(dummy).IsValid()
	}
}

func BenchmarkReflectPtrToValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reflect.ValueOf(&dummy)
	}
}

func BenchmarkReflectPtrToValueIsValid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reflect.ValueOf(&dummy).IsValid()
	}
}

func BenchmarkReflectPtrToValueElem(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reflect.ValueOf(&dummy).Elem()
	}
}

func BenchmarkReflectPtrToValueElemIsValid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reflect.ValueOf(&dummy).Elem().IsValid()
	}
}

func BenchmarkReflectInterfaceToValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reflect.ValueOf(new(dummyInterface1))
	}
}

func BenchmarkReflectInterfaceToValueIsValid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reflect.ValueOf(new(dummyInterface1)).IsValid()
	}
}

func BenchmarkReflectInterfaceToValueElem(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reflect.ValueOf(new(dummyInterface1)).Elem()
	}
}

func BenchmarkReflectInterfaceToValueElemIsValid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reflect.ValueOf(new(dummyInterface1)).Elem().IsValid()
	}
}

func paramFree(p []reflect.Value) (r []reflect.Value) {
	return
}

func makeParamFree(fptr any) {
	f := reflect.ValueOf(fptr).Elem()
	v := reflect.MakeFunc(f.Type(), paramFree)
	f.Set(v)
}

func BenchmarkReflectMakeFunc(b *testing.B) {
	var f func()
	for i := 0; i < b.N; i++ {
		makeParamFree(&f)
	}
}

func BenchmarkReflectMakeFuncThenCall(b *testing.B) {
	var f func()
	for i := 0; i < b.N; i++ {
		makeParamFree(&f)
		f()
	}
}

func sink(in []reflect.Value) []reflect.Value {
	return []reflect.Value{in[0]}
}

func makeSink(fptr any) {
	f := reflect.ValueOf(fptr).Elem()
	v := reflect.MakeFunc(f.Type(), sink)
	f.Set(v)
}

func BenchmarkReflectMakeFuncSinkArg(b *testing.B) {
	var f func(int)
	for i := 0; i < b.N; i++ {
		makeSink(&f)
	}
}

func BenchmarkReflectMakeFuncSinkThenCallArg(b *testing.B) {
	var f func(int) int
	for i := 0; i < b.N; i++ {
		makeSink(&f)
		f(0)
	}
}

func passthrough(in []reflect.Value) []reflect.Value {
	return []reflect.Value{in[0]}
}

func makePassthrough(fptr any) {
	f := reflect.ValueOf(fptr).Elem()
	v := reflect.MakeFunc(f.Type(), passthrough)
	f.Set(v)
}

func BenchmarkReflectMakeFuncPassthroughArg(b *testing.B) {
	var f func(int) int
	for i := 0; i < b.N; i++ {
		makePassthrough(&f)
	}
}

func BenchmarkReflectMakeFuncPassthroughThenCallArg(b *testing.B) {
	var f func(int) int
	for i := 0; i < b.N; i++ {
		makePassthrough(&f)
		f(0)
	}
}
