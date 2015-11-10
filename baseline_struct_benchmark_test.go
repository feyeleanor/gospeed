package gospeed

import "testing"

func BenchmarkBaselineNewStructureLiteral(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = dummyStructure{}
	}
}

func BenchmarkBaselineNewStructure(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = *new(dummyStructure)
	}
}

func BenchmarkBaselineFieldGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = dummy.i
	}
}

func BenchmarkBaselineFieldSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dummy.i = 1
	}
}
