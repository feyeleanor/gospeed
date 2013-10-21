package gospeed

import "testing"

func BenchmarkBaselineCastInt32ToInt(b *testing.B) {
	for i := 0; i < b.N; i++ { _ = int(x32) }
}

func BenchmarkBaselineCastIntToInt32(b *testing.B) {
	for i := 0; i < b.N; i++ { _ = int32(x) }
}

func BenchmarkBaselineCastInt64ToUint64(b *testing.B) {
	for i := 0; i < b.N; i++ { _ = uint64(x64) }
}

func BenchmarkBaselineCastUint64ToInt64(b *testing.B) {
	for i := 0; i < b.N; i++ { _ = int64(u64) }
}

func BenchmarkBaselineVariableGet(b *testing.B) {
	for i := 0; i < b.N; i++ { _ = x }
}

func BenchmarkBaselineVariableSet(b *testing.B) {
	for i := 0; i < b.N; i++ { x = 1 }
}

func BenchmarkBaselineVariableGetInterface(b *testing.B) {
	for i := 0; i < b.N; i++ { _ = in }
}

func BenchmarkBaselineVariableSetInterface(b *testing.B) {
	for i := 0; i < b.N; i++ { in = 1 }
}

func BenchmarkBaselineVariableIncrement(b *testing.B) {
	for i := 0; i < b.N; i++ { x++ }
}

func BenchmarkBaselineVariableDecrement(b *testing.B) {
	for i := 0; i < b.N; i++ { x-- }
}

func BenchmarkBaselineVariableIntegerAdd(b *testing.B) {
	x, y := 1, 2
	for i := 0; i < b.N; i++ { _ = x + y }
}

func BenchmarkBaselineVariableIntegerSubtract(b *testing.B) {
	x, y := 1, 2
	for i := 0; i < b.N; i++ { _ = x - y }
}

func BenchmarkBaselineVariableIntegerMultiply(b *testing.B) {
	x, y := 1, 2
	for i := 0; i < b.N; i++ { _ = x * y }
}

func BenchmarkBaselineVariableIntegerDivide(b *testing.B) {
	x, y := 1, 2
	for i := 0; i < b.N; i++ { _ = x / y }
}

func BenchmarkBaselineVariableIntegerRemainder(b *testing.B) {
	x, y := 1, 2
	for i := 0; i < b.N; i++ { _ = x % y }
}

func BenchmarkBaselineFieldGet(b *testing.B) {
	for i := 0; i < b.N; i++ { _ = dummy.i }
}

func BenchmarkBaselineFieldSet(b *testing.B) {
	for i := 0; i < b.N; i++ { dummy.i = 1 }
}

func BenchmarkBaselineSliceGet(b *testing.B) {
	for i := 0; i < b.N; i++ { _ = s[0] }
}

func BenchmarkBaselineSliceSet(b *testing.B) {
	for i := 0; i < b.N; i++ { s[0] = 1 }
}

func BenchmarkBaselineMapIntGet(b *testing.B) {
	for i := 0; i < b.N; i++ { _ = h[0] }
}

func BenchmarkBaselineMapIntSet(b *testing.B) {
	for i := 0; i < b.N; i++ { h[0] = 1 }
}

func BenchmarkBaselineMapStringGet(b *testing.B) {
	for i := 0; i < b.N; i++ { _ = hs["0"] }
}

func BenchmarkBaselineMapStringSet(b *testing.B) {
	for i := 0; i < b.N; i++ { hs["0"] = 1 }
}

func BenchmarkBaselineIf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if i >= 0 {}
	}
}

func BenchmarkBaselineIfElse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if i < 0 {} else {}
	}
}

func BenchmarkBaselineSwitchDefault(b *testing.B) {
	for i := 0; i < b.N; i++ {
		switch i {
		default: {}
		}
	}
}

func BenchmarkBaselineSwitchOneCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		switch i {
		case 0:	{}
		}
	}
}

func BenchmarkBaselineSwitchTwoCases(b *testing.B) {
	for i := 0; i < b.N; i++ {
		switch i {
		case 0:	{}
		case 1:	{}
		}
	}
}

func BenchmarkBaselineSwitchTwoCasesFallthrough(b *testing.B) {
	for i := 0; i < b.N; i++ {
		switch i {
		case 0:	fallthrough
		case 1:	{}
		}
	}
}

func BenchmarkBaselineForLoopIteration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < 1; j++ {  }
	}
}

func BenchmarkBaselineForReverseLoopIteration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 1; j > 0; j-- {  }
	}
}

func BenchmarkBaselineForRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, _ = range s {}
	}
}

func BenchmarkBaselineForSliceLength(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(s); j++ {}
	}
}

func BenchmarkBaselineForReverseSliceLength(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := len(s); j > 0; j-- {}
	}
}

func BenchmarkBaselineForLoopIteration10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10; j++ {  }
	}
}

func BenchmarkBaselineForReverseLoopIteration10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 10; j > 0; j-- {  }
	}
}

func BenchmarkBaselineForRange10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, _ = range s10 {}
	}
}

func BenchmarkBaselineForSliceLength10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(s10); j++ {}
	}
}

func BenchmarkBaselineForReverseSliceLength10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := len(s10); j > 0; j-- {}
	}
}

func BenchmarkBaselineForLoopIteration100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < 100; j++ {  }
	}
}

func BenchmarkBaselineForReverseLoopIteration100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 100; j > 0; j-- {  }
	}
}

func BenchmarkBaselineForRange100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, _ = range s100 {}
	}
}

func BenchmarkBaselineForSliceLength100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(s100); j++ {}
	}
}

func BenchmarkBaselineForReverseSliceLength100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := len(s100); j > 0; j-- {}
	}
}

func BenchmarkBaselineForLoopIteration10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {  }
	}
}

func BenchmarkBaselineForReverseLoopIteration10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 10000; j > 0; j-- {  }
	}
}

func BenchmarkBaselineForRange10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, _ = range s10000 {}
	}
}

func BenchmarkBaselineForSliceLength10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(s10000); j++ {}
	}
}

func BenchmarkBaselineForReverseSliceLength10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := len(s10000); j > 0; j-- {}
	}
}

func BenchmarkBaselineFunctionCall(b *testing.B) {
	for i := 0; i < b.N; i++ { f() }
}

func BenchmarkBaselineFunctionCallArg(b *testing.B) {
	for i := 0; i < b.N; i++ { farg(1) }
}

func BenchmarkBaselineFunctionCall5VarArgs(b *testing.B) {
	for i := 0; i < b.N; i++ { fvarargs(1, 2, 3, 4, 5) }
}

func BenchmarkBaselineFunctionCallInt(b *testing.B) {
	for i := 0; i < b.N; i++ { fint(1) }
}

func BenchmarkBaselineFunctionCall5VarInts(b *testing.B) {
	for i := 0; i < b.N; i++ { fvarints(1, 2, 3, 4, 5) }
}

func BenchmarkBaselineFunctionCallWithDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		func() {
			defer func() {}()
		}()
	}
}

func BenchmarkBaselineDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		defer func() {}()
	}
}

func BenchmarkBaselineTypeAssertion(b *testing.B) {
	for i := 0; i < b.N; i++ { _ = in.(int) }
}

func BenchmarkBaselineTypeAssertionEmptyInterface(b *testing.B) {
	for i := 0; i < b.N; i++ { _ = in.(interface{}) }
}

func BenchmarkBaselineTypeAssertionInterface1(b *testing.B) {
	for i := 0; i < b.N; i++ { _ = di.(dummyInterface1) }
}

func BenchmarkBaselineTypeAssertionInterface2(b *testing.B) {
	for i := 0; i < b.N; i++ { _ = di.(dummyInterface2) }
}

func BenchmarkBaselineTypeCheck(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, ok := in.(int); ok {}
	}
}

func BenchmarkBaselineTypeCheckEmptyInterface(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, ok := in.(interface{}); ok {}
	}
}

func BenchmarkBaselineTypeCheckInterface1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, ok := di.(dummyInterface1); ok {}
	}
}

func BenchmarkBaselineTypeCheckInterface2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, ok := di.(dummyInterface2); ok {}
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

func BenchmarkBaselineNewStructureLiteral(b *testing.B) {
	for i := 0; i < b.N; i++ { _ = dummyStructure{} }
}

func BenchmarkBaselineNewStructure(b *testing.B) {
	for i := 0; i < b.N; i++ { _ = *new(dummyStructure) }
}

func BenchmarkBaselineNewSliceLiteral(b *testing.B) {
	for i := 0; i < b.N; i++ { _ = []int{0, 1} }
}

func BenchmarkBaselineNewSlice(b *testing.B) {
	for i := 0; i < b.N; i++ { _ = make([]int, 2) }
}

func BenchmarkBaselineNewMapLiteralIntToInt(b *testing.B) {
	for i := 0; i < b.N; i++ { _ = map [int]int{} }
}

func BenchmarkBaselineNewMapLiteralIntToInterface(b *testing.B) {
	for i := 0; i < b.N; i++ { _ = map [int]interface{}{} }
}

func BenchmarkBaselineNewMapLiteralStringToInt(b *testing.B) {
	for i := 0; i < b.N; i++ { _ = map [string]int{} }
}

func BenchmarkBaselineNewMapLiteralStringToInterface(b *testing.B) {
	for i := 0; i < b.N; i++ { _ = map [string]interface{}{} }
}

func BenchmarkBaselineNewMapLiteralIntToInt2Item(b *testing.B) {
	for i := 0; i < b.N; i++ { _ = map [int]int{i: 0, i: 1} }
}

func BenchmarkBaselineNewMapLiteralIntToInterface2Item(b *testing.B) {
	for i := 0; i < b.N; i++ { _ = map [int]interface{}{i: 0, i: 1} }
}

func BenchmarkBaselineNewMapLiteralStringToInt2Item(b *testing.B) {
	for i := 0; i < b.N; i++ { _ = map [string]int{"0": i, "1": i} }
}

func BenchmarkBaselineNewMapLiteralStringToInterface2Item(b *testing.B) {
	for i := 0; i < b.N; i++ { _ = map [string]interface{}{"0": i, "1": i} }
}

func BenchmarkBaselineNewMapIntToInt(b *testing.B) {
	for i := 0; i < b.N; i++ { _ = make(map [int]int) }
}

func BenchmarkBaselineNewMapIntToInterface(b *testing.B) {
	for i := 0; i < b.N; i++ { _ = make(map [int]interface{}) }
}

func BenchmarkBaselineNewMapStringToInt(b *testing.B) {
	for i := 0; i < b.N; i++ { _ = make(map [string]int) }
}

func BenchmarkBaselineNewMapStringToInterface(b *testing.B) {
	for i := 0; i < b.N; i++ { _ = make(map [string]interface{}) }
}

func BenchmarkBaselineSliceCopy(b *testing.B) {
	for i := 0; i < b.N; i++ { copy(d10, s10) }
}

func BenchmarkBaselineNewSliceAppendElement1(b *testing.B) {
	for i := 0; i < b.N; i++ { _ = append(s, 1) }
}

func BenchmarkBaselineNewSliceAppendElement10(b *testing.B) {
	for i := 0; i < b.N; i++ { _ = append(s, s10...) }
}