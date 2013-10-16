package gospeed

import (
	"reflect"
	"testing"
)

func BenchmarkReflectPrimitiveToValue(b *testing.B) {
	for i := 0; i < b.N; i++ { reflect.ValueOf(in) }
}

func BenchmarkReflectPrimitiveToValueIsValid(b *testing.B) {
	for i := 0; i < b.N; i++ { reflect.ValueOf(in).IsValid() }
}

func BenchmarkReflectSliceToValue(b *testing.B) {
	for i := 0; i < b.N; i++ { reflect.ValueOf(s) }
}

func BenchmarkReflectSliceToValueIsValid(b *testing.B) {
	for i := 0; i < b.N; i++ { reflect.ValueOf(s).IsValid() }
}

func BenchmarkReflectStructToValue(b *testing.B) {
	for i := 0; i < b.N; i++ { reflect.ValueOf(dummy) }
}

func BenchmarkReflectStructToValueIsValid(b *testing.B) {
	for i := 0; i < b.N; i++ { reflect.ValueOf(dummy).IsValid() }
}

func BenchmarkReflectPtrToValue(b *testing.B) {
	for i := 0; i < b.N; i++ { reflect.ValueOf(&dummy) }
}

func BenchmarkReflectPtrToValueIsValid(b *testing.B) {
	for i := 0; i < b.N; i++ { reflect.ValueOf(&dummy).IsValid() }
}

func BenchmarkReflectPtrToValueElem(b *testing.B) {
	for i := 0; i < b.N; i++ { reflect.ValueOf(&dummy).Elem() }
}

func BenchmarkReflectPtrToValueElemIsValid(b *testing.B) {
	for i := 0; i < b.N; i++ { reflect.ValueOf(&dummy).Elem().IsValid() }
}

func BenchmarkReflectInterfaceToValue(b *testing.B) {
	for i := 0; i < b.N; i++ { reflect.ValueOf(new(dummyInterface1)) }
}

func BenchmarkReflectInterfaceToValueIsValid(b *testing.B) {
	for i := 0; i < b.N; i++ { reflect.ValueOf(new(dummyInterface1)).IsValid() }
}

func BenchmarkReflectInterfaceToValueElem(b *testing.B) {
	for i := 0; i < b.N; i++ { reflect.ValueOf(new(dummyInterface1)).Elem() }
}

func BenchmarkReflectInterfaceToValueElemIsValid(b *testing.B) {
	for i := 0; i < b.N; i++ { reflect.ValueOf(new(dummyInterface1)).Elem().IsValid() }
}