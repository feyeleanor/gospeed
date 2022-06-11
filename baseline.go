package gospeed

var x int
var u uint
var x32 int32
var u32 uint32
var x64 int64
var u64 uint64
var in any = 0
var h map[int]int = map[int]int{10000000: 0, 10000001: 1}
var hs map[string]int = map[string]int{"10000000": 0, "10000001": 1}
var f func() = func() {}
var farg1 func(x any) = func(x any) {}
var farg2 func(x, y any) = func(x, y any) {}
var farg3 func(x, y, z any) = func(x, y, z any) {}
var farg4 func(x, y, z, a any) = func(x, y, z, a any) {}
var farg5 func(x, y, z, a, b any) = func(x, y, z, a, b any) {}
var farg6 func(x, y, z, a, b, c any) = func(x, y, z, a, b, c any) {}
var farg7 func(x, y, z, a, b, c, d any) = func(x, y, z, a, b, c, d any) {}
var farg8 func(x, y, z, a, b, c, d, e any) = func(x, y, z, a, b, c, d, e any) {}
var farg9 func(x, y, z, a, b, c, d, e, f any) = func(x, y, z, a, b, c, d, e, f any) {}
var farg10 func(x, y, z, a, b, c, d, e, f, g any) = func(x, y, z, a, b, c, d, e, f, g any) {}
var fvarargs func(x ...any) = func(x ...any) {}
var fint1 func(x int) = func(x int) {}
var fint2 func(x, y int) = func(x, y int) {}
var fint3 func(x, y, z int) = func(x, y, z int) {}
var fint4 func(x, y, z, a int) = func(x, y, z, a int) {}
var fint5 func(x, y, z, a, b int) = func(x, y, z, a, b int) {}
var fint6 func(x, y, z, a, b, c int) = func(x, y, z, a, b, c int) {}
var fint7 func(x, y, z, a, b, c, d int) = func(x, y, z, a, b, c, d int) {}
var fint8 func(x, y, z, a, b, c, d, e int) = func(x, y, z, a, b, c, d, e int) {}
var fint9 func(x, y, z, a, b, c, d, e, f int) = func(x, y, z, a, b, c, d, e, f int) {}
var fint10 func(x, y, z, a, b, c, d, e, f, g int) = func(x, y, z, a, b, c, d, e, f, g int) {}
var fvarints func(x ...int) = func(x ...int) {}

type dummyInterface1 interface {
	m1()
}

type dummyInterface2 interface {
	dummyInterface1
	m1arg(x any)
}

type dummyStructure struct {
	i  int
	in any
	s  []int
	h  map[int]int
}

func (d dummyStructure) m1()                 {}
func (d dummyStructure) m1arg(x any)         {}
func (d dummyStructure) m1int(x int)         {}
func (d dummyStructure) m1varargs(x ...any)  {}
func (d dummyStructure) m1varints(x ...int)  {}
func (d *dummyStructure) m2()                {}
func (d *dummyStructure) m2arg(x any)        {}
func (d *dummyStructure) m2int(x int)        {}
func (d *dummyStructure) m2varargs(x ...any) {}
func (d *dummyStructure) m2varints(x ...int) {}
func (d dummyStructure) get(i int) (r int)   { return d.s[i] }
func (d dummyStructure) set(i, x int)        { d.s[i] = x }

var dummy dummyStructure = dummyStructure{}
var di any = dummyStructure{}

type dummyInterface3 interface {
	get(int) int
	set(int, int)
}

var dummyAccessor dummyStructure = dummyStructure{s: []int{0}}
var dai dummyInterface3 = dummyStructure{s: []int{0}}
